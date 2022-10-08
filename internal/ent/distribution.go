package ent

import (
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent/group"
    "github.com/chatpuppy/puppychat/internal/ent/groupmember"
    "github.com/chatpuppy/puppychat/internal/ent/key"
    "github.com/chatpuppy/puppychat/internal/ent/member"
    "github.com/chatpuppy/puppychat/internal/ent/message"
    "github.com/chatpuppy/puppychat/internal/g"
    jsoniter "github.com/json-iterator/go"
    rsaTools "github.com/liasica/go-encryption/rsa"
    log "github.com/sirupsen/logrus"
)

func toJsonBytes(v any, t string) []byte {
    b, err := jsoniter.Marshal(v)
    if err != nil {
        log.Errorf("%s to json bytes falied:%v", t, err)
    }
    return b
}

func (m *GroupMutation) Parse() *model.SyncData {
    data := m.SyncData()
    if data == nil {
        return nil
    }

    op := m.op
    if !op.Is(OpDelete | OpDeleteOne) {
        if data.Keys != "" {
            keys := new(model.GroupKeys)
            _ = keys.Decrypt(data.Keys)
            b, _ := jsoniter.Marshal(keys)
            data.Keys = string(b)
        }
    }

    return &model.SyncData{
        Table: group.Table,
        Value: toJsonBytes(data, group.Table),
        Op:    uint(op),
    }
}

func (m *GroupMutation) Columns() []string {
    return group.Columns
}

func (m *GroupMemberMutation) Parse() *model.SyncData {
    data := m.SyncData()
    op := m.op
    if data == nil {
        return nil
    }
    return &model.SyncData{
        Table: groupmember.Table,
        Value: toJsonBytes(data, groupmember.Table),
        Op:    uint(op),
    }
}

func (m *GroupMemberMutation) Columns() []string {
    return groupmember.Columns
}

func (m *KeyMutation) Parse() *model.SyncData {
    data := m.SyncData()
    if data == nil {
        return nil
    }

    op := m.op
    if !op.Is(OpDelete | OpDeleteOne) {
        if data.Keys != "" {
            keys := new(model.GroupMemberKeys)
            _ = keys.Decrypt(data.Keys)
            b, _ := jsoniter.Marshal(keys)
            data.Keys = string(b)
        }
    }

    return &model.SyncData{
        Table: key.Table,
        Value: toJsonBytes(data, key.Table),
        Op:    uint(op),
    }
}

func (m *KeyMutation) Columns() []string {
    return key.Columns
}

func (m *MemberMutation) Parse() *model.SyncData {
    data := m.SyncData()
    if data == nil {
        return nil
    }

    op := m.op
    return &model.SyncData{
        Table: member.Table,
        Value: toJsonBytes(data, member.Table),
        Op:    uint(op),
    }
}

func (m *MemberMutation) Columns() []string {
    return member.Columns
}

func (m *MessageMutation) Parse() *model.SyncData {
    data := m.SyncData()
    if data == nil {
        return nil
    }

    op := m.op
    if !op.Is(OpDelete | OpDeleteOne) {
        if data.Content != nil {
            content, _ := rsaTools.DecryptUsePrivateKey(data.Content, g.RsaPrivateKey())
            data.Content = content
        }
    }

    return &model.SyncData{
        Table: message.Table,
        Value: toJsonBytes(data, message.Table),
        Op:    uint(op),
    }
}

func (m *MessageMutation) Columns() []string {
    return message.Columns
}
