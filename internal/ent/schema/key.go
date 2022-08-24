package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "github.com/chatpuppy/puppychat/internal/ent/internal"
)

// Key holds the schema definition for the Key entity.
type Key struct {
    ent.Schema
}

// Annotations of the Key.
func (Key) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "key"},
    }
}

// Fields of the Key.
func (Key) Fields() []ent.Field {
    return []ent.Field{
        field.Uint64("group_id"),
        field.Uint64("member_id"),
        field.String("key"),
        field.Bool("enable").Default(true),
    }
}

// Edges of the Key.
func (Key) Edges() []ent.Edge {
    return []ent.Edge{}
}

func (Key) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.TimeMixin{},
    }
}

func (Key) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("group_id", "member_id"),
        index.Fields("enable"),
    }
}
