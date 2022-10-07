package internal

import (
    "context"
    "entgo.io/ent"
    "entgo.io/ent/entc/integration/ent/hook"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/mixin"
    "fmt"
    "github.com/chatpuppy/puppychat/internal/g"
)

type SnowflakeIDMixin struct {
    mixin.Schema
}

func (SnowflakeIDMixin) Fields() []ent.Field {
    return []ent.Field{
        field.String("id").MaxLen(25).NotEmpty().Unique().Immutable(),
    }
}

// Hooks of the Mixin.
func (SnowflakeIDMixin) Hooks() []ent.Hook {
    return []ent.Hook{
        hook.On(IDHook(), ent.OpCreate),
    }
}

type IDSetter interface {
    SetID(string)
    ID() (string, bool)
}

func IDHook() ent.Hook {
    return func(next ent.Mutator) ent.Mutator {
        return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
            is, ok := m.(IDSetter)
            if !ok {
                return nil, fmt.Errorf("unexpected mutation %T", m)
            }
            _, ok = is.ID()
            if !ok {
                id := g.SnowflakeNode().Generate()
                is.SetID(fmt.Sprintf("%d", id))
            }
            return next.Mutate(ctx, m)
        })
    }
}
