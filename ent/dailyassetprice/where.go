// Code generated by ent, DO NOT EDIT.

package dailyassetprice

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// ID filters vertices based on their ID field.
func ID(id pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTime), v))
	})
}

// Open applies equality check predicate on the "open" field. It's identical to OpenEQ.
func Open(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOpen), v))
	})
}

// High applies equality check predicate on the "high" field. It's identical to HighEQ.
func High(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHigh), v))
	})
}

// Low applies equality check predicate on the "low" field. It's identical to LowEQ.
func Low(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLow), v))
	})
}

// Close applies equality check predicate on the "close" field. It's identical to CloseEQ.
func Close(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClose), v))
	})
}

// AdjustedClose applies equality check predicate on the "adjusted_close" field. It's identical to AdjustedCloseEQ.
func AdjustedClose(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdjustedClose), v))
	})
}

// BaseAssetID applies equality check predicate on the "base_asset_id" field. It's identical to BaseAssetIDEQ.
func BaseAssetID(v pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaseAssetID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTime), v))
	})
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTime), v))
	})
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...time.Time) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTime), v...))
	})
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...time.Time) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTime), v...))
	})
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTime), v))
	})
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTime), v))
	})
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTime), v))
	})
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v time.Time) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTime), v))
	})
}

// OpenEQ applies the EQ predicate on the "open" field.
func OpenEQ(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOpen), v))
	})
}

// OpenNEQ applies the NEQ predicate on the "open" field.
func OpenNEQ(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOpen), v))
	})
}

// OpenIn applies the In predicate on the "open" field.
func OpenIn(vs ...float64) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOpen), v...))
	})
}

// OpenNotIn applies the NotIn predicate on the "open" field.
func OpenNotIn(vs ...float64) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOpen), v...))
	})
}

// OpenGT applies the GT predicate on the "open" field.
func OpenGT(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOpen), v))
	})
}

// OpenGTE applies the GTE predicate on the "open" field.
func OpenGTE(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOpen), v))
	})
}

// OpenLT applies the LT predicate on the "open" field.
func OpenLT(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOpen), v))
	})
}

// OpenLTE applies the LTE predicate on the "open" field.
func OpenLTE(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOpen), v))
	})
}

// OpenIsNil applies the IsNil predicate on the "open" field.
func OpenIsNil() predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOpen)))
	})
}

// OpenNotNil applies the NotNil predicate on the "open" field.
func OpenNotNil() predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOpen)))
	})
}

// HighEQ applies the EQ predicate on the "high" field.
func HighEQ(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHigh), v))
	})
}

// HighNEQ applies the NEQ predicate on the "high" field.
func HighNEQ(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHigh), v))
	})
}

// HighIn applies the In predicate on the "high" field.
func HighIn(vs ...float64) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHigh), v...))
	})
}

// HighNotIn applies the NotIn predicate on the "high" field.
func HighNotIn(vs ...float64) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHigh), v...))
	})
}

// HighGT applies the GT predicate on the "high" field.
func HighGT(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHigh), v))
	})
}

// HighGTE applies the GTE predicate on the "high" field.
func HighGTE(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHigh), v))
	})
}

// HighLT applies the LT predicate on the "high" field.
func HighLT(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHigh), v))
	})
}

// HighLTE applies the LTE predicate on the "high" field.
func HighLTE(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHigh), v))
	})
}

// HighIsNil applies the IsNil predicate on the "high" field.
func HighIsNil() predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHigh)))
	})
}

// HighNotNil applies the NotNil predicate on the "high" field.
func HighNotNil() predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHigh)))
	})
}

// LowEQ applies the EQ predicate on the "low" field.
func LowEQ(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLow), v))
	})
}

// LowNEQ applies the NEQ predicate on the "low" field.
func LowNEQ(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLow), v))
	})
}

// LowIn applies the In predicate on the "low" field.
func LowIn(vs ...float64) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLow), v...))
	})
}

// LowNotIn applies the NotIn predicate on the "low" field.
func LowNotIn(vs ...float64) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLow), v...))
	})
}

// LowGT applies the GT predicate on the "low" field.
func LowGT(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLow), v))
	})
}

// LowGTE applies the GTE predicate on the "low" field.
func LowGTE(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLow), v))
	})
}

// LowLT applies the LT predicate on the "low" field.
func LowLT(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLow), v))
	})
}

// LowLTE applies the LTE predicate on the "low" field.
func LowLTE(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLow), v))
	})
}

// LowIsNil applies the IsNil predicate on the "low" field.
func LowIsNil() predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLow)))
	})
}

// LowNotNil applies the NotNil predicate on the "low" field.
func LowNotNil() predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLow)))
	})
}

// CloseEQ applies the EQ predicate on the "close" field.
func CloseEQ(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClose), v))
	})
}

// CloseNEQ applies the NEQ predicate on the "close" field.
func CloseNEQ(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClose), v))
	})
}

// CloseIn applies the In predicate on the "close" field.
func CloseIn(vs ...float64) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClose), v...))
	})
}

// CloseNotIn applies the NotIn predicate on the "close" field.
func CloseNotIn(vs ...float64) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClose), v...))
	})
}

// CloseGT applies the GT predicate on the "close" field.
func CloseGT(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClose), v))
	})
}

// CloseGTE applies the GTE predicate on the "close" field.
func CloseGTE(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClose), v))
	})
}

// CloseLT applies the LT predicate on the "close" field.
func CloseLT(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClose), v))
	})
}

// CloseLTE applies the LTE predicate on the "close" field.
func CloseLTE(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClose), v))
	})
}

// CloseIsNil applies the IsNil predicate on the "close" field.
func CloseIsNil() predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldClose)))
	})
}

// CloseNotNil applies the NotNil predicate on the "close" field.
func CloseNotNil() predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldClose)))
	})
}

// AdjustedCloseEQ applies the EQ predicate on the "adjusted_close" field.
func AdjustedCloseEQ(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdjustedClose), v))
	})
}

// AdjustedCloseNEQ applies the NEQ predicate on the "adjusted_close" field.
func AdjustedCloseNEQ(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdjustedClose), v))
	})
}

// AdjustedCloseIn applies the In predicate on the "adjusted_close" field.
func AdjustedCloseIn(vs ...float64) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAdjustedClose), v...))
	})
}

// AdjustedCloseNotIn applies the NotIn predicate on the "adjusted_close" field.
func AdjustedCloseNotIn(vs ...float64) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAdjustedClose), v...))
	})
}

// AdjustedCloseGT applies the GT predicate on the "adjusted_close" field.
func AdjustedCloseGT(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAdjustedClose), v))
	})
}

// AdjustedCloseGTE applies the GTE predicate on the "adjusted_close" field.
func AdjustedCloseGTE(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAdjustedClose), v))
	})
}

// AdjustedCloseLT applies the LT predicate on the "adjusted_close" field.
func AdjustedCloseLT(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAdjustedClose), v))
	})
}

// AdjustedCloseLTE applies the LTE predicate on the "adjusted_close" field.
func AdjustedCloseLTE(v float64) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAdjustedClose), v))
	})
}

// BaseAssetIDEQ applies the EQ predicate on the "base_asset_id" field.
func BaseAssetIDEQ(v pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaseAssetID), v))
	})
}

// BaseAssetIDNEQ applies the NEQ predicate on the "base_asset_id" field.
func BaseAssetIDNEQ(v pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaseAssetID), v))
	})
}

// BaseAssetIDIn applies the In predicate on the "base_asset_id" field.
func BaseAssetIDIn(vs ...pulid.PULID) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBaseAssetID), v...))
	})
}

// BaseAssetIDNotIn applies the NotIn predicate on the "base_asset_id" field.
func BaseAssetIDNotIn(vs ...pulid.PULID) predicate.DailyAssetPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBaseAssetID), v...))
	})
}

// BaseAssetIDGT applies the GT predicate on the "base_asset_id" field.
func BaseAssetIDGT(v pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaseAssetID), v))
	})
}

// BaseAssetIDGTE applies the GTE predicate on the "base_asset_id" field.
func BaseAssetIDGTE(v pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaseAssetID), v))
	})
}

// BaseAssetIDLT applies the LT predicate on the "base_asset_id" field.
func BaseAssetIDLT(v pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaseAssetID), v))
	})
}

// BaseAssetIDLTE applies the LTE predicate on the "base_asset_id" field.
func BaseAssetIDLTE(v pulid.PULID) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaseAssetID), v))
	})
}

// BaseAssetIDContains applies the Contains predicate on the "base_asset_id" field.
func BaseAssetIDContains(v pulid.PULID) predicate.DailyAssetPrice {
	vc := string(v)
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBaseAssetID), vc))
	})
}

// BaseAssetIDHasPrefix applies the HasPrefix predicate on the "base_asset_id" field.
func BaseAssetIDHasPrefix(v pulid.PULID) predicate.DailyAssetPrice {
	vc := string(v)
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBaseAssetID), vc))
	})
}

// BaseAssetIDHasSuffix applies the HasSuffix predicate on the "base_asset_id" field.
func BaseAssetIDHasSuffix(v pulid.PULID) predicate.DailyAssetPrice {
	vc := string(v)
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBaseAssetID), vc))
	})
}

// BaseAssetIDEqualFold applies the EqualFold predicate on the "base_asset_id" field.
func BaseAssetIDEqualFold(v pulid.PULID) predicate.DailyAssetPrice {
	vc := string(v)
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBaseAssetID), vc))
	})
}

// BaseAssetIDContainsFold applies the ContainsFold predicate on the "base_asset_id" field.
func BaseAssetIDContainsFold(v pulid.PULID) predicate.DailyAssetPrice {
	vc := string(v)
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBaseAssetID), vc))
	})
}

// HasBaseAsset applies the HasEdge predicate on the "base_asset" edge.
func HasBaseAsset() predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaseAssetTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BaseAssetTable, BaseAssetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaseAssetWith applies the HasEdge predicate on the "base_asset" edge with a given conditions (other predicates).
func HasBaseAssetWith(preds ...predicate.Asset) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaseAssetInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BaseAssetTable, BaseAssetColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DailyAssetPrice) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DailyAssetPrice) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
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
func Not(p predicate.DailyAssetPrice) predicate.DailyAssetPrice {
	return predicate.DailyAssetPrice(func(s *sql.Selector) {
		p(s.Not())
	})
}