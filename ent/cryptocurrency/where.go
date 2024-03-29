// Code generated by ent, DO NOT EDIT.

package cryptocurrency

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// ID filters vertices based on their ID field.
func ID(id pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Symbol applies equality check predicate on the "symbol" field. It's identical to SymbolEQ.
func Symbol(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// AssetID applies equality check predicate on the "asset_id" field. It's identical to AssetIDEQ.
func AssetID(v pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssetID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// SymbolEQ applies the EQ predicate on the "symbol" field.
func SymbolEQ(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// SymbolNEQ applies the NEQ predicate on the "symbol" field.
func SymbolNEQ(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSymbol), v))
	})
}

// SymbolIn applies the In predicate on the "symbol" field.
func SymbolIn(vs ...string) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSymbol), v...))
	})
}

// SymbolNotIn applies the NotIn predicate on the "symbol" field.
func SymbolNotIn(vs ...string) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSymbol), v...))
	})
}

// SymbolGT applies the GT predicate on the "symbol" field.
func SymbolGT(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSymbol), v))
	})
}

// SymbolGTE applies the GTE predicate on the "symbol" field.
func SymbolGTE(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSymbol), v))
	})
}

// SymbolLT applies the LT predicate on the "symbol" field.
func SymbolLT(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSymbol), v))
	})
}

// SymbolLTE applies the LTE predicate on the "symbol" field.
func SymbolLTE(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSymbol), v))
	})
}

// SymbolContains applies the Contains predicate on the "symbol" field.
func SymbolContains(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSymbol), v))
	})
}

// SymbolHasPrefix applies the HasPrefix predicate on the "symbol" field.
func SymbolHasPrefix(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSymbol), v))
	})
}

// SymbolHasSuffix applies the HasSuffix predicate on the "symbol" field.
func SymbolHasSuffix(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSymbol), v))
	})
}

// SymbolEqualFold applies the EqualFold predicate on the "symbol" field.
func SymbolEqualFold(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSymbol), v))
	})
}

// SymbolContainsFold applies the ContainsFold predicate on the "symbol" field.
func SymbolContainsFold(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSymbol), v))
	})
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIcon), v))
	})
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIcon), v...))
	})
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIcon), v...))
	})
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIcon), v))
	})
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIcon), v))
	})
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIcon), v))
	})
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIcon), v))
	})
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIcon), v))
	})
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIcon), v))
	})
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIcon), v))
	})
}

// IconIsNil applies the IsNil predicate on the "icon" field.
func IconIsNil() predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIcon)))
	})
}

// IconNotNil applies the NotNil predicate on the "icon" field.
func IconNotNil() predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIcon)))
	})
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIcon), v))
	})
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIcon), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// AssetIDEQ applies the EQ predicate on the "asset_id" field.
func AssetIDEQ(v pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssetID), v))
	})
}

// AssetIDNEQ applies the NEQ predicate on the "asset_id" field.
func AssetIDNEQ(v pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAssetID), v))
	})
}

// AssetIDIn applies the In predicate on the "asset_id" field.
func AssetIDIn(vs ...pulid.PULID) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAssetID), v...))
	})
}

// AssetIDNotIn applies the NotIn predicate on the "asset_id" field.
func AssetIDNotIn(vs ...pulid.PULID) predicate.Cryptocurrency {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAssetID), v...))
	})
}

// AssetIDGT applies the GT predicate on the "asset_id" field.
func AssetIDGT(v pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAssetID), v))
	})
}

// AssetIDGTE applies the GTE predicate on the "asset_id" field.
func AssetIDGTE(v pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAssetID), v))
	})
}

// AssetIDLT applies the LT predicate on the "asset_id" field.
func AssetIDLT(v pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAssetID), v))
	})
}

// AssetIDLTE applies the LTE predicate on the "asset_id" field.
func AssetIDLTE(v pulid.PULID) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAssetID), v))
	})
}

// AssetIDContains applies the Contains predicate on the "asset_id" field.
func AssetIDContains(v pulid.PULID) predicate.Cryptocurrency {
	vc := string(v)
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAssetID), vc))
	})
}

// AssetIDHasPrefix applies the HasPrefix predicate on the "asset_id" field.
func AssetIDHasPrefix(v pulid.PULID) predicate.Cryptocurrency {
	vc := string(v)
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAssetID), vc))
	})
}

// AssetIDHasSuffix applies the HasSuffix predicate on the "asset_id" field.
func AssetIDHasSuffix(v pulid.PULID) predicate.Cryptocurrency {
	vc := string(v)
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAssetID), vc))
	})
}

// AssetIDEqualFold applies the EqualFold predicate on the "asset_id" field.
func AssetIDEqualFold(v pulid.PULID) predicate.Cryptocurrency {
	vc := string(v)
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAssetID), vc))
	})
}

// AssetIDContainsFold applies the ContainsFold predicate on the "asset_id" field.
func AssetIDContainsFold(v pulid.PULID) predicate.Cryptocurrency {
	vc := string(v)
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAssetID), vc))
	})
}

// HasAsset applies the HasEdge predicate on the "asset" edge.
func HasAsset() predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssetTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, AssetTable, AssetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetWith applies the HasEdge predicate on the "asset" edge with a given conditions (other predicates).
func HasAssetWith(preds ...predicate.Asset) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssetInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, AssetTable, AssetColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBlockchains applies the HasEdge predicate on the "blockchains" edge.
func HasBlockchains() predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlockchainsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, BlockchainsTable, BlockchainsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlockchainsWith applies the HasEdge predicate on the "blockchains" edge with a given conditions (other predicates).
func HasBlockchainsWith(preds ...predicate.Blockchain) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlockchainsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, BlockchainsTable, BlockchainsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Cryptocurrency) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Cryptocurrency) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
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
func Not(p predicate.Cryptocurrency) predicate.Cryptocurrency {
	return predicate.Cryptocurrency(func(s *sql.Selector) {
		p(s.Not())
	})
}
