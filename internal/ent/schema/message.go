package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "entgo.io/ent/schema/mixin"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent/internal"
)

type MessageMixin struct {
    mixin.Schema
    Optional     bool
    WithoutIndex bool
}

func (m MessageMixin) Fields() []ent.Field {
    f := field.String("message_id")
    if m.Optional {
        f.Optional().Nillable()
    }
    return []ent.Field{f}
}

func (m MessageMixin) Edges() []ent.Edge {
    e := edge.To("message", Message.Type).Unique().Field("message_id")
    if !m.Optional {
        e.Required()
    }
    return []ent.Edge{e}
}

func (m MessageMixin) Indexes() (arr []ent.Index) {
    if !m.WithoutIndex {
        arr = append(arr, index.Fields("message_id"))
    }
    return
}

// Message holds the schema definition for the Message entity.
type Message struct {
    ent.Schema
}

// Annotations of the Message.
func (Message) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "message"},
    }
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
    return []ent.Field{
        field.String("group_id"),
        field.String("member_id"),
        field.Bytes("content"),
        field.String("parent_id").Optional().Nillable(),
        field.JSON("owner", &model.Member{}).Comment("message's owner"),
        field.Int64("last_node"),
    }
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("member", Member.Type).Ref("messages").Required().Unique().Field("member_id"),
        edge.From("group", Group.Type).Ref("messages").Required().Unique().Field("group_id"),
        edge.To("children", Message.Type).From("parent").Unique().Field("parent_id"),
    }
}

func (Message) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.SnowflakeIDMixin{},
        internal.TimeMixin{},
    }
}

func (Message) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("group_id"),
        index.Fields("member_id"),
        index.Fields("parent_id"),
    }
}

func (Message) Hooks() []ent.Hook {
    return []ent.Hook{
        internal.DistributeHook(),
    }
}
