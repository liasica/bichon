package service

import (
    "context"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "errors"
    "github.com/chatpuppy/puppychat/internal/ent"
    "math/big"
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

func (s *groupService) GenerateKey() (pri *ecdsa.PrivateKey, pub *ecdsa.PublicKey, err error) {
    pri, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        return
    }
    pub = &pri.PublicKey
    return
}

func (s *groupService) GenerateShareKey(pubKey []string) (shared [sha256.Size]byte, err error) {
    if len(pubKey) != 2 {
        err = errors.New("public key error")
        return
    }

    x := new(big.Int)
    x.SetString(pubKey[0], 16)

    y := new(big.Int)
    y.SetString(pubKey[1], 16)

    var pri *ecdsa.PrivateKey
    var pub *ecdsa.PublicKey

    pri, pub, err = s.GenerateKey()
    if err != nil {
        return
    }
    a, _ := pub.Curve.ScalarMult(x, y, pri.D.Bytes())
    shared = sha256.Sum256(a.Bytes())
    return
}
