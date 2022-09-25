package model

import (
    "database/sql/driver"
    "fmt"
    "github.com/liasica/go-encryption/hexutil"
    "strconv"
)

type GroupMemberPerm uint8

const (
    GroupMemberPermDefault GroupMemberPerm = 0 // ordinary member permission
    GroupMemberPermManager GroupMemberPerm = 8 // manager member permission
    GroupMemberPermOwner   GroupMemberPerm = 9 // owner
)

func (p *GroupMemberPerm) Scan(v interface{}) (err error) {
    switch x := v.(type) {
    case nil:
    case uint8:
        *p = GroupMemberPerm(x)
    case int64:
        *p = GroupMemberPerm(x)
    case int, int8, int16, int32, uint, uint16, uint32, uint64, uintptr, float32, float64:
        var n uint64
        n, err = strconv.ParseUint(fmt.Sprintf("%x", x), 10, 64)
        if err != nil {
            return
        }
        *p = GroupMemberPerm(n)
    default:
        err = fmt.Errorf("unexpected type %T", v)
    }
    return
}

func (p GroupMemberPerm) Value() (driver.Value, error) {
    return p.Uint8(), nil
}

func (p GroupMemberPerm) Uint8() uint8 {
    return uint8(p)
}

func (p GroupMemberPerm) IsManager() bool {
    return p >= GroupMemberPermManager
}

func (p GroupMemberPerm) IsOwner() bool {
    return p == GroupMemberPermOwner
}

type GroupMemberKeyShareReq struct {
    SharedPublic string `json:"sharedPublic" validate:"required"` // Member's public key for `ecdh` exchange
}

type GroupMemberKeys struct {
    Public  string `json:"public"`
    Private string `json:"private"`
    Shared  string `json:"shared"`
}

func (keys *GroupMemberKeys) Encrypt() (hex string, err error) {
    return KeysEncrypt(keys)
}

func (keys *GroupMemberKeys) Decrypt(hex string) (err error) {
    var decrypted *GroupMemberKeys
    decrypted, err = KeysDecrypt[GroupMemberKeys](hex)
    if err != nil {
        return
    }
    *keys = *decrypted
    return
}

type GroupShareKeyReq struct {
    GroupID string `json:"groupId" validate:"required"` // group id
    *GroupMemberKeyShareReq
}

type GroupMemberRawKeys struct {
    Public  []byte
    Private []byte
    Shared  []byte
}

func (keys *GroupMemberKeys) Raw() *GroupMemberRawKeys {
    pub, _ := hexutil.Decode(keys.Public)
    priv, _ := hexutil.Decode(keys.Private)
    shared, _ := hexutil.Decode(keys.Shared)
    return &GroupMemberRawKeys{
        Public:  pub,
        Private: priv,
        Shared:  shared,
    }
}
