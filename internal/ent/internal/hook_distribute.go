package internal

import (
    "context"
    "entgo.io/ent"
    "entgo.io/ent/entc/integration/ent/hook"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/g"
)

type DistributeMutator interface {
    Parse(op ent.Op) *model.SyncData
    Columns() []string
    LastNode() (int64, bool)
}

func DistributeHook() ent.Hook {
    hk := func(next ent.Mutator) ent.Mutator {
        return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (v ent.Value, err error) {
            v, err = next.Mutate(ctx, mutation)
            if err == nil {
                m, ok := mutation.(DistributeMutator)
                if ok {
                    data := m.Parse(mutation.Op())
                    lastNode, _ := m.LastNode()
                    if data != nil && (g.IsDistributionNode() || !g.IsDistributionNodeID(lastNode)) {
                        data.NodeID, _ = m.LastNode()
                        model.DistributionBroadcastChan <- data
                    }
                }
            }
            return
        })
    }

    return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne)
}
