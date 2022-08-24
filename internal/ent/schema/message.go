package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
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
        field.Uint64("group_id"),
        field.Uint64("member_id"),
        field.Text("content"),
    }
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("owner", Member.Type).Ref("messages").Required().Unique().Field("member_id"),
        edge.From("group", Group.Type).Ref("messages").Required().Unique().Field("group_id"),
    }
}

func (Message) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.TimeMixin{},
    }
}

func (Message) Indexes() []ent.Index {
    return []ent.Index{}
}
