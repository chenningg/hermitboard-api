package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transaction_type", TransactionType.Type).
			Unique().
			Required().
			Annotations(entgql.MapsTo("transactionType")),
		edge.To("base_asset", Asset.Type).
			Required().
			Unique().
			Annotations(entgql.MapsTo("baseAsset")),
		edge.To("quote_asset", Asset.Type).
			Unique().
			Annotations(entgql.MapsTo("quoteAsset")),
		edge.From("portfolio", Portfolio.Type).
			Ref("transactions").
			Unique().
			Required().
			Annotations(entgql.MapsTo("portfolio")),
		edge.From("exchange", Exchange.Type).
			Ref("transactions").
			Required().
			Unique().
			Annotations(entgql.MapsTo("exchange")),
		edge.To("blockchain", Blockchain.Type).
			Unique().
			Annotations(entgql.MapsTo("blockchain")),
	}
}

func (Transaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
