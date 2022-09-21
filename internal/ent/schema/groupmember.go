package schema

import (
    "context"
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/entc/integration/ent/hook"
    "entgo.io/ent/schema"
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
        field.Uint8("permission").Default(model.GroupMemberPermDefault).Comment("member's permission in group"),
        field.String("sn").Unique().Comment("user's share sn"),
    }
}

// Edges of the GroupMember.
func (GroupMember) Edges() []ent.Edge {
    return []ent.Edge{}
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
        // index.Fields("group_id", "member_id"),
    }
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
