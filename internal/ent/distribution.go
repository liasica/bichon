package ent

import (
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent/group"
    "github.com/chatpuppy/puppychat/internal/ent/groupmember"
    "github.com/chatpuppy/puppychat/internal/ent/key"
    "github.com/chatpuppy/puppychat/internal/ent/member"
    "github.com/chatpuppy/puppychat/internal/ent/message"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/chatpuppy/puppychat/pkg/tea"
    jsoniter "github.com/json-iterator/go"
    rsaTools "github.com/liasica/go-encryption/rsa"
)

func toJsonBytes(v any) []byte {
    b, _ := jsoniter.Marshal(v)
    return b
}

func (m *GroupMutation) Parse(op Op) *model.SyncData {
    data := m.SyncData()
    if data == nil {
        return nil
    }
    if !op.Is(OpDelete | OpDeleteOne) {
        if data.Keys != nil {
            keys := new(model.GroupKeys)
            _ = keys.Decrypt(*data.Keys)
            b, _ := jsoniter.Marshal(keys)
            data.Keys = tea.String(string(b))
        }
    }

    return &model.SyncData{
        Table: group.Table,
        Value: toJsonBytes(data),
        Op:    uint(op),
    }
}

func (m *GroupMutation) Columns() []string {
    return group.Columns
}

func (m *GroupMemberMutation) Parse(op Op) *model.SyncData {
    data := m.SyncData()
    if data == nil {
        return nil
    }
    return &model.SyncData{
        Table: groupmember.Table,
        Value: toJsonBytes(data),
        Op:    uint(op),
    }
}

func (m *GroupMemberMutation) Columns() []string {
    return groupmember.Columns
}

func (m *KeyMutation) Parse(op Op) *model.SyncData {
    data := m.SyncData()
    if data == nil {
        return nil
    }

    if !op.Is(OpDelete | OpDeleteOne) {
        if data.Keys != nil {
            keys := new(model.GroupMemberKeys)
            _ = keys.Decrypt(*data.Keys)
            b, _ := jsoniter.Marshal(keys)
            data.Keys = tea.String(string(b))
        }
    }

    return &model.SyncData{
        Table: key.Table,
        Value: toJsonBytes(data),
        Op:    uint(op),
    }
}

func (m *KeyMutation) Columns() []string {
    return key.Columns
}

func (m *MemberMutation) Parse(op Op) *model.SyncData {
    data := m.SyncData()
    if data == nil {
        return nil
    }

    return &model.SyncData{
        Table: member.Table,
        Value: toJsonBytes(m.SyncData()),
        Op:    uint(op),
    }
}

func (m *MemberMutation) Columns() []string {
    return member.Columns
}

func (m *MessageMutation) Parse(op Op) *model.SyncData {
    data := m.SyncData()
    if data == nil {
        return nil
    }

    if !op.Is(OpDelete | OpDeleteOne) {
        if data.Content != nil {
            content, _ := rsaTools.DecryptUsePrivateKey(*data.Content, g.RsaPrivateKey())
            data.Content = tea.Pointer(content)
        }
    }

    return &model.SyncData{
        Table: message.Table,
        Value: toJsonBytes(data),
        Op:    uint(op),
    }
}

func (m *MessageMutation) Columns() []string {
    return message.Columns
}
