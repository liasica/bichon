package model

import (
    "fmt"
    "golang.org/x/exp/slices"
    "sync"
    "time"
)

const (
    // GroupMax Maximum number of groups to create
    GroupMax = 100
    // GroupMaxMembers Maximum members of group
    GroupMaxMembers = 1000
    // GroupCreateFrequency create group interval seconds
    GroupCreateFrequency = 60
    // GroupKeyShareFrequency group key share interval seconds
    GroupKeyShareFrequency = 60
    // GroupInviteCodeExpires group invite expiration days
    GroupInviteCodeExpires = 7 // 7 days
)

type GroupCategory string

var (
    GroupCategories = []GroupCategory{"TEST", "test1", "test2", "test3"}

    keysCache sync.Map

    // GroupActived member current actived group
    // type is memberID => groupID
    GroupActived sync.Map
)

func (cate GroupCategory) IsValid() bool {
    return slices.Contains(GroupCategories, cate)
}

func StoreSharedKeys(memberID, groupID string, keys *GroupMemberRawKeys) {
    keysCache.Store(fmt.Sprintf("%s-%s", memberID, groupID), keys)
}

func LoadSharedKeys(memberID, groupID string) (keys *GroupMemberRawKeys, err error) {
    // load shared keys
    k, _ := keysCache.Load(fmt.Sprintf("%s-%s", memberID, groupID))
    var ok bool
    keys, ok = k.(*GroupMemberRawKeys)
    if !ok {
        err = ErrShareKeyNotFound
        return
    }
    return
}

type GroupCreateReq struct {
    Category   GroupCategory `json:"category" validate:"required,enum"` // group's category
    Name       string        `json:"name" validate:"required,min=2,max=20"`
    Intro      *string       `json:"intro,omitempty"`  // group intro
    Public     *bool         `json:"public,omitempty"` // `true` create public group, `false` create private group
    MaxMembers int           `json:"maxMembers"`       // group's max members
    *GroupMemberKeyShareReq
}

type GroupPublicKey struct {
    GroupPublicKey string `json:"groupPublicKey"`
    KeyID          string `json:"keyId"`
}

type GroupDetailWithPublicKey struct {
    *GroupPublicKey
    *GroupDetail
}

type GroupMeta struct {
    ID           string `json:"id"`
    Name         string `json:"name"`
    Address      string `json:"address"`
    MembersMax   int    `json:"membersMax"`
    MembersCount int    `json:"membersCount"`
    Intro        string `json:"intro,omitempty"` // optional
    Category     string `json:"category"`
}

type GroupJoinedListRes struct {
    GroupMeta
    UnreadCount int        `json:"unreadCount"`          // unread message count
    UnreadID    *string    `json:"unreadId,omitempty"`   // first unread message id
    UnreadTime  *time.Time `json:"unreadTime,omitempty"` // first unread message time
}

type GroupDetail struct {
    GroupMeta
    Public     bool   `json:"public"`          // group is public or private
    Owner      bool   `json:"owner,omitempty"` // group is belongs to current member
    InviteCode string `json:"inviteCode"`
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

type GroupJoinReq struct {
    GroupID    string `json:"groupId,omitempty" validate:"required_without=InviteCode"`
    InviteCode string `json:"inviteCode,omitempty" validate:"required_without=GroupID"` // invite code from group member
    *GroupMemberKeyShareReq
}

type GroupMetaReq struct {
    ID string `json:"id" validate:"required" param:"id"`
}

type GroupDetailRes struct {
    *GroupDetail
    Members []MemberWithPermission `json:"members"`
}

type GroupKeyUsed struct {
    GroupID  string              `json:"groupId" validate:"required"`
    KeyID    string              `json:"keyId" validate:"required"`
    MemberID string              `json:"memberId" swaggerignore:"true"`
    Keys     *GroupMemberRawKeys `json:"keys" swaggerignore:"true"`
}

type GroupKeyUsedReq struct {
    Keys []*GroupKeyUsed `json:"keys" validate:"required,min=1"`
}

type GroupListReq struct {
    PaginationReq
    Name     string        `json:"name" query:"name"`
    Category GroupCategory `json:"category" query:"category"`
}

type GroupListRes struct {
    GroupMeta
    Joined bool `json:"joined,omitempty"`
}

type GroupIDReq struct {
    GroupID string `json:"groupId" validate:"required"`
}

type GroupUpdateReq struct {
    GroupID    string         `json:"groupId" validate:"required"`
    Category   *GroupCategory `json:"category"`
    Name       *string        `json:"name"`
    Intro      *string        `json:"intro"`
    Public     *bool          `json:"public"`
    MaxMembers *int           `json:"maxMembers"`
}

type GroupInviteCodeRes struct {
    InviteCode string `json:"inviteCode"`
}

type GroupActiveData struct {
    GroupID  string `json:"groupId"`
    MemberID string `json:"memberId"`
}

type GroupMemberReq struct {
    GroupID  string `json:"groupId"`
    MemberID string `json:"memberId"`
}
