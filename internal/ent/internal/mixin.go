package internal

import (
    "context"
    "entgo.io/ent"
    "entgo.io/ent/entc/integration/ent/hook"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/mixin"
    "fmt"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/sony/sonyflake"
)

type SonyflakeIDMixin struct {
    mixin.Schema
}

func (SonyflakeIDMixin) Fields() []ent.Field {
    return []ent.Field{
        field.String("id").MaxLen(25).NotEmpty().Unique().Immutable(),
    }
}

// Hooks of the Mixin.
func (SonyflakeIDMixin) Hooks() []ent.Hook {
    return []ent.Hook{
        hook.On(IDHook(), ent.OpCreate),
    }
}

type IDSetter interface {
    SetID(string)
}

func IDHook() ent.Hook {
    return func(next ent.Mutator) ent.Mutator {
        sf := sonyflake.NewSonyflake(sonyflake.Settings{
            MachineID: func() (uint16, error) {
                return g.NodeID(), nil
            },
            CheckMachineID: func(id uint16) bool {
                return g.NodeID() == id
            },
        })
        return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
            is, ok := m.(IDSetter)
            if !ok {
                return nil, fmt.Errorf("unexpected mutation %T", m)
            }
            id, err := sf.NextID()
            if err != nil {
                return nil, err
            }
            is.SetID(fmt.Sprintf("%d", id))
            return next.Mutate(ctx, m)
        })
    }
}
