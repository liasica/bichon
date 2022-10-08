package internal

import (
    "context"
    "entgo.io/ent"
    "entgo.io/ent/entc/integration/ent/hook"
    "github.com/chatpuppy/puppychat/app/model"
)

type DistributeMutator interface {
    Parse() *model.SyncData
    Columns() []string
}

func DistributeHook() ent.Hook {
    hk := func(next ent.Mutator) ent.Mutator {
        return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (v ent.Value, err error) {
            v, err = next.Mutate(ctx, mutation)
            if err == nil {
                m, ok := mutation.(DistributeMutator)
                if ok {
                    data := m.Parse()
                    nd, _ := ctx.Value("doNotDistribute").(bool)
                    if data != nil && !nd {
                        model.DistributionBroadcastChan <- data
                    }
                }
            }
            return
        })
    }

    return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne)
}
