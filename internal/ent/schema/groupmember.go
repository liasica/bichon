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
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent/internal"
)

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
    return []ent.Index{}
}

type GroupMemberMutator interface {
    Cache(delete bool)
}

func (GroupMember) Hooks() []ent.Hook {
    return []ent.Hook{
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
