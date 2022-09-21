package service

import (
    "context"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/ent/message"
    "github.com/chatpuppy/puppychat/internal/ent/messageread"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/liasica/go-encryption/aes"
    "github.com/liasica/go-encryption/rsa"
    "time"
)

var (
    // BroadcastChan broadcast to other online members in group
    BroadcastChan = make(chan *model.ChatMessage, 1024)
)

type messageService struct {
    ctx context.Context
    orm *ent.MessageClient
}

func NewMessage() *messageService {
    return &messageService{
        ctx: context.Background(),
        orm: ent.Database.Message,
    }
}

// Create message and broadcast
func (s *messageService) Create(mem *ent.Member, req *model.ChatMessage) (res model.ChatMessageCreateRes, err error) {
    var keys *model.GroupMemberRawKeys
    keys, err = model.LoadSharedKeys(mem.ID, req.GroupID)
    if err != nil {
        return
    }

    // decrypt message
    err = req.Decrypt(keys)
    if err != nil {
        return
    }

    // encrypt message by node keys
    var b []byte
    b, err = rsa.Encrypt([]byte(req.Decrypted), g.RsaPublicKey())

    // save message
    var m *ent.Message
    m, err = s.orm.Create().
        SetGroupID(req.GroupID).
        SetMemberID(req.MemberID).
        SetContent(b).
        SetNillableParentID(req.ParentID).
        Save(s.ctx)
    if err != nil {
        return
    }

    // broadcast message
    req.CreatedAt = m.CreatedAt
    req.ID = m.ID
    BroadcastChan <- req

    res = model.ChatMessageCreateRes{
        ID:        m.ID,
        CreatedAt: m.CreatedAt,
    }

    return
}

// EncryptMessageFromDB decrypt database content and encrypt content by member's share key
func (s *messageService) EncryptMessageFromDB(keys *model.GroupMemberRawKeys, content []byte) (b64 string, err error) {
    var b []byte
    b, err = rsa.Decrypt(content, g.RsaPrivateKey())
    if err != nil {
        return
    }

    return aes.EncryptToBase64(b, keys.Shared)
}

// Detail get message detail
func (s *messageService) Detail(keys *model.GroupMemberRawKeys, msg *ent.Message) (res *model.Message, err error) {
    mem := msg.Edges.Member
    if mem == nil {
        mem, _ = NewMember().QueryID(msg.MemberID)
    }
    if mem == nil {
        err = model.ErrMemberNotFound
    }
    res = &model.Message{
        ID: msg.ID,
        Member: &model.Member{
            ID:       mem.ID,
            Address:  mem.Address,
            Nickname: mem.Nickname,
            Avatar:   mem.Avatar,
        },
        CreatedAt: msg.CreatedAt,
    }

    res.Content, err = s.EncryptMessageFromDB(keys, msg.Content)

    if msg.Edges.Parent != nil {
        res.Quote, err = s.Detail(keys, msg.Edges.Parent)
    }

    return
}

// List get group message list
func (s *messageService) List(mem *ent.Member, req *model.MessageListReq) (res []*model.Message, err error) {
    last := time.Now()
    if req.Time != nil {
        last = *req.Time
    }

    // load shared keys
    var keys *model.GroupMemberRawKeys
    keys, err = model.LoadSharedKeys(mem.ID, req.GroupID)

    items, _ := s.orm.Query().
        Where(
            message.GroupID(req.GroupID),
            message.CreatedAtLT(last),
        ).
        WithMember().
        WithParent(func(query *ent.MessageQuery) {
            query.WithMember()
        }).
        Limit(20).
        All(s.ctx)

    res = make([]*model.Message, len(items))

    for i, item := range items {
        res[i], err = s.Detail(keys, item)
        if err != nil {
            return
        }
    }

    return
}

// Read mask message as read
func (s *messageService) Read(mem *ent.Member, req *model.MessageReadReq) {
    msg, err := s.orm.Query().Where(message.ID(req.ID)).First(s.ctx)
    if err != nil {
        return
    }
    _ = ent.Database.MessageRead.Create().
        SetGroupID(msg.GroupID).
        SetMemberID(mem.ID).
        SetLastTime(msg.CreatedAt).
        SetLastID(msg.ID).
        OnConflictColumns(messageread.FieldGroupID, messageread.FieldMemberID).
        Exec(s.ctx)
}
