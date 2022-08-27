package model

import (
    "github.com/chatpuppy/puppychat/internal/g"
    jsoniter "github.com/json-iterator/go"
    "github.com/liasica/go-encryption/rsa"
)

const (
    // GroupMax Maximum number of groups to create
    GroupMax = 100
    // GroupMaxMembers Maximum members of group
    GroupMaxMembers = 1000
)

type GroupCreateReq struct {
    Category   string  `json:"category" validate:"required"` // group's category
    Name       string  `json:"name" validate:"required,min=2,max=20"`
    Intro      *string `json:"intro,omitempty"`  // group intro
    Public     *bool   `json:"public,omitempty"` // `true` create public group, `false` create private group
    MaxMembers int     `json:"maxMembers"`       // group's max members
}

type GroupDetail struct {
    Name       string `json:"name"`
    Category   string `json:"category"`
    Address    string `json:"address"`
    MembersMax int    `json:"membersMax"`
    Intro      string `json:"intro"`
    Public     bool   `json:"public"`          // group is public or private
    Owner      bool   `json:"owner,omitempty"` // group is belongs to current member
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
    return rsa.Encrypt(data, g.RsaPublicKey())
}

// Decrypt decrypt group keys
func (keys *GroupKeys) Decrypt(b []byte) (err error) {
    var data []byte
    data, err = rsa.Decrypt(b, g.RsaPrivateKey())
    if err != nil {
        return
    }
    return keys.Unmarshal(data)
}
