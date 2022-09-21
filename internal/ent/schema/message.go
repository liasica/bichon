package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "github.com/chatpuppy/puppychat/internal/ent/internal"
)

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
