package service

import (
    "context"
    "crypto/ecdsa"
    "entgo.io/ent/dialect/sql"
    "fmt"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/ent/group"
    "github.com/chatpuppy/puppychat/internal/ent/groupmember"
    "github.com/chatpuppy/puppychat/internal/ent/key"
    "github.com/chatpuppy/puppychat/utils"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/liasica/go-encryption/hexutil"
    "sort"
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

// QueryID query group by id
func (s *groupService) QueryID(id string) (*ent.Group, error) {
    return s.orm.Query().Where(group.ID(id)).First(s.ctx)
}

// QueryInviteCode query group and inviter
func (s *groupService) QueryInviteCode(code string) (gr *ent.Group, mem *ent.Member, err error) {
    gm, _ := ent.Database.GroupMember.Query().Where(
        groupmember.InviteCode(code),
        groupmember.InviteExpireGTE(time.Now()),
    ).WithGroup().WithMember().First(s.ctx)
    if gm == nil {
        err = model.ErrInviteCodeInvalid
        return
    }
    gr = gm.Edges.Group
    mem = gm.Edges.Member
    if gr == nil || mem == nil {
        err = model.ErrInviteCodeInvalid
    }
    return
}

// Meta group meta data
func (s *groupService) Meta(gro *ent.Group) model.GroupMeta {
    return model.GroupMeta{
        ID:           gro.ID,
        Name:         gro.Name,
        Address:      gro.Address,
        MembersMax:   gro.MembersMax,
        MembersCount: gro.MembersCount,
        Intro:        gro.Intro,
        Category:     gro.Category,
        CreatedAt:    gro.CreatedAt,
    }
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
            SetCategory(string(req.Category)).
            SetMembersCount(1).
            Save(s.ctx)
        if err != nil {
            return
        }

        // join group
        var gkeys *model.GroupMemberKeys
        var keyID, code string
        gkeys, keyID, code, err = s.joinGroup(tx, mem, gro, model.GroupMemberPermOwner, req.SharedPublic, true, nil)
        if err != nil {
            return
        }

        res = &model.GroupDetailWithPublicKey{
            GroupPublicKey: &model.GroupPublicKey{
                GroupPublicKey: gkeys.Public,
                KeyID:          keyID,
            },
            GroupDetail: s.detail(code, gro, mem),
        }
        return
    })

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
func (s *groupService) joinGroup(tx *ent.Tx, mem *ent.Member, gro *ent.Group, perm model.GroupMemberPerm, spKey string, isCreate bool, inviter *ent.Member) (keys *model.GroupMemberKeys, keyID, code string, err error) {
    keys, keyID, err = s.shareKey(tx, mem, gro, spKey)

    // create group member set share sn and permission
    ic, it := s.GenerateInviteCode(mem.ID, gro.ID)
    creater := tx.GroupMember.Create().
        SetGroup(gro).
        SetMember(mem).
        SetInviteCode(ic).
        SetInviteExpire(it).
        SetPermission(perm)
    if inviter != nil {
        creater.SetInviter(inviter)
    }
    _, err = creater.Save(s.ctx)
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
func (s *groupService) detail(code string, gro *ent.Group, mem *ent.Member) *model.GroupDetail {
    return &model.GroupDetail{
        GroupMeta:  s.Meta(gro),
        Public:     gro.Public,
        Owner:      mem.ID == gro.OwnerID,
        InviteCode: code,
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

// GenerateInviteCode generate group invite code
func (s *groupService) GenerateInviteCode(memberID, groupID string) (string, time.Time) {
    return utils.Md5String([]byte(fmt.Sprintf("%s%s%d", memberID, groupID, time.Now().UnixNano()))), time.Now().AddDate(0, 0, model.GroupInviteCodeExpires)
}

// regenerate group invite code
func (s *groupService) reGenerateInviteCode(grm *ent.GroupMember) (string, time.Time, error) {
    ic, it := s.GenerateInviteCode(grm.MemberID, grm.GroupID)
    return ic, it, ent.Database.GroupMember.UpdateOne(grm).SetInviteCode(ic).SetInviteExpire(it).Exec(s.ctx)
}

// Detail get group's detail
func (s *groupService) Detail(mem *ent.Member, req *model.GroupMetaReq) (res *model.GroupDetailRes, err error) {
    grm, _ := ent.Database.GroupMember.Query().Where(groupmember.GroupID(req.ID), groupmember.MemberID(mem.ID)).WithGroup(func(query *ent.GroupQuery) {
        query.WithGroupMembers(func(gmq *ent.GroupMemberQuery) {
            gmq.WithMember()
        })
    }).First(s.ctx)

    if grm == nil || grm.Edges.Group == nil {
        err = model.ErrNotFoundGroup
        return
    }

    gro := grm.Edges.Group

    var ic string

    if grm.InviteExpire.Before(time.Now()) {
        ic, _, err = s.reGenerateInviteCode(grm)
        if err != nil {
            return
        }
    } else {
        ic = grm.InviteCode
    }

    members := make([]model.MemberWithPermission, 0)
    for _, gm := range gro.Edges.GroupMembers {
        m := gm.Edges.Member
        if m != nil {
            members = append(members, model.MemberWithPermission{
                Permission: gm.Permission,
                Member: model.Member{
                    ID:       m.ID,
                    Address:  m.Address,
                    Nickname: m.Nickname,
                    Avatar:   m.Avatar,
                },
            })
        }
    }

    sort.Slice(members, func(i, j int) bool {
        return members[i].Permission > members[j].Permission
    })

    res = &model.GroupDetailRes{
        GroupDetail: s.detail(ic, gro, mem),
        Members:     members,
    }

    return
}

// KeyUsed set current key of group which member used
func (s *groupService) KeyUsed(mem *ent.Member, req *model.GroupKeyUsedReq) (res any, err error) {
    for _, item := range req.Keys {
        item.Keys, err = NewKey().QueryRawKeys(item.KeyID, mem.ID, item.GroupID)
        if err != nil {
            return
        }
        item.MemberID = mem.ID
        // cache keys
        model.StoreSharedKeys(mem.ID, item.GroupID, item.Keys)
    }
    return
}

// JoinedList list all joined group
func (s *groupService) JoinedList(mem *ent.Member) []model.GroupJoinedListRes {
    // get unread messages
    unread := NewMessage().UnreadList(mem.ID)
    // get groups
    items, _ := s.orm.Query().Where(group.HasGroupMembersWith(groupmember.MemberID(mem.ID))).All(s.ctx)
    res := make([]model.GroupJoinedListRes, len(items))
    ids := make([]string, len(items))
    for i, gro := range items {
        ur := unread[gro.ID]
        res[i] = model.GroupJoinedListRes{
            GroupMeta:   s.Meta(gro),
            UnreadCount: ur.Count,
            UnreadID:    ur.ID,
            UnreadTime:  ur.Time,
        }
        ids[i] = gro.ID
    }

    return res
}

// List pagination request all public groups
func (s *groupService) List(mem *ent.Member, req *model.GroupListReq) (res *model.PaginationRes, err error) {
    q := s.orm.Query().
        Where(
            group.Public(true),
        )
    if req.Name != "" {
        q.Where(group.NameContainsFold(req.Name))
    }
    if req.Category != "" {
        q.Where(group.Category(string(req.Category)))
    }
    var result []struct {
        ID           string    `json:"id"`
        Name         string    `json:"name"`
        Address      string    `json:"address"`
        Intro        string    `json:"intro"`
        Category     string    `json:"category"`
        MembersMax   int       `json:"members_max"`
        MembersCount int       `json:"members_count"`
        Joined       bool      `json:"joined"`
        CreatedAt    time.Time `json:"created_at"`
    }
    err = q.
        Limit(req.GetLimit()).
        Offset(req.GetOffset()).
        Order(ent.Desc(group.FieldID)).
        Modify(func(sel *sql.Selector) {
            sel.SetDistinct(true).
                Select(
                    sel.C(group.FieldID),
                    sel.C(group.FieldName),
                    sel.C(group.FieldAddress),
                    sel.C(group.FieldMembersMax),
                    sel.C(group.FieldMembersCount),
                    sel.C(group.FieldIntro),
                    sel.C(group.FieldCategory),
                    sel.C(group.FieldCreatedAt),
                )
            if mem != nil {
                t := sql.Table(groupmember.Table)
                p := sql.And(
                    sql.ColumnsEQ(t.C(groupmember.FieldGroupID), sel.C(group.FieldID)),
                    sql.EQ(t.C(groupmember.FieldMemberID), mem.ID),
                )
                sel.AppendSelectExprAs(sql.Exists(sql.Select().From(t).Where(p)), "joined")
            }
        }).
        Scan(s.ctx, &result)

    if err != nil {
        return
    }

    items := make([]model.GroupListRes, len(result))
    for i, gro := range result {
        items[i] = model.GroupListRes{
            GroupMeta: model.GroupMeta{
                ID:           gro.ID,
                Name:         gro.Name,
                Address:      gro.Address,
                MembersMax:   gro.MembersMax,
                MembersCount: gro.MembersCount,
                Intro:        gro.Intro,
                Category:     gro.Category,
                CreatedAt:    gro.CreatedAt,
            },
            Joined: gro.Joined,
        }
    }

    res = &model.PaginationRes{
        Pagination: q.PaginationResult(req.PaginationReq),
        Items:      items,
    }

    return
}

// Join group
func (s *groupService) Join(mem *ent.Member, req *model.GroupJoinReq) (res *model.GroupDetailWithPublicKey, err error) {
    // find group
    var (
        gro     *ent.Group
        inviter *ent.Member
    )

    switch true {
    case req.GroupID != "":
        gro, _ = s.QueryID(req.GroupID)
        if !gro.Public {
            err = model.ErrGroupPrivate
        }
    case req.InviteCode != "":
        gro, inviter, err = s.QueryInviteCode(req.InviteCode)
    }

    if err != nil {
        return
    }

    if gro == nil {
        err = model.ErrNotFoundGroup
        return
    }

    // if member already in group
    if s.MemberIn(mem.ID, gro.ID) {
        err = model.ErrAlreadyInGroup
        return
    }

    var (
        keys        *model.GroupMemberKeys
        keyID, code string
    )

    err = ent.WithTx(s.ctx, func(tx *ent.Tx) error {
        keys, keyID, code, err = s.joinGroup(tx, mem, gro, model.GroupMemberPermDefault, req.SharedPublic, false, inviter)
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
        GroupDetail: s.detail(code, gro, mem),
    }
    return
}

// Leave group
func (s *groupService) Leave(mem *ent.Member, req *model.GroupIDReq) error {
    if !NewGroup().MemberIn(mem.ID, req.GroupID) {
        return model.ErrNotInGroup
    }
    return ent.WithTx(s.ctx, func(tx *ent.Tx) (err error) {
        _, err = ent.Database.GroupMember.Delete().Where(groupmember.GroupID(req.GroupID), groupmember.MemberID(mem.ID)).Exec(s.ctx)
        if err != nil {
            return
        }
        _, err = tx.Group.UpdateOneID(req.GroupID).AddMembersCount(-1).Save(s.ctx)
        return
    })
}

// Update group
func (s *groupService) Update(mem *ent.Member, req *model.GroupUpdateReq) error {
    // find group and verify permission
    gro, _ := s.orm.Query().Where(group.OwnerID(mem.ID), group.ID(req.GroupID)).First(s.ctx)
    if gro == nil {
        return model.ErrNotFoundGroup
    }
    // change info
    updater := s.orm.UpdateOne(gro)
    if req.MaxMembers != nil {
        if gro.MembersCount > *req.MaxMembers {
            return model.ErrMaxMemberSetting
        }
        updater.SetMembersMax(*req.MaxMembers)
    }
    if req.Category != nil {
        if !req.Category.IsValid() {
            return model.ErrGroupCategory
        }
        updater.SetCategory(string(*req.Category))
    }
    if req.Public != nil {
        updater.SetPublic(*req.Public)
    }
    if req.Intro != nil {
        updater.SetNillableIntro(req.Intro)
    }
    return updater.Exec(s.ctx)
}

// ReGenerateInviteCode regenerate group invite code request
func (s *groupService) ReGenerateInviteCode(mem *ent.Member, req *model.GroupIDReq) (res *model.GroupInviteCodeRes, err error) {
    grm, _ := ent.Database.GroupMember.Query().Where(groupmember.GroupID(req.GroupID), groupmember.MemberID(mem.ID)).First(s.ctx)
    if grm == nil {
        err = model.ErrNotInGroup
        return
    }
    var ic string
    ic, _, err = s.reGenerateInviteCode(grm)
    if err != nil {
        return
    }
    res = &model.GroupInviteCodeRes{InviteCode: ic}
    return
}

// Active setting actived group
func (s *groupService) Active(mem *ent.Member, req *model.GroupIDReq) {
    model.GroupActived.Store(mem.ID, req.GroupID)
}

// Kickout member
// permission neede manager
func (s *groupService) Kickout(mem *ent.Member, req *model.GroupMemberReq) error {
    q := ent.Database.GroupMember.Query()
    // check permission
    gr, _ := q.Clone().Where(groupmember.GroupID(req.GroupID), groupmember.MemberID(mem.ID)).First(s.ctx)
    if !gr.Permission.IsManager() {
        return model.ErrInsufficientPermission
    }
    // check target member and permission
    tgr, _ := q.Clone().Where(groupmember.GroupID(req.GroupID), groupmember.MemberID(req.MemberID)).First(s.ctx)
    if tgr == nil {
        return model.ErrMemberNotInGroup
    }
    if tgr.Permission.IsManager() && !gr.Permission.IsOwner() {
        return model.ErrInsufficientPermission
    }
    // kick out
    return ent.Database.GroupMember.DeleteOne(tgr).Exec(s.ctx)
}

// SetManager set member permission to manager
func (s *groupService) SetManager(mem *ent.Member, req *model.GroupMemberReq) error {
    q := ent.Database.GroupMember.Query()
    // check permission
    gr, _ := q.Clone().Where(groupmember.GroupID(req.GroupID), groupmember.MemberID(mem.ID)).First(s.ctx)
    if !gr.Permission.IsOwner() {
        return model.ErrInsufficientPermission
    }
    // check target member and permission
    tgr, _ := q.Clone().Where(groupmember.GroupID(req.GroupID), groupmember.MemberID(req.MemberID)).First(s.ctx)
    if tgr == nil {
        return model.ErrMemberNotInGroup
    }
    if tgr.Permission.IsManager() {
        return model.ErrAlreadyManager
    }
    // set manager
    return ent.Database.GroupMember.UpdateOne(tgr).SetPermission(model.GroupMemberPermManager).Exec(s.ctx)
}
