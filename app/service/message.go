package service

import (
    "context"
    "entgo.io/ent/dialect/sql"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/ent/group"
    "github.com/chatpuppy/puppychat/internal/ent/member"
    "github.com/chatpuppy/puppychat/internal/ent/message"
    "github.com/chatpuppy/puppychat/internal/ent/messageread"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/liasica/go-encryption/aes"
    "github.com/liasica/go-encryption/rsa"
    "time"
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
func (s *messageService) Create(mem *ent.Member, req *model.MessageCreateReq) (res *model.Message, err error) {
    // find group
    gro, _ := ent.Database.Group.Query().Where(group.HasMembersWith(member.ID(mem.ID))).First(s.ctx)
    if gro == nil {
        err = model.ErrNotFoundGroup
        return
    }
    // owner
    owner := &model.Member{
        ID:       mem.ID,
        Address:  mem.Address,
        Nickname: mem.Nickname,
        Avatar:   mem.Avatar,
    }

    var keys *model.GroupMemberRawKeys
    keys, err = model.LoadSharedKeys(mem.ID, req.GroupID)
    if err != nil {
        return
    }

    mc := &model.MessageContent{
        Encrypted: req.Content,
    }

    // decrypt message
    err = mc.Decrypt(keys)
    if err != nil {
        return
    }

    // encrypt message by node keys
    var b []byte
    b, err = rsa.Encrypt(mc.Decrypted, g.RsaPublicKey())

    // get quote message
    var (
        parent *ent.Message
        bq     *model.Message
        rq     *model.Message
    )
    if req.ParentID != nil {
        parent, _ = s.orm.Query().Where(message.ID(*req.ParentID), message.GroupID(req.GroupID)).First(s.ctx)
        if parent != nil {
            err = model.ErrQuoteMessageNotFound
            return
        }

        // decrypt parent message content
        var pb []byte
        pb, err = rsa.Decrypt(parent.Content, g.RsaPrivateKey())
        if err != nil {
            return
        }

        // quote message
        quote := model.Message{
            ID:        parent.ID,
            CreatedAt: parent.CreatedAt,
            Member:    parent.Owner,

            MessageContent: &model.MessageContent{Decrypted: pb},
        }
        bq = &quote
        rq = &quote

        err = s.Share(keys, rq)
        if err != nil {
            return
        }
    }

    // save message
    var m *ent.Message
    m, err = s.orm.Create().
        SetGroupID(req.GroupID).
        SetMemberID(req.MemberID).
        SetContent(b).
        SetNillableParentID(req.ParentID).
        SetOwner(owner).
        Save(s.ctx)
    if err != nil {
        return
    }

    // broadcast message
    model.SendBroadcast(model.MessageBroadcast{
        Message: model.Message{
            ID:        m.ID,
            CreatedAt: m.CreatedAt,
            GroupID:   req.GroupID,
            Member:    owner,
            Quote:     bq,

            MessageContent: &model.MessageContent{Decrypted: mc.Decrypted},
        },
        GroupAddress: gro.Address,
    })

    res = &model.Message{
        Content:   req.Content,
        GroupID:   req.GroupID,
        Member:    owner,
        ID:        m.ID,
        CreatedAt: m.CreatedAt,
        Quote:     rq,
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

// Share message content encrypt use sharekey
func (s *messageService) Share(keys *model.GroupMemberRawKeys, msg *model.Message) (err error) {
    err = msg.MessageContent.Encrypt(keys)
    if err != nil {
        return
    }
    msg.Content = msg.MessageContent.Encrypted.(string)

    if msg.Quote != nil {
        err = s.Share(keys, msg.Quote)
    }

    return
}

// Detail get message detail
func (s *messageService) Detail(keys *model.GroupMemberRawKeys, msg *ent.Message) (res *model.Message, err error) {
    res = &model.Message{
        ID:        msg.ID,
        GroupID:   msg.GroupID,
        Member:    msg.Owner,
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

    if keys == nil {
        err = model.ErrShareKeyNotFound
        return
    }

    q := s.orm.Query().
        Where(
            message.GroupID(req.GroupID),
            message.CreatedAtLT(last),
        ).
        WithParent().
        Order(ent.Desc(message.FieldCreatedAt)).
        Limit(20)

    items, _ := q.All(s.ctx)

    res = make([]*model.Message, len(items))

    for i, item := range items {
        res[i], err = s.Detail(keys, item)
        if err != nil {
            return
        }
    }

    return
}

// Read mask group message as read
func (s *messageService) Read(mem *ent.Member, req *model.MessageReadReq) error {
    // is member in group
    if !NewGroup().MemberIn(mem.ID, req.GroupID) {
        return model.ErrNotInGroup
    }
    msg, err := s.orm.Query().Where(message.GroupID(req.GroupID)).Order(ent.Desc(message.FieldCreatedAt)).First(s.ctx)
    if err != nil {
        return err
    }
    return s.UpdateRead(msg.ID, mem.ID, req.GroupID, msg.CreatedAt)
}

// UpdateRead update read message attributes
func (s *messageService) UpdateRead(id, memID, groupID string, createdAt time.Time) error {
    return ent.Database.MessageRead.Create().
        SetGroupID(groupID).
        SetMemberID(memID).
        SetLastTime(createdAt).
        SetLastID(id).
        OnConflictColumns(messageread.FieldGroupID, messageread.FieldMemberID).
        Exec(s.ctx)
}

// UnreadList group message unread info
func (s *messageService) UnreadList(memberID string) (items map[string]model.MessageUnread) {
    var result []struct {
        Count   int        `json:"count"`
        GroupID string     `json:"group_id"`
        ID      *string    `json:"id"`   // first unread message id
        Time    *time.Time `json:"time"` // first unread message time
    }
    _ = s.orm.Query().Modify(func(sel *sql.Selector) {
        t := sql.Table(messageread.Table)
        sel.LeftJoin(t).On(t.C(messageread.FieldGroupID), sel.C(message.FieldGroupID)).
            Where(
                sql.And(
                    sql.EQ(sel.C(message.FieldMemberID), memberID),
                    sql.Or(
                        sql.ColumnsGT(sel.C(message.FieldCreatedAt), t.C(messageread.FieldLastTime)),
                        sql.IsNull(t.C(messageread.FieldLastTime)),
                    ),
                ),
            ).
            Select(
                sel.C(message.FieldGroupID),
                sql.As(sql.Count(sel.C(message.FieldID)), "count"),
                sql.As(sql.Min(sel.C(message.FieldID)), "id"),
                sql.As(sql.Min(sel.C(message.FieldCreatedAt)), "time"),
            ).
            GroupBy(sel.C(message.FieldGroupID))
    }).Scan(s.ctx, &result)
    items = make(map[string]model.MessageUnread)
    for _, r := range result {
        items[r.GroupID] = model.MessageUnread{
            Count:   r.Count,
            GroupID: r.GroupID,
            ID:      r.ID,
            Time:    r.Time,
        }
    }
    return
}

// AutoRead auto update read message attributes from actived group
func (s *messageService) AutoRead(msg *model.Message) {
    groupID, ok := model.GroupActived.Load(msg.Member.ID)
    if !ok {
        return
    }
    if groupID != msg.GroupID {
        return
    }
    _ = s.UpdateRead(msg.ID, msg.Member.ID, msg.GroupID, msg.CreatedAt)
}
