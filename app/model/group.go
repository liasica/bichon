package model

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
    *GroupMemberKeyShareReq
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

// Encrypt encrypt group keys
func (keys *GroupKeys) Encrypt() (hex string, err error) {
    return KeysEncrypt(keys)
}

// Decrypt decrypt group keys
func (keys *GroupKeys) Decrypt(hex string) (err error) {
    var decrypted *GroupKeys
    decrypted, err = KeysDecrypt[GroupKeys](hex)
    if err != nil {
        return
    }
    *keys = *decrypted
    return
}
