package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "entgo.io/ent/schema/mixin"
    "github.com/chatpuppy/puppychat/internal/ent/internal"
)

type GroupMixin struct {
    mixin.Schema
    Optional     bool
    WithoutIndex bool
}

func (m GroupMixin) Fields() []ent.Field {
    f := field.String("group_id")
    if m.Optional {
        f.Optional().Nillable()
    }
    return []ent.Field{f}
}

func (m GroupMixin) Edges() []ent.Edge {
    e := edge.To("group", Group.Type).Unique().Field("group_id")
    if !m.Optional {
        e.Required()
    }
    return []ent.Edge{e}
}

func (m GroupMixin) Indexes() (arr []ent.Index) {
    if !m.WithoutIndex {
        arr = append(arr, index.Fields("group_id"))
    }
    return
}

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
        field.String("category"),
        field.String("owner_id").Comment("created by"),
        field.Int("members_max"),
        field.Int("members_count").Default(1).Comment("members count of group"),
        field.Bool("public"),
        field.String("address").Immutable().Unique(),
        field.String("intro").Optional(),
        field.Text("keys").Comment("group's ethereum keys"),
    }
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("owner", Member.Type).Ref("own_groups").Required().Unique().Field("owner_id"),
        edge.To("messages", Message.Type),

        edge.To("members", Member.Type).Through("group_members", GroupMember.Type),
    }
}

func (Group) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.SnowflakeIDMixin{},
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
        index.Fields("owner_id"),
    }
}

func (Group) Hooks() []ent.Hook {
    return []ent.Hook{
        internal.DistributeHook(),
    }
}
