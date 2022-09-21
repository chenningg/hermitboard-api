// Code generated by ent, DO NOT EDIT.

package staffaccount

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// ID filters vertices based on their ID field.
func ID(id pulid.PULID) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.PULID) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.PULID) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.PULID) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.PULID) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.PULID) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.PULID) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.PULID) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.PULID) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordUpdatedAt applies equality check predicate on the "password_updated_at" field. It's identical to PasswordUpdatedAtEQ.
func PasswordUpdatedAt(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPasswordUpdatedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// AuthTypeEQ applies the EQ predicate on the "auth_type" field.
func AuthTypeEQ(v AuthType) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthType), v))
	})
}

// AuthTypeNEQ applies the NEQ predicate on the "auth_type" field.
func AuthTypeNEQ(v AuthType) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthType), v))
	})
}

// AuthTypeIn applies the In predicate on the "auth_type" field.
func AuthTypeIn(vs ...AuthType) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAuthType), v...))
	})
}

// AuthTypeNotIn applies the NotIn predicate on the "auth_type" field.
func AuthTypeNotIn(vs ...AuthType) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAuthType), v...))
	})
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNickname), v))
	})
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNickname), v...))
	})
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNickname), v...))
	})
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNickname), v))
	})
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNickname), v))
	})
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNickname), v))
	})
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNickname), v))
	})
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNickname), v))
	})
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNickname), v))
	})
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNickname), v))
	})
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNickname), v))
	})
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNickname), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPassword)))
	})
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPassword)))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// PasswordUpdatedAtEQ applies the EQ predicate on the "password_updated_at" field.
func PasswordUpdatedAtEQ(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPasswordUpdatedAt), v))
	})
}

// PasswordUpdatedAtNEQ applies the NEQ predicate on the "password_updated_at" field.
func PasswordUpdatedAtNEQ(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPasswordUpdatedAt), v))
	})
}

// PasswordUpdatedAtIn applies the In predicate on the "password_updated_at" field.
func PasswordUpdatedAtIn(vs ...time.Time) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPasswordUpdatedAt), v...))
	})
}

// PasswordUpdatedAtNotIn applies the NotIn predicate on the "password_updated_at" field.
func PasswordUpdatedAtNotIn(vs ...time.Time) predicate.StaffAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPasswordUpdatedAt), v...))
	})
}

// PasswordUpdatedAtGT applies the GT predicate on the "password_updated_at" field.
func PasswordUpdatedAtGT(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPasswordUpdatedAt), v))
	})
}

// PasswordUpdatedAtGTE applies the GTE predicate on the "password_updated_at" field.
func PasswordUpdatedAtGTE(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPasswordUpdatedAt), v))
	})
}

// PasswordUpdatedAtLT applies the LT predicate on the "password_updated_at" field.
func PasswordUpdatedAtLT(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPasswordUpdatedAt), v))
	})
}

// PasswordUpdatedAtLTE applies the LTE predicate on the "password_updated_at" field.
func PasswordUpdatedAtLTE(v time.Time) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPasswordUpdatedAt), v))
	})
}

// HasAuthRoles applies the HasEdge predicate on the "auth_roles" edge.
func HasAuthRoles() predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthRolesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AuthRolesTable, AuthRolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthRolesWith applies the HasEdge predicate on the "auth_roles" edge with a given conditions (other predicates).
func HasAuthRolesWith(preds ...predicate.AuthRole) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthRolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AuthRolesTable, AuthRolesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStaffAccountAuthRoles applies the HasEdge predicate on the "staff_account_auth_roles" edge.
func HasStaffAccountAuthRoles() predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StaffAccountAuthRolesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, StaffAccountAuthRolesTable, StaffAccountAuthRolesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStaffAccountAuthRolesWith applies the HasEdge predicate on the "staff_account_auth_roles" edge with a given conditions (other predicates).
func HasStaffAccountAuthRolesWith(preds ...predicate.StaffAccountAuthRole) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StaffAccountAuthRolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, StaffAccountAuthRolesTable, StaffAccountAuthRolesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StaffAccount) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StaffAccount) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
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
func Not(p predicate.StaffAccount) predicate.StaffAccount {
	return predicate.StaffAccount(func(s *sql.Selector) {
		p(s.Not())
	})
}
