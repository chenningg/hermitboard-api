// Code generated by ent, DO NOT EDIT.

package connection

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// ID filters vertices based on their ID field.
func ID(id pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// AccessToken applies equality check predicate on the "access_token" field. It's identical to AccessTokenEQ.
func AccessToken(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// AccessTokenEQ applies the EQ predicate on the "access_token" field.
func AccessTokenEQ(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenNEQ applies the NEQ predicate on the "access_token" field.
func AccessTokenNEQ(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenIn applies the In predicate on the "access_token" field.
func AccessTokenIn(vs ...string) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenNotIn applies the NotIn predicate on the "access_token" field.
func AccessTokenNotIn(vs ...string) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenGT applies the GT predicate on the "access_token" field.
func AccessTokenGT(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenGTE applies the GTE predicate on the "access_token" field.
func AccessTokenGTE(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLT applies the LT predicate on the "access_token" field.
func AccessTokenLT(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLTE applies the LTE predicate on the "access_token" field.
func AccessTokenLTE(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContains applies the Contains predicate on the "access_token" field.
func AccessTokenContains(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasPrefix applies the HasPrefix predicate on the "access_token" field.
func AccessTokenHasPrefix(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasSuffix applies the HasSuffix predicate on the "access_token" field.
func AccessTokenHasSuffix(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenEqualFold applies the EqualFold predicate on the "access_token" field.
func AccessTokenEqualFold(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContainsFold applies the ContainsFold predicate on the "access_token" field.
func AccessTokenContainsFold(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccessToken), v))
	})
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountID), v))
	})
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...pulid.PULID) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAccountID), v...))
	})
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...pulid.PULID) predicate.Connection {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAccountID), v...))
	})
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountID), v))
	})
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountID), v))
	})
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountID), v))
	})
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v pulid.PULID) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountID), v))
	})
}

// AccountIDContains applies the Contains predicate on the "account_id" field.
func AccountIDContains(v pulid.PULID) predicate.Connection {
	vc := string(v)
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccountID), vc))
	})
}

// AccountIDHasPrefix applies the HasPrefix predicate on the "account_id" field.
func AccountIDHasPrefix(v pulid.PULID) predicate.Connection {
	vc := string(v)
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccountID), vc))
	})
}

// AccountIDHasSuffix applies the HasSuffix predicate on the "account_id" field.
func AccountIDHasSuffix(v pulid.PULID) predicate.Connection {
	vc := string(v)
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccountID), vc))
	})
}

// AccountIDEqualFold applies the EqualFold predicate on the "account_id" field.
func AccountIDEqualFold(v pulid.PULID) predicate.Connection {
	vc := string(v)
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccountID), vc))
	})
}

// AccountIDContainsFold applies the ContainsFold predicate on the "account_id" field.
func AccountIDContainsFold(v pulid.PULID) predicate.Connection {
	vc := string(v)
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccountID), vc))
	})
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPortfolios applies the HasEdge predicate on the "portfolios" edge.
func HasPortfolios() predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PortfoliosTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PortfoliosTable, PortfoliosPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPortfoliosWith applies the HasEdge predicate on the "portfolios" edge with a given conditions (other predicates).
func HasPortfoliosWith(preds ...predicate.Portfolio) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PortfoliosInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PortfoliosTable, PortfoliosPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Connection) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Connection) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
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
func Not(p predicate.Connection) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		p(s.Not())
	})
}
