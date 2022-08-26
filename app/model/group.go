package model

import (
    "github.com/chatpuppy/puppychat/encryption"
    "github.com/chatpuppy/puppychat/internal/g"
    jsoniter "github.com/json-iterator/go"
)

const (
    // GroupMax Maximum number of groups to create
    GroupMax = 100
)

type GroupCreateReq struct {
    Name  string  `json:"name" validate:"required,min=2,max=20"`
    Intro *string `json:"intro"`
}

type GroupCreateRes struct {
    Name    string  `json:"name"`
    Address string  `json:"address"`
    Intro   *string `json:"intro,omitempty"`
}

type GroupJoinReq struct {
    GroupID uint64 `json:"groupId" validate:"required" trans:"Group ID"`
}

type GroupKeys struct {
    PrivateKey string `json:"privateKey"`
    PublicKey  string `json:"publicKey"`
}

// Marshal group keys to bytes
func (keys *GroupKeys) Marshal() ([]byte, error) {
    return jsoniter.Marshal(keys)
}

// Unmarshal bytes to group keys
func (keys *GroupKeys) Unmarshal(b []byte) (err error) {
    return jsoniter.Unmarshal(b, keys)
}

// Encrypt encrypt group keys
func (keys *GroupKeys) Encrypt() (b []byte, err error) {
    var data []byte
    data, err = keys.Marshal()
    if err != nil {
        return
    }
    return encryption.RSAEncrypt(data, g.RsaPublicKey())
}

// Decrypt decrypt group keys
func (keys *GroupKeys) Decrypt(b []byte) (err error) {
    var data []byte
    data, err = encryption.RSADecrypt(b, g.RsaPrivateKey())
    if err != nil {
        return
    }
    return keys.Unmarshal(data)
}
