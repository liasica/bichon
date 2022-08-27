package service

import (
    "context"
    "crypto/ecdsa"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/liasica/go-encryption/ecdh"
)

type keyService struct {
    ctx context.Context
}

func NewKey() *keyService {
    return &keyService{
        ctx: context.Background(),
    }
}

// GenerateAndShare generate public and private keys and calculate shared key
func (s *keyService) GenerateAndShare(spKey string) (keys *model.GroupMemberKeys, err error) {
    var pub *ecdsa.PublicKey
    var priv *ecdsa.PrivateKey

    priv, pub, err = ecdh.GenerateKey()
    if err != nil {
        return
    }
    var privs string
    privs, err = ecdh.PrivateKeyEncode(priv)
    if err != nil {
        return
    }

    var shared string
    shared, err = ecdh.ShareKey(spKey, privs)
    if err != nil {
        return
    }

    keys = &model.GroupMemberKeys{
        Public:  ecdh.PublicKeyEncode(pub),
        Private: privs,
        Shared:  shared,
    }
    return
}
