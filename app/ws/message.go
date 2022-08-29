package ws

import (
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/service"
    "github.com/liasica/go-encryption/aes"
)

type Action uint8

const (
    ActionEnter Action = iota + 1 // enter group
    ActionLeave                   // leave group
    ActionText                    // text message
)

type Message[T MessageType] struct {
    MemberID string `json:"memberId"`
    GroupID  string `json:"groupId"`
    KeyID    string `json:"keyId"`  // group key id
    Action   Action `json:"action"` // actions, 1: enter group, 2: leave group, 3: text message
    // encrypted message
    Encrypted string `json:"encrypted,omitempty"`
    // decrypted message
    Data T `json:"data,omitempty"`

    client *Client
}

type MessageType interface {
    *TextMessage

    Marshal() ([]byte, error)
    Unmarshal([]byte) error
}

// TextMessage text message
type TextMessage string

func NewTextMessage(message string) *TextMessage {
    m := TextMessage(message)
    return &m
}

func (m *TextMessage) Marshal() ([]byte, error) {
    return []byte(*m), nil
}

func (m *TextMessage) Unmarshal(b []byte) error {
    *m = TextMessage(b)
    return nil
}

// SetClient setting message's client
func (m *Message[T]) SetClient(client *Client) {
    m.client = client
}

// GetKeys getting keys from cache or database
func (m *Message[T]) GetKeys() (keys *model.GroupMemberRawKeys, err error) {
    v, ok := m.client.hub.keys.Load(m.KeyID)
    if ok {
        keys = v.(*model.GroupMemberRawKeys)
    } else {
        keys, err = service.NewKey().QueryRawKeys(m.KeyID, m.MemberID, m.GroupID)
        if err != nil {
            return
        }
    }
    return
}

// DecryptDataFromMember decrypt message data using ECDH shared key from member
func (m *Message[T]) DecryptDataFromMember() error {
    keys, err := m.GetKeys()
    if err != nil {
        return err
    }

    // decrypt data
    var b []byte
    b, err = aes.DecryptFromBase64(m.Encrypted, keys.Shared)
    if err != nil {
        return err
    }

    return m.Data.Unmarshal(b)
}

// EncryptDataToMember encrypt message data using ECDH shared key to member
func (m *Message[T]) EncryptDataToMember() error {
    keys, err := m.GetKeys()
    if err != nil {
        return err
    }
    var data []byte
    data, err = m.Data.Marshal()
    if err != nil {
        return err
    }
    m.Encrypted, err = aes.EncryptToBase64(data, keys.Shared)
    return err
}
