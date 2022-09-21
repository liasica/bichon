package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "github.com/chatpuppy/puppychat/internal/ent/internal"
)

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
        MemberMixin{},
        GroupMixin{},
    }
}

func (MessageRead) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("last_id"),
        index.Fields("last_time"),
    }
}
