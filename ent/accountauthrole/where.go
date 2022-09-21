// Code generated by ent, DO NOT EDIT.

package accountauthrole

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// ID filters vertices based on their ID field.
func ID(id pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// AuthRoleID applies equality check predicate on the "auth_role_id" field. It's identical to AuthRoleIDEQ.
func AuthRoleID(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthRoleID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AccountAuthRole {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AccountAuthRole {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AccountAuthRole {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AccountAuthRole {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.AccountAuthRole {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.AccountAuthRole {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountID), v))
	})
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...pulid.PULID) predicate.AccountAuthRole {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAccountID), v...))
	})
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...pulid.PULID) predicate.AccountAuthRole {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAccountID), v...))
	})
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountID), v))
	})
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountID), v))
	})
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountID), v))
	})
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountID), v))
	})
}

// AccountIDContains applies the Contains predicate on the "account_id" field.
func AccountIDContains(v pulid.PULID) predicate.AccountAuthRole {
	vc := string(v)
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccountID), vc))
	})
}

// AccountIDHasPrefix applies the HasPrefix predicate on the "account_id" field.
func AccountIDHasPrefix(v pulid.PULID) predicate.AccountAuthRole {
	vc := string(v)
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccountID), vc))
	})
}

// AccountIDHasSuffix applies the HasSuffix predicate on the "account_id" field.
func AccountIDHasSuffix(v pulid.PULID) predicate.AccountAuthRole {
	vc := string(v)
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccountID), vc))
	})
}

// AccountIDEqualFold applies the EqualFold predicate on the "account_id" field.
func AccountIDEqualFold(v pulid.PULID) predicate.AccountAuthRole {
	vc := string(v)
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccountID), vc))
	})
}

// AccountIDContainsFold applies the ContainsFold predicate on the "account_id" field.
func AccountIDContainsFold(v pulid.PULID) predicate.AccountAuthRole {
	vc := string(v)
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccountID), vc))
	})
}

// AuthRoleIDEQ applies the EQ predicate on the "auth_role_id" field.
func AuthRoleIDEQ(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthRoleID), v))
	})
}

// AuthRoleIDNEQ applies the NEQ predicate on the "auth_role_id" field.
func AuthRoleIDNEQ(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthRoleID), v))
	})
}

// AuthRoleIDIn applies the In predicate on the "auth_role_id" field.
func AuthRoleIDIn(vs ...pulid.PULID) predicate.AccountAuthRole {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAuthRoleID), v...))
	})
}

// AuthRoleIDNotIn applies the NotIn predicate on the "auth_role_id" field.
func AuthRoleIDNotIn(vs ...pulid.PULID) predicate.AccountAuthRole {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAuthRoleID), v...))
	})
}

// AuthRoleIDGT applies the GT predicate on the "auth_role_id" field.
func AuthRoleIDGT(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthRoleID), v))
	})
}

// AuthRoleIDGTE applies the GTE predicate on the "auth_role_id" field.
func AuthRoleIDGTE(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthRoleID), v))
	})
}

// AuthRoleIDLT applies the LT predicate on the "auth_role_id" field.
func AuthRoleIDLT(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthRoleID), v))
	})
}

// AuthRoleIDLTE applies the LTE predicate on the "auth_role_id" field.
func AuthRoleIDLTE(v pulid.PULID) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthRoleID), v))
	})
}

// AuthRoleIDContains applies the Contains predicate on the "auth_role_id" field.
func AuthRoleIDContains(v pulid.PULID) predicate.AccountAuthRole {
	vc := string(v)
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthRoleID), vc))
	})
}

// AuthRoleIDHasPrefix applies the HasPrefix predicate on the "auth_role_id" field.
func AuthRoleIDHasPrefix(v pulid.PULID) predicate.AccountAuthRole {
	vc := string(v)
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthRoleID), vc))
	})
}

// AuthRoleIDHasSuffix applies the HasSuffix predicate on the "auth_role_id" field.
func AuthRoleIDHasSuffix(v pulid.PULID) predicate.AccountAuthRole {
	vc := string(v)
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthRoleID), vc))
	})
}

// AuthRoleIDEqualFold applies the EqualFold predicate on the "auth_role_id" field.
func AuthRoleIDEqualFold(v pulid.PULID) predicate.AccountAuthRole {
	vc := string(v)
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthRoleID), vc))
	})
}

// AuthRoleIDContainsFold applies the ContainsFold predicate on the "auth_role_id" field.
func AuthRoleIDContainsFold(v pulid.PULID) predicate.AccountAuthRole {
	vc := string(v)
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthRoleID), vc))
	})
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAuthRole applies the HasEdge predicate on the "auth_role" edge.
func HasAuthRole() predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthRoleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AuthRoleTable, AuthRoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthRoleWith applies the HasEdge predicate on the "auth_role" edge with a given conditions (other predicates).
func HasAuthRoleWith(preds ...predicate.AuthRole) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthRoleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AuthRoleTable, AuthRoleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccountAuthRole) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccountAuthRole) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
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
func Not(p predicate.AccountAuthRole) predicate.AccountAuthRole {
	return predicate.AccountAuthRole(func(s *sql.Selector) {
		p(s.Not())
	})
}