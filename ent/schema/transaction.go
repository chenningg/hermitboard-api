package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
	"github.com/chenningg/hermitboard-api/ent/schema/pulid"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Transaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("TRX"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Time("time"),
		field.Int("units"),
		field.Float("price_per_unit"),
		field.String("transaction_type_id").
			GoType(pulid.PULID("")),
		field.String("base_asset_id").
			GoType(pulid.PULID("")),
		field.String("quote_asset_id").
			GoType(pulid.PULID("")).
			Optional().
			Nillable(),
		field.String("portfolio_id").
			GoType(pulid.PULID("")),
		field.String("exchange_id").
			GoType(pulid.PULID("")),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transaction_type", TransactionType.Type).
			Unique().
			Required().
			Field("transaction_type_id"),
		edge.To("base_asset", Asset.Type).
			Required().
			Unique().
			Field("base_asset_id"),
		edge.To("quote_asset", Asset.Type).
			Unique().
			Field("quote_asset_id"),
		edge.From("portfolio", Portfolio.Type).
			Ref("transactions").
			Unique().
			Required().
			Field("portfolio_id"),
		edge.From("exchange", Exchange.Type).
			Ref("transactions").
			Required().
			Unique().
			Field("exchange_id"),
	}
}
