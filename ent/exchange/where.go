// Code generated by ent, DO NOT EDIT.

package exchange

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/oklog/ulid/v2"
	ulid "github.com/oklog/ulid/v2"
)

// ID filters vertices based on their ID field.
func ID(id ulid.ULID) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ulid.ULID) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ulid.ULID) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ulid.ULID) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...ulid.ULID) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id ulid.ULID) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ulid.ULID) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ulid.ULID) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ulid.ULID) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Exchange {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Exchange {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Exchange {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Exchange {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Exchange {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Exchange {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIcon), v))
	})
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.Exchange {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIcon), v...))
	})
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.Exchange {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIcon), v...))
	})
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIcon), v))
	})
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIcon), v))
	})
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIcon), v))
	})
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIcon), v))
	})
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIcon), v))
	})
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIcon), v))
	})
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIcon), v))
	})
}

// IconIsNil applies the IsNil predicate on the "icon" field.
func IconIsNil() predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIcon)))
	})
}

// IconNotNil applies the NotNil predicate on the "icon" field.
func IconNotNil() predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIcon)))
	})
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIcon), v))
	})
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIcon), v))
	})
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Exchange {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Exchange {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLIsNil applies the IsNil predicate on the "url" field.
func URLIsNil() predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldURL)))
	})
}

// URLNotNil applies the NotNil predicate on the "url" field.
func URLNotNil() predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldURL)))
	})
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// HasTransactions applies the HasEdge predicate on the "transactions" edge.
func HasTransactions() predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, TransactionsTable, TransactionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionsWith applies the HasEdge predicate on the "transactions" edge with a given conditions (other predicates).
func HasTransactionsWith(preds ...predicate.Transaction) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, TransactionsTable, TransactionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Exchange) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Exchange) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
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
func Not(p predicate.Exchange) predicate.Exchange {
	return predicate.Exchange(func(s *sql.Selector) {
		p(s.Not())
	})
}
