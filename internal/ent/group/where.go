// Code generated by ent, DO NOT EDIT.

package group

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategory), v))
	})
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwnerID), v))
	})
}

// MembersMax applies equality check predicate on the "members_max" field. It's identical to MembersMaxEQ.
func MembersMax(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMembersMax), v))
	})
}

// MembersCount applies equality check predicate on the "members_count" field. It's identical to MembersCountEQ.
func MembersCount(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMembersCount), v))
	})
}

// Public applies equality check predicate on the "public" field. It's identical to PublicEQ.
func Public(v bool) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublic), v))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// Intro applies equality check predicate on the "intro" field. It's identical to IntroEQ.
func Intro(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntro), v))
	})
}

// Keys applies equality check predicate on the "keys" field. It's identical to KeysEQ.
func Keys(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKeys), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategory), v))
	})
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategory), v))
	})
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCategory), v...))
	})
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCategory), v...))
	})
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCategory), v))
	})
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCategory), v))
	})
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCategory), v))
	})
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCategory), v))
	})
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCategory), v))
	})
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCategory), v))
	})
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCategory), v))
	})
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCategory), v))
	})
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCategory), v))
	})
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwnerID), v))
	})
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOwnerID), v))
	})
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOwnerID), v...))
	})
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOwnerID), v...))
	})
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOwnerID), v))
	})
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOwnerID), v))
	})
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOwnerID), v))
	})
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOwnerID), v))
	})
}

// OwnerIDContains applies the Contains predicate on the "owner_id" field.
func OwnerIDContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOwnerID), v))
	})
}

// OwnerIDHasPrefix applies the HasPrefix predicate on the "owner_id" field.
func OwnerIDHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOwnerID), v))
	})
}

// OwnerIDHasSuffix applies the HasSuffix predicate on the "owner_id" field.
func OwnerIDHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOwnerID), v))
	})
}

// OwnerIDEqualFold applies the EqualFold predicate on the "owner_id" field.
func OwnerIDEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOwnerID), v))
	})
}

// OwnerIDContainsFold applies the ContainsFold predicate on the "owner_id" field.
func OwnerIDContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOwnerID), v))
	})
}

// MembersMaxEQ applies the EQ predicate on the "members_max" field.
func MembersMaxEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMembersMax), v))
	})
}

// MembersMaxNEQ applies the NEQ predicate on the "members_max" field.
func MembersMaxNEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMembersMax), v))
	})
}

// MembersMaxIn applies the In predicate on the "members_max" field.
func MembersMaxIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMembersMax), v...))
	})
}

// MembersMaxNotIn applies the NotIn predicate on the "members_max" field.
func MembersMaxNotIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMembersMax), v...))
	})
}

// MembersMaxGT applies the GT predicate on the "members_max" field.
func MembersMaxGT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMembersMax), v))
	})
}

// MembersMaxGTE applies the GTE predicate on the "members_max" field.
func MembersMaxGTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMembersMax), v))
	})
}

// MembersMaxLT applies the LT predicate on the "members_max" field.
func MembersMaxLT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMembersMax), v))
	})
}

// MembersMaxLTE applies the LTE predicate on the "members_max" field.
func MembersMaxLTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMembersMax), v))
	})
}

// MembersCountEQ applies the EQ predicate on the "members_count" field.
func MembersCountEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMembersCount), v))
	})
}

// MembersCountNEQ applies the NEQ predicate on the "members_count" field.
func MembersCountNEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMembersCount), v))
	})
}

// MembersCountIn applies the In predicate on the "members_count" field.
func MembersCountIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMembersCount), v...))
	})
}

// MembersCountNotIn applies the NotIn predicate on the "members_count" field.
func MembersCountNotIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMembersCount), v...))
	})
}

// MembersCountGT applies the GT predicate on the "members_count" field.
func MembersCountGT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMembersCount), v))
	})
}

// MembersCountGTE applies the GTE predicate on the "members_count" field.
func MembersCountGTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMembersCount), v))
	})
}

// MembersCountLT applies the LT predicate on the "members_count" field.
func MembersCountLT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMembersCount), v))
	})
}

// MembersCountLTE applies the LTE predicate on the "members_count" field.
func MembersCountLTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMembersCount), v))
	})
}

// PublicEQ applies the EQ predicate on the "public" field.
func PublicEQ(v bool) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublic), v))
	})
}

// PublicNEQ applies the NEQ predicate on the "public" field.
func PublicNEQ(v bool) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPublic), v))
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// IntroEQ applies the EQ predicate on the "intro" field.
func IntroEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntro), v))
	})
}

// IntroNEQ applies the NEQ predicate on the "intro" field.
func IntroNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIntro), v))
	})
}

// IntroIn applies the In predicate on the "intro" field.
func IntroIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIntro), v...))
	})
}

// IntroNotIn applies the NotIn predicate on the "intro" field.
func IntroNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIntro), v...))
	})
}

// IntroGT applies the GT predicate on the "intro" field.
func IntroGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIntro), v))
	})
}

// IntroGTE applies the GTE predicate on the "intro" field.
func IntroGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIntro), v))
	})
}

// IntroLT applies the LT predicate on the "intro" field.
func IntroLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIntro), v))
	})
}

// IntroLTE applies the LTE predicate on the "intro" field.
func IntroLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIntro), v))
	})
}

// IntroContains applies the Contains predicate on the "intro" field.
func IntroContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIntro), v))
	})
}

// IntroHasPrefix applies the HasPrefix predicate on the "intro" field.
func IntroHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIntro), v))
	})
}

// IntroHasSuffix applies the HasSuffix predicate on the "intro" field.
func IntroHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIntro), v))
	})
}

// IntroIsNil applies the IsNil predicate on the "intro" field.
func IntroIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIntro)))
	})
}

// IntroNotNil applies the NotNil predicate on the "intro" field.
func IntroNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIntro)))
	})
}

// IntroEqualFold applies the EqualFold predicate on the "intro" field.
func IntroEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIntro), v))
	})
}

// IntroContainsFold applies the ContainsFold predicate on the "intro" field.
func IntroContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIntro), v))
	})
}

// KeysEQ applies the EQ predicate on the "keys" field.
func KeysEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKeys), v))
	})
}

// KeysNEQ applies the NEQ predicate on the "keys" field.
func KeysNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKeys), v))
	})
}

// KeysIn applies the In predicate on the "keys" field.
func KeysIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldKeys), v...))
	})
}

// KeysNotIn applies the NotIn predicate on the "keys" field.
func KeysNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldKeys), v...))
	})
}

// KeysGT applies the GT predicate on the "keys" field.
func KeysGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKeys), v))
	})
}

// KeysGTE applies the GTE predicate on the "keys" field.
func KeysGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKeys), v))
	})
}

// KeysLT applies the LT predicate on the "keys" field.
func KeysLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKeys), v))
	})
}

// KeysLTE applies the LTE predicate on the "keys" field.
func KeysLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKeys), v))
	})
}

// KeysContains applies the Contains predicate on the "keys" field.
func KeysContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKeys), v))
	})
}

// KeysHasPrefix applies the HasPrefix predicate on the "keys" field.
func KeysHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKeys), v))
	})
}

// KeysHasSuffix applies the HasSuffix predicate on the "keys" field.
func KeysHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKeys), v))
	})
}

// KeysEqualFold applies the EqualFold predicate on the "keys" field.
func KeysEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKeys), v))
	})
}

// KeysContainsFold applies the ContainsFold predicate on the "keys" field.
func KeysContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKeys), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Member) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMessages applies the HasEdge predicate on the "messages" edge.
func HasMessages() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MessagesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessagesWith applies the HasEdge predicate on the "messages" edge with a given conditions (other predicates).
func HasMessagesWith(preds ...predicate.Message) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MessagesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMembers applies the HasEdge predicate on the "members" edge.
func HasMembers() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MembersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MembersTable, MembersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembersWith applies the HasEdge predicate on the "members" edge with a given conditions (other predicates).
func HasMembersWith(preds ...predicate.Member) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MembersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MembersTable, MembersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroupMembers applies the HasEdge predicate on the "group_members" edge.
func HasGroupMembers() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupMembersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, GroupMembersTable, GroupMembersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupMembersWith applies the HasEdge predicate on the "group_members" edge with a given conditions (other predicates).
func HasGroupMembersWith(preds ...predicate.GroupMember) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupMembersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, GroupMembersTable, GroupMembersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		p(s.Not())
	})
}
