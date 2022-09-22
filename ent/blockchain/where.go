// Code generated by ent, DO NOT EDIT.

package blockchain

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// ID filters vertices based on their ID field.
func ID(id pulid.PULID) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.PULID) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.PULID) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.PULID) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.PULID) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.PULID) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.PULID) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.PULID) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.PULID) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Symbol applies equality check predicate on the "symbol" field. It's identical to SymbolEQ.
func Symbol(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// ChainID applies equality check predicate on the "chain_id" field. It's identical to ChainIDEQ.
func ChainID(v int64) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SymbolEQ applies the EQ predicate on the "symbol" field.
func SymbolEQ(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// SymbolNEQ applies the NEQ predicate on the "symbol" field.
func SymbolNEQ(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSymbol), v))
	})
}

// SymbolIn applies the In predicate on the "symbol" field.
func SymbolIn(vs ...string) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSymbol), v...))
	})
}

// SymbolNotIn applies the NotIn predicate on the "symbol" field.
func SymbolNotIn(vs ...string) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSymbol), v...))
	})
}

// SymbolGT applies the GT predicate on the "symbol" field.
func SymbolGT(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSymbol), v))
	})
}

// SymbolGTE applies the GTE predicate on the "symbol" field.
func SymbolGTE(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSymbol), v))
	})
}

// SymbolLT applies the LT predicate on the "symbol" field.
func SymbolLT(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSymbol), v))
	})
}

// SymbolLTE applies the LTE predicate on the "symbol" field.
func SymbolLTE(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSymbol), v))
	})
}

// SymbolContains applies the Contains predicate on the "symbol" field.
func SymbolContains(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSymbol), v))
	})
}

// SymbolHasPrefix applies the HasPrefix predicate on the "symbol" field.
func SymbolHasPrefix(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSymbol), v))
	})
}

// SymbolHasSuffix applies the HasSuffix predicate on the "symbol" field.
func SymbolHasSuffix(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSymbol), v))
	})
}

// SymbolEqualFold applies the EqualFold predicate on the "symbol" field.
func SymbolEqualFold(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSymbol), v))
	})
}

// SymbolContainsFold applies the ContainsFold predicate on the "symbol" field.
func SymbolContainsFold(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSymbol), v))
	})
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIcon), v))
	})
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIcon), v...))
	})
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIcon), v...))
	})
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIcon), v))
	})
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIcon), v))
	})
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIcon), v))
	})
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIcon), v))
	})
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIcon), v))
	})
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIcon), v))
	})
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIcon), v))
	})
}

// IconIsNil applies the IsNil predicate on the "icon" field.
func IconIsNil() predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIcon)))
	})
}

// IconNotNil applies the NotNil predicate on the "icon" field.
func IconNotNil() predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIcon)))
	})
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIcon), v))
	})
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIcon), v))
	})
}

// ChainIDEQ applies the EQ predicate on the "chain_id" field.
func ChainIDEQ(v int64) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainID), v))
	})
}

// ChainIDNEQ applies the NEQ predicate on the "chain_id" field.
func ChainIDNEQ(v int64) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChainID), v))
	})
}

// ChainIDIn applies the In predicate on the "chain_id" field.
func ChainIDIn(vs ...int64) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldChainID), v...))
	})
}

// ChainIDNotIn applies the NotIn predicate on the "chain_id" field.
func ChainIDNotIn(vs ...int64) predicate.Blockchain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldChainID), v...))
	})
}

// ChainIDGT applies the GT predicate on the "chain_id" field.
func ChainIDGT(v int64) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChainID), v))
	})
}

// ChainIDGTE applies the GTE predicate on the "chain_id" field.
func ChainIDGTE(v int64) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChainID), v))
	})
}

// ChainIDLT applies the LT predicate on the "chain_id" field.
func ChainIDLT(v int64) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChainID), v))
	})
}

// ChainIDLTE applies the LTE predicate on the "chain_id" field.
func ChainIDLTE(v int64) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChainID), v))
	})
}

// ChainIDIsNil applies the IsNil predicate on the "chain_id" field.
func ChainIDIsNil() predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldChainID)))
	})
}

// ChainIDNotNil applies the NotNil predicate on the "chain_id" field.
func ChainIDNotNil() predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldChainID)))
	})
}

// HasCryptocurrencies applies the HasEdge predicate on the "cryptocurrencies" edge.
func HasCryptocurrencies() predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CryptocurrenciesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CryptocurrenciesTable, CryptocurrenciesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCryptocurrenciesWith applies the HasEdge predicate on the "cryptocurrencies" edge with a given conditions (other predicates).
func HasCryptocurrenciesWith(preds ...predicate.Cryptocurrency) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CryptocurrenciesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CryptocurrenciesTable, CryptocurrenciesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Blockchain) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Blockchain) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
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
func Not(p predicate.Blockchain) predicate.Blockchain {
	return predicate.Blockchain(func(s *sql.Selector) {
		p(s.Not())
	})
}
