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

// Member holds the schema definition for the Member entity.
type Member struct {
    ent.Schema
}

// Annotations of the Member.
func (Member) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "member"},
    }
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
    return []ent.Field{
        field.String("address").Unique(),
        field.String("nickname").Optional(),
        field.String("avatar").Optional(),
        field.String("intro").Optional(),
        field.String("public_key").Optional(),
        field.String("nonce"),
        field.Bool("show_nickname").Default(true),
    }
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("own_groups", Group.Type),
        edge.To("messages", Message.Type),
        edge.From("groups", Group.Type).Ref("members"),
    }
}

func (Member) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.SonyflakeIDMixin{},
        internal.TimeMixin{},
    }
}

func (Member) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("nonce"),
    }
}
