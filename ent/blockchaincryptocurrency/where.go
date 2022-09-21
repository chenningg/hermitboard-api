// Code generated by ent, DO NOT EDIT.

package blockchaincryptocurrency

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// ID filters vertices based on their ID field.
func ID(id pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// BlockchainID applies equality check predicate on the "blockchain_id" field. It's identical to BlockchainIDEQ.
func BlockchainID(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockchainID), v))
	})
}

// CryptocurrencyID applies equality check predicate on the "cryptocurrency_id" field. It's identical to CryptocurrencyIDEQ.
func CryptocurrencyID(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCryptocurrencyID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BlockchainCryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BlockchainCryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.BlockchainCryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.BlockchainCryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.BlockchainCryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.BlockchainCryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// BlockchainIDEQ applies the EQ predicate on the "blockchain_id" field.
func BlockchainIDEQ(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockchainID), v))
	})
}

// BlockchainIDNEQ applies the NEQ predicate on the "blockchain_id" field.
func BlockchainIDNEQ(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBlockchainID), v))
	})
}

// BlockchainIDIn applies the In predicate on the "blockchain_id" field.
func BlockchainIDIn(vs ...pulid.PULID) predicate.BlockchainCryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBlockchainID), v...))
	})
}

// BlockchainIDNotIn applies the NotIn predicate on the "blockchain_id" field.
func BlockchainIDNotIn(vs ...pulid.PULID) predicate.BlockchainCryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBlockchainID), v...))
	})
}

// BlockchainIDGT applies the GT predicate on the "blockchain_id" field.
func BlockchainIDGT(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBlockchainID), v))
	})
}

// BlockchainIDGTE applies the GTE predicate on the "blockchain_id" field.
func BlockchainIDGTE(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBlockchainID), v))
	})
}

// BlockchainIDLT applies the LT predicate on the "blockchain_id" field.
func BlockchainIDLT(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBlockchainID), v))
	})
}

// BlockchainIDLTE applies the LTE predicate on the "blockchain_id" field.
func BlockchainIDLTE(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBlockchainID), v))
	})
}

// BlockchainIDContains applies the Contains predicate on the "blockchain_id" field.
func BlockchainIDContains(v pulid.PULID) predicate.BlockchainCryptocurrency {
	vc := string(v)
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBlockchainID), vc))
	})
}

// BlockchainIDHasPrefix applies the HasPrefix predicate on the "blockchain_id" field.
func BlockchainIDHasPrefix(v pulid.PULID) predicate.BlockchainCryptocurrency {
	vc := string(v)
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBlockchainID), vc))
	})
}

// BlockchainIDHasSuffix applies the HasSuffix predicate on the "blockchain_id" field.
func BlockchainIDHasSuffix(v pulid.PULID) predicate.BlockchainCryptocurrency {
	vc := string(v)
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBlockchainID), vc))
	})
}

// BlockchainIDEqualFold applies the EqualFold predicate on the "blockchain_id" field.
func BlockchainIDEqualFold(v pulid.PULID) predicate.BlockchainCryptocurrency {
	vc := string(v)
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBlockchainID), vc))
	})
}

// BlockchainIDContainsFold applies the ContainsFold predicate on the "blockchain_id" field.
func BlockchainIDContainsFold(v pulid.PULID) predicate.BlockchainCryptocurrency {
	vc := string(v)
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBlockchainID), vc))
	})
}

// CryptocurrencyIDEQ applies the EQ predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDEQ(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCryptocurrencyID), v))
	})
}

// CryptocurrencyIDNEQ applies the NEQ predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDNEQ(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCryptocurrencyID), v))
	})
}

// CryptocurrencyIDIn applies the In predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDIn(vs ...pulid.PULID) predicate.BlockchainCryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCryptocurrencyID), v...))
	})
}

// CryptocurrencyIDNotIn applies the NotIn predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDNotIn(vs ...pulid.PULID) predicate.BlockchainCryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCryptocurrencyID), v...))
	})
}

// CryptocurrencyIDGT applies the GT predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDGT(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCryptocurrencyID), v))
	})
}

// CryptocurrencyIDGTE applies the GTE predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDGTE(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCryptocurrencyID), v))
	})
}

// CryptocurrencyIDLT applies the LT predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDLT(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCryptocurrencyID), v))
	})
}

// CryptocurrencyIDLTE applies the LTE predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDLTE(v pulid.PULID) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCryptocurrencyID), v))
	})
}

// CryptocurrencyIDContains applies the Contains predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDContains(v pulid.PULID) predicate.BlockchainCryptocurrency {
	vc := string(v)
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCryptocurrencyID), vc))
	})
}

// CryptocurrencyIDHasPrefix applies the HasPrefix predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDHasPrefix(v pulid.PULID) predicate.BlockchainCryptocurrency {
	vc := string(v)
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCryptocurrencyID), vc))
	})
}

// CryptocurrencyIDHasSuffix applies the HasSuffix predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDHasSuffix(v pulid.PULID) predicate.BlockchainCryptocurrency {
	vc := string(v)
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCryptocurrencyID), vc))
	})
}

// CryptocurrencyIDEqualFold applies the EqualFold predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDEqualFold(v pulid.PULID) predicate.BlockchainCryptocurrency {
	vc := string(v)
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCryptocurrencyID), vc))
	})
}

// CryptocurrencyIDContainsFold applies the ContainsFold predicate on the "cryptocurrency_id" field.
func CryptocurrencyIDContainsFold(v pulid.PULID) predicate.BlockchainCryptocurrency {
	vc := string(v)
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCryptocurrencyID), vc))
	})
}

// HasBlockchain applies the HasEdge predicate on the "blockchain" edge.
func HasBlockchain() predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlockchainTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BlockchainTable, BlockchainColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlockchainWith applies the HasEdge predicate on the "blockchain" edge with a given conditions (other predicates).
func HasBlockchainWith(preds ...predicate.Blockchain) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlockchainInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BlockchainTable, BlockchainColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCryptocurrency applies the HasEdge predicate on the "cryptocurrency" edge.
func HasCryptocurrency() predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CryptocurrencyTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CryptocurrencyTable, CryptocurrencyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCryptocurrencyWith applies the HasEdge predicate on the "cryptocurrency" edge with a given conditions (other predicates).
func HasCryptocurrencyWith(preds ...predicate.Cryptocurrency) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CryptocurrencyInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CryptocurrencyTable, CryptocurrencyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BlockchainCryptocurrency) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BlockchainCryptocurrency) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
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
func Not(p predicate.BlockchainCryptocurrency) predicate.BlockchainCryptocurrency {
	return predicate.BlockchainCryptocurrency(func(s *sql.Selector) {
		p(s.Not())
	})
}
