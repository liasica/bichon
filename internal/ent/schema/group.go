package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "github.com/chatpuppy/puppychat/internal/ent/internal"
)

// Group holds the schema definition for the Group entity.
type Group struct {
    ent.Schema
}

// Annotations of the Group.
func (Group) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "group"},
    }
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
    return []ent.Field{
        field.String("name"),
        field.Uint64("member_id").Comment("created by"),
        field.Int("members_max"),
        field.Int("members_count"),
        field.Bool("public"),
        field.String("sn"),
        field.String("intro"),
    }
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("owner", Member.Type).Ref("own_groups").Required().Unique().Field("member_id"),
        edge.To("messages", Message.Type),
        edge.To("members", Member.Type),
    }
}

func (Group) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.TimeMixin{},
    }
}

func (Group) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("name").Annotations(
            entsql.IndexTypes(map[string]string{
                dialect.Postgres: "GIN",
            }),
        ),
        index.Fields("member_id"),
        index.Fields("sn"),
    }
}
