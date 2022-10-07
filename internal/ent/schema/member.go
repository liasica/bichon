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

type MemberMixin struct {
    mixin.Schema
    Optional     bool
    WithoutIndex bool
}

func (m MemberMixin) Fields() []ent.Field {
    f := field.String("member_id")
    if m.Optional {
        f.Optional().Nillable()
    }
    return []ent.Field{f}
}

func (m MemberMixin) Edges() []ent.Edge {
    e := edge.To("member", Member.Type).Unique().Field("member_id")
    if !m.Optional {
        e.Required()
    }
    return []ent.Edge{e}
}

func (m MemberMixin) Indexes() (arr []ent.Index) {
    if !m.WithoutIndex {
        arr = append(arr, index.Fields("member_id"))
    }
    return
}

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
        field.Text("public_key").Optional(),
        field.String("nonce"),
        field.Bool("show_nickname").Default(true),
        field.Int64("last_node"),
    }
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("own_groups", Group.Type),
        edge.To("messages", Message.Type),
        edge.From("groups", Group.Type).Ref("members").Through("group_members", GroupMember.Type),
    }
}

func (Member) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.SnowflakeIDMixin{},
        internal.TimeMixin{},
    }
}

func (Member) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("nonce"),
    }
}

func (Member) Hooks() []ent.Hook {
    return []ent.Hook{
        internal.DistributeHook(),
    }
}
