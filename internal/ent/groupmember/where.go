// Code generated by ent, DO NOT EDIT.

package groupmember

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// MemberID applies equality check predicate on the "member_id" field. It's identical to MemberIDEQ.
func MemberID(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemberID), v))
	})
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// Permission applies equality check predicate on the "permission" field. It's identical to PermissionEQ.
func Permission(v uint8) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermission), v))
	})
}

// Sn applies equality check predicate on the "sn" field. It's identical to SnEQ.
func Sn(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSn), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// MemberIDEQ applies the EQ predicate on the "member_id" field.
func MemberIDEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemberID), v))
	})
}

// MemberIDNEQ applies the NEQ predicate on the "member_id" field.
func MemberIDNEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMemberID), v))
	})
}

// MemberIDIn applies the In predicate on the "member_id" field.
func MemberIDIn(vs ...string) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMemberID), v...))
	})
}

// MemberIDNotIn applies the NotIn predicate on the "member_id" field.
func MemberIDNotIn(vs ...string) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMemberID), v...))
	})
}

// MemberIDGT applies the GT predicate on the "member_id" field.
func MemberIDGT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMemberID), v))
	})
}

// MemberIDGTE applies the GTE predicate on the "member_id" field.
func MemberIDGTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMemberID), v))
	})
}

// MemberIDLT applies the LT predicate on the "member_id" field.
func MemberIDLT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMemberID), v))
	})
}

// MemberIDLTE applies the LTE predicate on the "member_id" field.
func MemberIDLTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMemberID), v))
	})
}

// MemberIDContains applies the Contains predicate on the "member_id" field.
func MemberIDContains(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMemberID), v))
	})
}

// MemberIDHasPrefix applies the HasPrefix predicate on the "member_id" field.
func MemberIDHasPrefix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMemberID), v))
	})
}

// MemberIDHasSuffix applies the HasSuffix predicate on the "member_id" field.
func MemberIDHasSuffix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMemberID), v))
	})
}

// MemberIDEqualFold applies the EqualFold predicate on the "member_id" field.
func MemberIDEqualFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMemberID), v))
	})
}

// MemberIDContainsFold applies the ContainsFold predicate on the "member_id" field.
func MemberIDContainsFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMemberID), v))
	})
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroupID), v))
	})
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...string) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGroupID), v...))
	})
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...string) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGroupID), v...))
	})
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGroupID), v))
	})
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGroupID), v))
	})
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGroupID), v))
	})
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGroupID), v))
	})
}

// GroupIDContains applies the Contains predicate on the "group_id" field.
func GroupIDContains(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGroupID), v))
	})
}

// GroupIDHasPrefix applies the HasPrefix predicate on the "group_id" field.
func GroupIDHasPrefix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGroupID), v))
	})
}

// GroupIDHasSuffix applies the HasSuffix predicate on the "group_id" field.
func GroupIDHasSuffix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGroupID), v))
	})
}

// GroupIDEqualFold applies the EqualFold predicate on the "group_id" field.
func GroupIDEqualFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGroupID), v))
	})
}

// GroupIDContainsFold applies the ContainsFold predicate on the "group_id" field.
func GroupIDContainsFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGroupID), v))
	})
}

// PermissionEQ applies the EQ predicate on the "permission" field.
func PermissionEQ(v uint8) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermission), v))
	})
}

// PermissionNEQ applies the NEQ predicate on the "permission" field.
func PermissionNEQ(v uint8) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPermission), v))
	})
}

// PermissionIn applies the In predicate on the "permission" field.
func PermissionIn(vs ...uint8) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPermission), v...))
	})
}

// PermissionNotIn applies the NotIn predicate on the "permission" field.
func PermissionNotIn(vs ...uint8) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPermission), v...))
	})
}

// PermissionGT applies the GT predicate on the "permission" field.
func PermissionGT(v uint8) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPermission), v))
	})
}

// PermissionGTE applies the GTE predicate on the "permission" field.
func PermissionGTE(v uint8) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPermission), v))
	})
}

// PermissionLT applies the LT predicate on the "permission" field.
func PermissionLT(v uint8) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPermission), v))
	})
}

// PermissionLTE applies the LTE predicate on the "permission" field.
func PermissionLTE(v uint8) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPermission), v))
	})
}

// SnEQ applies the EQ predicate on the "sn" field.
func SnEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSn), v))
	})
}

// SnNEQ applies the NEQ predicate on the "sn" field.
func SnNEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSn), v))
	})
}

// SnIn applies the In predicate on the "sn" field.
func SnIn(vs ...string) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSn), v...))
	})
}

// SnNotIn applies the NotIn predicate on the "sn" field.
func SnNotIn(vs ...string) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSn), v...))
	})
}

// SnGT applies the GT predicate on the "sn" field.
func SnGT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSn), v))
	})
}

// SnGTE applies the GTE predicate on the "sn" field.
func SnGTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSn), v))
	})
}

// SnLT applies the LT predicate on the "sn" field.
func SnLT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSn), v))
	})
}

// SnLTE applies the LTE predicate on the "sn" field.
func SnLTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSn), v))
	})
}

// SnContains applies the Contains predicate on the "sn" field.
func SnContains(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSn), v))
	})
}

// SnHasPrefix applies the HasPrefix predicate on the "sn" field.
func SnHasPrefix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSn), v))
	})
}

// SnHasSuffix applies the HasSuffix predicate on the "sn" field.
func SnHasSuffix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSn), v))
	})
}

// SnEqualFold applies the EqualFold predicate on the "sn" field.
func SnEqualFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSn), v))
	})
}

// SnContainsFold applies the ContainsFold predicate on the "sn" field.
func SnContainsFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSn), v))
	})
}

// HasMember applies the HasEdge predicate on the "member" edge.
func HasMember() predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MemberTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberWith applies the HasEdge predicate on the "member" edge with a given conditions (other predicates).
func HasMemberWith(preds ...predicate.Member) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MemberInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
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
func Not(p predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		p(s.Not())
	})
}