package service

import (
    "context"
    "crypto/ecdsa"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/ent/group"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
    "strings"
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
    }
    // check if the number of created groups exceeds the maximum number
    cnt, _ := s.orm.Query().Where(group.MemberID(mem.ID)).Count(s.ctx)
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

    var b []byte
    b, err = keys.Encrypt()
    if err != nil {
        return
    }

    address := strings.ToLower(crypto.PubkeyToAddress(pub).Hex())

    gp := true
    if req.Public != nil {
        gp = *req.Public
    }

    // create group
    var gro *ent.Group
    gro, err = s.orm.Create().
        SetName(req.Name).
        SetMembersMax(req.MaxMembers).
        SetNillableIntro(req.Intro).
        SetPublic(gp).
        SetKeys(hexutil.Encode(b)).
        SetAddress(address).
        SetMemberID(mem.ID).
        SetCategory(req.Category).
        Save(s.ctx)

    res = s.detail(gro, mem)

    return
}

func (s *groupService) detail(gro *ent.Group, mem *ent.Member) *model.GroupDetail {
    return &model.GroupDetail{
        Name:       gro.Name,
        Address:    gro.Address,
        Intro:      gro.Intro,
        MembersMax: gro.MembersMax,
        Public:     gro.Public,
        Owner:      mem.ID == gro.MemberID,
        Category:   gro.Category,
    }
}
