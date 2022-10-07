package schema

import (
    "ariga.io/atlas/sql/postgres"
    "context"
    "entgo.io/ent"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/entc/integration/ent/hook"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "entgo.io/ent/schema/mixin"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent/internal"
)

type GroupMemberMixin struct {
    mixin.Schema
    Optional     bool
    WithoutIndex bool
}

func (m GroupMemberMixin) Fields() []ent.Field {
    f := field.String("group_member_id")
    if m.Optional {
        f.Optional().Nillable()
    }
    return []ent.Field{f}
}

func (m GroupMemberMixin) Edges() []ent.Edge {
    e := edge.To("groupMember", GroupMember.Type).Unique().Field("group_member_id")
    if !m.Optional {
        e.Required()
    }
    return []ent.Edge{e}
}

func (m GroupMemberMixin) Indexes() (arr []ent.Index) {
    if !m.WithoutIndex {
        arr = append(arr, index.Fields("group_member_id"))
    }
    return
}

// GroupMember holds the schema definition for the GroupMember entity.
type GroupMember struct {
    ent.Schema
}

// Annotations of the GroupMember.
func (GroupMember) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "group_member"},
    }
}

// Fields of the GroupMember.
func (GroupMember) Fields() []ent.Field {
    return []ent.Field{
        field.Other("permission", model.GroupMemberPerm(0)).Default(model.GroupMemberPermDefault).SchemaType(map[string]string{
            dialect.Postgres: postgres.TypeSmallInt,
        }).Comment("member's permission in group"),
        field.String("inviter_id").Optional().Nillable(),
        field.String("invite_code").Unique().Comment("invite code"),
        field.Time("invite_expire").Comment("invite code expire time"),
        field.String("read_id").Optional().Nillable().Comment("last read message id"),
        field.Time("read_time").Optional().Nillable().Comment("last read message time"),
        field.Int64("last_node"),
    }
}

// Edges of the GroupMember.
func (GroupMember) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("inviter", Member.Type).Unique().Field("inviter_id"),
    }
}

func (GroupMember) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.SnowflakeIDMixin{},
        internal.TimeMixin{},
        MemberMixin{},
        GroupMixin{},
    }
}

func (GroupMember) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("read_id"),
        index.Fields("read_time"),
    }
}

type GroupMemberMutator interface {
    Cache(delete bool)
}

func (GroupMember) Hooks() []ent.Hook {
    return []ent.Hook{
        internal.DistributeHook(),
        hook.On(
            func(next ent.Mutator) ent.Mutator {
                return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
                    m, ok := mutation.(GroupMemberMutator)
                    if ok {
                        m.Cache(mutation.Op().Is(ent.OpDelete | ent.OpDeleteOne))
                    }
                    return next.Mutate(ctx, mutation)
                })
            },
            ent.OpCreate|ent.OpDelete,
        ),
    }
}
