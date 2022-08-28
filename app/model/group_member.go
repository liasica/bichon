package model

const (
    GroupMemberPermDefault uint8 = 0 // ordinary member permission
    GroupMemberPermManager uint8 = 8 // manager member permission
    GroupMemberPermOwner   uint8 = 9 // owner
)

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
