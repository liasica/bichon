package model

import (
    "github.com/liasica/go-encryption/aes"
    "time"
)

type Message struct {
    ID        string    `json:"id"`
    CreatedAt time.Time `json:"createdAt"`
    Content   string    `json:"content"`
    GroupID   string    `json:"groupId"`
    Member    *Member   `json:"member,omitempty"`
    Quote     *Message  `json:"quote,omitempty"`

    MessageContent *MessageContent `json:"-"`
}

type MessageCreateReq struct {
    ID       string  `json:"id"`
    MemberID string  `json:"memberId" validate:"required"`
    GroupID  string  `json:"groupId" validate:"required"`
    Content  string  `json:"content" validate:"required"` // encrypted message
    ParentID *string `json:"parentId,omitempty"`          // quote message id
}

type MessageContent struct {
    Encrypted any    `json:"-"` // decrypted message (base64 string or bytes)
    Decrypted []byte `json:"-"` // encrypted message bytes
}

// Decrypt message data using ECDH shared key from member
func (m *MessageContent) Decrypt(keys *GroupMemberRawKeys) (err error) {
    // decrypt data
    var b []byte
    switch encrypted := m.Encrypted.(type) {
    case string:
        b, err = aes.DecryptFromBase64(encrypted, keys.Shared)
    case []byte:
        b, err = aes.Decrypt(encrypted, keys.Shared)
    }

    if err != nil {
        return
    }

    m.Decrypted = b
    return
}

// Encrypt message data using ECDH shared key to member
func (m *MessageContent) Encrypt(keys *GroupMemberRawKeys) (err error) {
    m.Encrypted, err = aes.EncryptToBase64(m.Decrypted, keys.Shared)
    return
}

type ChatMessageCreateRes struct {
    ID        string    `json:"id"`
    CreatedAt time.Time `json:"createdAt"`
}

type MessageListReq struct {
    Time    *time.Time `json:"time,omitempty" query:"time"` // last time
    GroupID string     `json:"groupId" query:"groupId" validate:"required"`
}

type MessageReadReq struct {
    GroupID string `json:"groupId" validate:"required"`
}

type MessageUnread struct {
    Count   int        `json:"count"`
    GroupID string     `json:"groupId"`
    ID      *string    `json:"id,omitempty"`   // first unread message id
    Time    *time.Time `json:"time,omitempty"` // first unread message time
}
