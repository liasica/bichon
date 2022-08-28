package service

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/ent/group"
    "github.com/chatpuppy/puppychat/internal/ent/groupmember"
    "github.com/chatpuppy/puppychat/internal/ent/key"
    "github.com/chatpuppy/puppychat/utils"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
    "strings"
    "time"
)

type groupService struct {
    ctx context.Context
    orm *ent.GroupClient
}

func NewGroup() *groupService {
    return &groupService{
        ctx: context.Background(),
        orm: ent.Database.Group,
    }
}

func (s *groupService) QueryID(id string) (*ent.Group, error) {
    return s.orm.Query().Where(group.ID(id)).First(s.ctx)
}

// Create create group
func (s *groupService) Create(mem *ent.Member, req *model.GroupCreateReq) (res *model.GroupDetailWithPublicKey, err error) {
    // check group create frequency limit
    if exist, _ := s.orm.Query().Where(group.OwnerID(mem.ID), group.CreatedAtGT(time.Now().Add(-time.Duration(model.GroupCreateFrequency)*time.Second))).Exist(s.ctx); exist {
        err = model.ErrCreateFrequency
        return
    }

    // check group max members count
    if req.MaxMembers > model.GroupMaxMembers {
        err = model.ErrMaximunMembers
    }
    // check if the number of created groups exceeds the maximum number
    cnt, _ := s.orm.Query().Where(group.OwnerID(mem.ID)).Count(s.ctx)
    if cnt >= model.GroupMax {
        err = model.ErrMaximunGroups
        return
    }

    // create an account
    var priv *ecdsa.PrivateKey
    priv, err = crypto.GenerateKey()
    if err != nil {
        return
    }

    pub := priv.PublicKey

    keys := &model.GroupKeys{
        PrivateKey: hexutil.Encode(crypto.FromECDSA(priv)),
        PublicKey:  hexutil.Encode(crypto.FromECDSAPub(&pub)),
    }

    var keysHex string
    keysHex, err = keys.Encrypt()
    if err != nil {
        return
    }

    address := strings.ToLower(crypto.PubkeyToAddress(pub).Hex())

    gp := true
    if req.Public != nil {
        gp = *req.Public
    }

    err = ent.WithTx(s.ctx, func(tx *ent.Tx) (err error) {
        // create group
        var gro *ent.Group
        gro, err = tx.Group.Create().
            SetName(req.Name).
            SetMembersMax(req.MaxMembers).
            SetNillableIntro(req.Intro).
            SetPublic(gp).
            SetKeys(keysHex).
            SetAddress(address).
            SetOwnerID(mem.ID).
            SetCategory(req.Category).
            SetMembersCount(1).
            Save(s.ctx)
        if err != nil {
            return
        }

        // join group
        var gkeys *model.GroupMemberKeys
        var keyID string
        gkeys, keyID, err = s.joinGroup(tx, mem, gro, model.GroupMemberPermOwner, req.SharedPublic, true)
        if err != nil {
            return
        }

        res = &model.GroupDetailWithPublicKey{
            GroupPublicKey: &model.GroupPublicKey{
                GroupPublicKey: gkeys.Public,
                KeyID:          keyID,
            },
            GroupDetail: s.detail(gro, mem),
        }
        return
    })

    return
}

// Join join group
func (s *groupService) Join(mem *ent.Member, req *model.GroupJoinReq) (res *model.GroupDetailWithPublicKey, err error) {
    // find group
    var gro *ent.Group
    gro, _ = s.QueryID(req.GroupID)
    if gro == nil {
        err = model.ErrNotFoundGroup
        return
    }
    // if member already in group
    if s.MemberIn(mem.ID, gro.ID) {
        err = model.ErrAlreadyInGroup
        return
    }

    var keys *model.GroupMemberKeys
    var keyID string
    err = ent.WithTx(s.ctx, func(tx *ent.Tx) error {
        keys, keyID, err = s.joinGroup(tx, mem, gro, model.GroupMemberPermDefault, req.SharedPublic, false)
        return err
    })
    if err != nil {
        return
    }

    res = &model.GroupDetailWithPublicKey{
        GroupPublicKey: &model.GroupPublicKey{
            GroupPublicKey: keys.Public,
            KeyID:          keyID,
        },
        GroupDetail: s.detail(gro, mem),
    }
    return
}

// MemberIn check member in group or not
func (s *groupService) MemberIn(memberID string, groupID string) (exists bool) {
    exists, _ = ent.Database.GroupMember.Query().Where(groupmember.GroupID(groupID), groupmember.MemberID(memberID)).Exist(s.ctx)
    return
}

// shareKey share key with group
func (s *groupService) shareKey(tx *ent.Tx, mem *ent.Member, gro *ent.Group, spKey string) (keys *model.GroupMemberKeys, keyID string, err error) {
    // generate keys
    keys, err = NewKey().GenerateAndShare(spKey)
    if err != nil {
        return
    }

    var hex string
    hex, err = keys.Encrypt()
    if err != nil {
        return
    }

    // create group member keys
    var k *ent.Key
    k, err = tx.Key.Create().SetKeys(hex).SetGroupID(gro.ID).SetMemberID(mem.ID).Save(s.ctx)
    if err != nil {
        return
    }

    keyID = k.ID

    return
}

// joinGroup join group and share ecdh key
func (s *groupService) joinGroup(tx *ent.Tx, mem *ent.Member, gro *ent.Group, perm uint8, spKey string, isCreate bool) (keys *model.GroupMemberKeys, keyID string, err error) {
    keys, keyID, err = s.shareKey(tx, mem, gro, spKey)

    // create group member set share sn and permission
    _, err = tx.GroupMember.Create().
        SetGroup(gro).
        SetMember(mem).
        SetSn(utils.Md5String([]byte(fmt.Sprintf("%d%d%d", mem.ID, gro.ID, time.Now().UnixNano())))).
        SetPermission(perm).
        Save(s.ctx)
    if err != nil {
        return
    }

    if !isCreate {
        // update group's member count
        _, err = tx.Group.UpdateOne(gro).AddMembersCount(1).Save(s.ctx)
    }
    return
}

// detail group's detail
func (s *groupService) detail(gro *ent.Group, mem *ent.Member) *model.GroupDetail {
    return &model.GroupDetail{
        ID:         gro.ID,
        Name:       gro.Name,
        Address:    gro.Address,
        Intro:      gro.Intro,
        MembersMax: gro.MembersMax,
        Public:     gro.Public,
        Owner:      mem.ID == gro.OwnerID,
        Category:   gro.Category,
    }
}

// ShareKey share key
func (s *groupService) ShareKey(mem *ent.Member, req *model.GroupShareKeyReq) (res *model.GroupPublicKey, err error) {
    // find group
    var gro *ent.Group
    gro, _ = s.QueryID(req.GroupID)
    if gro == nil {
        err = model.ErrNotFoundGroup
        return
    }

    // if member not in group
    if !s.MemberIn(mem.ID, gro.ID) {
        err = model.ErrNotInGroup
        return
    }

    // check group create frequency limit
    if exist, _ := ent.Database.Key.Query().Where(
        key.MemberID(mem.ID),
        key.GroupID(req.GroupID),
        key.CreatedAtGT(time.Now().Add(-time.Duration(model.GroupKeyShareFrequency)*time.Second)),
    ).Exist(s.ctx); exist {
        err = model.ErrKeyShareFrequency
        return
    }

    // share key
    var keys *model.GroupMemberKeys
    var keyID string
    err = ent.WithTx(s.ctx, func(tx *ent.Tx) error {
        keys, keyID, err = s.shareKey(tx, mem, gro, req.SharedPublic)
        return err
    })
    if err != nil {
        return
    }

    res = &model.GroupPublicKey{
        GroupPublicKey: keys.Public,
        KeyID:          keyID,
    }
    return
}
