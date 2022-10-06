package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
	"github.com/chenningg/hermitboard-api/pulid"
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
		field.Float("price_per_unit").
			StructTag("json:\"pricePerUnit,omitempty\""),
		field.String("blockchain_id").
			GoType(pulid.PULID("")).
			StructTag("json:\"blockchainID,omitempty\"").
			Optional().
			Nillable(),
		field.String("exchange_id").
			StructTag("json:\"exchangeID,omitempty\"").
			GoType(pulid.PULID("")),
		field.String("portfolio_id").
			StructTag("json:\"portfolioID,omitempty\"").
			GoType(pulid.PULID("")),
		field.String("base_asset_id").
			StructTag("json:\"baseAssetID,omitempty\"").
			GoType(pulid.PULID("")),
		field.String("quote_asset_id").
			StructTag("json:\"quoteAssetID,omitempty\"").
			GoType(pulid.PULID("")).
			Optional().
			Nillable(),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transaction_type", TransactionType.Type).
			Unique().
			Required().
			StructTag("json:\"transactionType,omitempty\"").
			Annotations(entgql.MapsTo("transactionType")),
		edge.To("base_asset", Asset.Type).
			Field("base_asset_id").
			Required().
			Unique().
			StructTag("json:\"baseAsset,omitempty\"").
			Annotations(entgql.MapsTo("baseAsset")),
		edge.To("quote_asset", Asset.Type).
			Field("quote_asset_id").
			Unique().
			StructTag("json:\"quoteAsset,omitempty\"").
			Annotations(entgql.MapsTo("quoteAsset")),
		edge.From("portfolio", Portfolio.Type).
			Ref("transactions").
			Field("portfolio_id").
			Unique().
			Required().
			Annotations(entgql.MapsTo("portfolio")),
		edge.From("exchange", Exchange.Type).
			Ref("transactions").
			Field("exchange_id").
			Required().
			Unique().
			Annotations(entgql.MapsTo("exchange")),
		edge.From("blockchain", Blockchain.Type).
			Ref("transactions").
			Field("blockchain_id").
			Unique().
			Annotations(entgql.MapsTo("blockchain")),
	}
}

func (Transaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("transactions"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
