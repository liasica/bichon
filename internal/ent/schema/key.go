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

type KeyMixin struct {
    mixin.Schema
    Optional bool
}

func (m KeyMixin) Fields() []ent.Field {
    f := field.Uint64("key_id")
    if m.Optional {
        f.Optional().Nillable()
    }
    return []ent.Field{f}
}

func (m KeyMixin) Edges() []ent.Edge {
    e := edge.To("key", Key.Type).Unique().Field("key_id")
    if !m.Optional {
        e.Required()
    }
    return []ent.Edge{e}
}

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
        internal.SonyflakeIDMixin{},
        internal.TimeMixin{},
    }
}

func (Key) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("group_id", "member_id"),
        index.Fields("enable"),
    }
}
