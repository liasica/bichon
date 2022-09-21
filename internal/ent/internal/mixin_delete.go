package internal

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/mixin"
)

// DeleteMixin Soft delete mixin
type DeleteMixin struct {
    mixin.Schema
}

func (DeleteMixin) Fields() []ent.Field {
    return []ent.Field{
        field.Time("deleted_at").Nillable().Optional(),
    }
}
