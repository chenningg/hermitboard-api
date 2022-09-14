package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Transaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("TRX"),
		mixin.TimeMixin{},
	}
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Time("time"),
		field.Int("units"),
		field.Float("price_per_unit"),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transaction_type", TransactionType.Type).
			Unique().
			Required(),
		edge.To("base_asset", Asset.Type).
			Required(),
		edge.To("quote_asset", Asset.Type),
		edge.From("portfolio", Portfolio.Type).
			Ref("transactions").
			Unique().
			Required(),
		edge.To("exchange", Exchange.Type).
			Unique(),
	}
}

func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("time").
			Edges("portfolio"),
	}
}
