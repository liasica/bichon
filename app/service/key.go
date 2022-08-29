package service

import (
    "context"
    "crypto/ecdsa"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/ent/key"
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

// QueryKeys get group key
func (s *keyService) QueryKeys(id, memberID, groupID string) (keys *model.GroupMemberKeys, err error) {
    k, _ := ent.Database.Key.Query().Where(
        key.ID(id),
        key.MemberID(memberID),
        key.GroupID(groupID),
    ).First(s.ctx)
    if k == nil {
        err = model.ErrGroupKey
        return
    }
    // decrypt keys use node's key
    keys = new(model.GroupMemberKeys)
    err = keys.Decrypt(k.Keys)
    if err != nil {
        return
    }
    return
}

func (s *keyService) QueryRawKeys(id, memberID, groupID string) (keys *model.GroupMemberRawKeys, err error) {
    var hexKeys *model.GroupMemberKeys
    hexKeys, err = s.QueryKeys(id, memberID, groupID)
    if err != nil {
        return
    }
    keys = hexKeys.Raw()
    return
}
