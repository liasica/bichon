package service

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/ent/group"
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

// Create create group
func (s *groupService) Create(mem *ent.Member, req *model.GroupCreateReq) (res *model.GroupDetail, err error) {
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
        err = s.joinGroup(tx, mem, gro, model.GroupMemberPermOwner, req.SharedPublic, true)
        if err != nil {
            return
        }

        res = s.detail(gro, mem)
        return
    })

    return
}

func (s *groupService) Join(groupID uint64, memberID uint64, keys *model.GroupMemberKeys) {

}

func (s *groupService) joinGroup(tx *ent.Tx, mem *ent.Member, gro *ent.Group, perm uint8, spKey string, isCreate bool) (err error) {
    // generate keys
    var keys *model.GroupMemberKeys
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
    k, err = tx.Key.Create().SetKeys(hex).Save(s.ctx)
    if err != nil {
        return
    }

    // create group member set share sn and permission
    var gm *ent.GroupMember
    gm, err = tx.GroupMember.Create().
        SetGroup(gro).
        SetMember(mem).
        SetSn(utils.Md5String([]byte(fmt.Sprintf("%d%d%d", mem.ID, gro.ID, time.Now().UnixNano())))).
        SetPermission(perm).
        SetKey(k).
        Save(s.ctx)
    if err != nil {
        return
    }

    if !isCreate {
        // update group's member count
        _, err = tx.Group.UpdateOne(gro).AddMembersCount(1).AddGroupMembers(gm).Save(s.ctx)
    }
    return
}

func (s *groupService) detail(gro *ent.Group, mem *ent.Member) *model.GroupDetail {
    return &model.GroupDetail{
        Name:       gro.Name,
        Address:    gro.Address,
        Intro:      gro.Intro,
        MembersMax: gro.MembersMax,
        Public:     gro.Public,
        Owner:      mem.ID == gro.OwnerID,
        Category:   gro.Category,
    }
}
