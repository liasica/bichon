// Code generated by ent, DO NOT EDIT.

package member

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// Intro applies equality check predicate on the "intro" field. It's identical to IntroEQ.
func Intro(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntro), v))
	})
}

// PublicKey applies equality check predicate on the "public_key" field. It's identical to PublicKeyEQ.
func PublicKey(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublicKey), v))
	})
}

// Nonce applies equality check predicate on the "nonce" field. It's identical to NonceEQ.
func Nonce(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNonce), v))
	})
}

// ShowNickname applies equality check predicate on the "show_nickname" field. It's identical to ShowNicknameEQ.
func ShowNickname(v bool) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShowNickname), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNickname), v))
	})
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNickname), v...))
	})
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNickname), v...))
	})
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNickname), v))
	})
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNickname), v))
	})
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNickname), v))
	})
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNickname), v))
	})
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNickname), v))
	})
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNickname), v))
	})
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNickname), v))
	})
}

// NicknameIsNil applies the IsNil predicate on the "nickname" field.
func NicknameIsNil() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNickname)))
	})
}

// NicknameNotNil applies the NotNil predicate on the "nickname" field.
func NicknameNotNil() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNickname)))
	})
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNickname), v))
	})
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNickname), v))
	})
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvatar), v))
	})
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAvatar), v...))
	})
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAvatar), v...))
	})
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvatar), v))
	})
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvatar), v))
	})
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvatar), v))
	})
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvatar), v))
	})
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAvatar), v))
	})
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAvatar), v))
	})
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAvatar), v))
	})
}

// AvatarIsNil applies the IsNil predicate on the "avatar" field.
func AvatarIsNil() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAvatar)))
	})
}

// AvatarNotNil applies the NotNil predicate on the "avatar" field.
func AvatarNotNil() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAvatar)))
	})
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAvatar), v))
	})
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAvatar), v))
	})
}

// IntroEQ applies the EQ predicate on the "intro" field.
func IntroEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntro), v))
	})
}

// IntroNEQ applies the NEQ predicate on the "intro" field.
func IntroNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIntro), v))
	})
}

// IntroIn applies the In predicate on the "intro" field.
func IntroIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIntro), v...))
	})
}

// IntroNotIn applies the NotIn predicate on the "intro" field.
func IntroNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIntro), v...))
	})
}

// IntroGT applies the GT predicate on the "intro" field.
func IntroGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIntro), v))
	})
}

// IntroGTE applies the GTE predicate on the "intro" field.
func IntroGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIntro), v))
	})
}

// IntroLT applies the LT predicate on the "intro" field.
func IntroLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIntro), v))
	})
}

// IntroLTE applies the LTE predicate on the "intro" field.
func IntroLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIntro), v))
	})
}

// IntroContains applies the Contains predicate on the "intro" field.
func IntroContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIntro), v))
	})
}

// IntroHasPrefix applies the HasPrefix predicate on the "intro" field.
func IntroHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIntro), v))
	})
}

// IntroHasSuffix applies the HasSuffix predicate on the "intro" field.
func IntroHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIntro), v))
	})
}

// IntroIsNil applies the IsNil predicate on the "intro" field.
func IntroIsNil() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIntro)))
	})
}

// IntroNotNil applies the NotNil predicate on the "intro" field.
func IntroNotNil() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIntro)))
	})
}

// IntroEqualFold applies the EqualFold predicate on the "intro" field.
func IntroEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIntro), v))
	})
}

// IntroContainsFold applies the ContainsFold predicate on the "intro" field.
func IntroContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIntro), v))
	})
}

// PublicKeyEQ applies the EQ predicate on the "public_key" field.
func PublicKeyEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublicKey), v))
	})
}

// PublicKeyNEQ applies the NEQ predicate on the "public_key" field.
func PublicKeyNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPublicKey), v))
	})
}

// PublicKeyIn applies the In predicate on the "public_key" field.
func PublicKeyIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPublicKey), v...))
	})
}

// PublicKeyNotIn applies the NotIn predicate on the "public_key" field.
func PublicKeyNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPublicKey), v...))
	})
}

// PublicKeyGT applies the GT predicate on the "public_key" field.
func PublicKeyGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPublicKey), v))
	})
}

// PublicKeyGTE applies the GTE predicate on the "public_key" field.
func PublicKeyGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPublicKey), v))
	})
}

// PublicKeyLT applies the LT predicate on the "public_key" field.
func PublicKeyLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPublicKey), v))
	})
}

// PublicKeyLTE applies the LTE predicate on the "public_key" field.
func PublicKeyLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPublicKey), v))
	})
}

// PublicKeyContains applies the Contains predicate on the "public_key" field.
func PublicKeyContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPublicKey), v))
	})
}

// PublicKeyHasPrefix applies the HasPrefix predicate on the "public_key" field.
func PublicKeyHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPublicKey), v))
	})
}

// PublicKeyHasSuffix applies the HasSuffix predicate on the "public_key" field.
func PublicKeyHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPublicKey), v))
	})
}

// PublicKeyIsNil applies the IsNil predicate on the "public_key" field.
func PublicKeyIsNil() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPublicKey)))
	})
}

// PublicKeyNotNil applies the NotNil predicate on the "public_key" field.
func PublicKeyNotNil() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPublicKey)))
	})
}

// PublicKeyEqualFold applies the EqualFold predicate on the "public_key" field.
func PublicKeyEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPublicKey), v))
	})
}

// PublicKeyContainsFold applies the ContainsFold predicate on the "public_key" field.
func PublicKeyContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPublicKey), v))
	})
}

// NonceEQ applies the EQ predicate on the "nonce" field.
func NonceEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNonce), v))
	})
}

// NonceNEQ applies the NEQ predicate on the "nonce" field.
func NonceNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNonce), v))
	})
}

// NonceIn applies the In predicate on the "nonce" field.
func NonceIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNonce), v...))
	})
}

// NonceNotIn applies the NotIn predicate on the "nonce" field.
func NonceNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNonce), v...))
	})
}

// NonceGT applies the GT predicate on the "nonce" field.
func NonceGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNonce), v))
	})
}

// NonceGTE applies the GTE predicate on the "nonce" field.
func NonceGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNonce), v))
	})
}

// NonceLT applies the LT predicate on the "nonce" field.
func NonceLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNonce), v))
	})
}

// NonceLTE applies the LTE predicate on the "nonce" field.
func NonceLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNonce), v))
	})
}

// NonceContains applies the Contains predicate on the "nonce" field.
func NonceContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNonce), v))
	})
}

// NonceHasPrefix applies the HasPrefix predicate on the "nonce" field.
func NonceHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNonce), v))
	})
}

// NonceHasSuffix applies the HasSuffix predicate on the "nonce" field.
func NonceHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNonce), v))
	})
}

// NonceEqualFold applies the EqualFold predicate on the "nonce" field.
func NonceEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNonce), v))
	})
}

// NonceContainsFold applies the ContainsFold predicate on the "nonce" field.
func NonceContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNonce), v))
	})
}

// ShowNicknameEQ applies the EQ predicate on the "show_nickname" field.
func ShowNicknameEQ(v bool) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShowNickname), v))
	})
}

// ShowNicknameNEQ applies the NEQ predicate on the "show_nickname" field.
func ShowNicknameNEQ(v bool) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShowNickname), v))
	})
}

// HasOwnGroups applies the HasEdge predicate on the "own_groups" edge.
func HasOwnGroups() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnGroupsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OwnGroupsTable, OwnGroupsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnGroupsWith applies the HasEdge predicate on the "own_groups" edge with a given conditions (other predicates).
func HasOwnGroupsWith(preds ...predicate.Group) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnGroupsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OwnGroupsTable, OwnGroupsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMessages applies the HasEdge predicate on the "messages" edge.
func HasMessages() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MessagesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessagesWith applies the HasEdge predicate on the "messages" edge with a given conditions (other predicates).
func HasMessagesWith(preds ...predicate.Message) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
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

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupsTable, GroupsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupsTable, GroupsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroupMembers applies the HasEdge predicate on the "group_members" edge.
func HasGroupMembers() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupMembersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, GroupMembersTable, GroupMembersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupMembersWith applies the HasEdge predicate on the "group_members" edge with a given conditions (other predicates).
func HasGroupMembersWith(preds ...predicate.GroupMember) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
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
func And(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
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
func Not(p predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		p(s.Not())
	})
}
