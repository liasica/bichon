package service

import (
    "context"
    "crypto/ecdsa"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/ent/group"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
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
func (s *groupService) Create(mem *ent.Member, req *model.GroupCreateReq) (gro any, err error) {
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

    address := crypto.PubkeyToAddress(pub).Hex()

    // create group
    _, _ = s.orm.Create().
        SetName(req.Name).
        SetNillableIntro(req.Intro).
        SetKeys(b).
        SetAddress(address).
        SetMemberID(mem.ID).
        Save(s.ctx)

    return
}
