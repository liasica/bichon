package internal

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "entgo.io/ent/schema/mixin"
    "time"
)

type TimeMixin struct {
    mixin.Schema
    DoNotIndexCreatedAt bool
}

func (TimeMixin) Fields() []ent.Field {
    return []ent.Field{
        field.Time("created_at").Immutable().Default(time.Now()),
    }
}

func (t TimeMixin) Indexes() []ent.Index {
    var list []ent.Index
    if !t.DoNotIndexCreatedAt {
        list = append(list, index.Fields("created_at"))
    }
    return list
}
