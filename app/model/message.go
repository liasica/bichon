package model

import (
    "github.com/liasica/go-encryption/aes"
    "time"
)

type ChatMessage struct {
    ID        string    `json:"id"`
    MemberID  string    `json:"memberId" validate:"required"`
    GroupID   string    `json:"groupId" validate:"required"`
    Content   string    `json:"content" validate:"required"` // encrypted message
    ParentID  *string   `json:"parentId,omitempty"`          // quote message id
    CreatedAt time.Time `json:"createdAt"`

    Decrypted string `json:"-"` // decrypted message
}

// Decrypt message data using ECDH shared key from member
func (m *ChatMessage) Decrypt(keys *GroupMemberRawKeys) (err error) {
    // decrypt data
    var b []byte
    b, err = aes.DecryptFromBase64(m.Content, keys.Shared)
    if err != nil {
        return
    }

    m.Decrypted = string(b)
    return
}

// Encrypt message data using ECDH shared key to member
func (m *ChatMessage) Encrypt(keys *GroupMemberRawKeys) (err error) {
    m.Content, err = aes.EncryptToBase64([]byte(m.Decrypted), keys.Shared)
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

type Message struct {
    ID        string    `json:"id"`
    CreatedAt time.Time `json:"createdAt"`
    Content   string    `json:"content"`
    Member    *Member   `json:"member,omitempty"`
    Quote     *Message  `json:"quote,omitempty"`
}

type MessageReadReq struct {
    ID string `json:"id"`
}
