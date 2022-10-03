package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "entgo.io/ent/schema/mixin"
    "github.com/chatpuppy/puppychat/internal/ent/internal"
)

type MessageReadMixin struct {
    mixin.Schema
    Optional     bool
    WithoutIndex bool
}

func (m MessageReadMixin) Fields() []ent.Field {
    f := field.String("message_read_id")
    if m.Optional {
        f.Optional().Nillable()
    }
    return []ent.Field{f}
}

func (m MessageReadMixin) Edges() []ent.Edge {
    e := edge.To("messageRead", MessageRead.Type).Unique().Field("message_read_id")
    if !m.Optional {
        e.Required()
    }
    return []ent.Edge{e}
}

func (m MessageReadMixin) Indexes() (arr []ent.Index) {
    if !m.WithoutIndex {
        arr = append(arr, index.Fields("message_read_id"))
    }
    return
}

// MessageRead holds the schema definition for the MessageRead entity.
type MessageRead struct {
    ent.Schema
}

// Annotations of the MessageRead.
func (MessageRead) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "message_read"},
    }
}

// Fields of the MessageRead.
func (MessageRead) Fields() []ent.Field {
    return []ent.Field{
        field.String("last_id"),
        field.Time("last_time"),
    }
}

// Edges of the MessageRead.
func (MessageRead) Edges() []ent.Edge {
    return []ent.Edge{}
}

func (MessageRead) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.SnowflakeIDMixin{},
        MemberMixin{WithoutIndex: true},
        GroupMixin{WithoutIndex: true},
    }
}

func (MessageRead) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("last_id"),
        index.Fields("last_time"),
        index.Fields("member_id", "group_id").Unique(),
    }
}
