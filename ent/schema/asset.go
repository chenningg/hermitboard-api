package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// Asset holds the schema definition for the Asset entity.
type Asset struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Asset) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("AST"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the Asset.
func (Asset) Fields() []ent.Field {
	return nil
}

// Edges of the Asset.
func (Asset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("asset_class", AssetClass.Type).
			Unique().
			Required().
			Annotations(entgql.MapsTo("assetClass")),
		edge.To("cryptocurrency", Cryptocurrency.Type).
			Unique().
			Annotations(entgql.MapsTo("cryptocurrency")),
		edge.From("transaction_bases", Transaction.Type).
			Ref("base_asset").
			Annotations(entgql.MapsTo("transactionBases")),
		edge.From("transaction_quotes", Transaction.Type).
			Ref("quote_asset").
			Annotations(entgql.MapsTo("transactionQuotes")),
		edge.To("daily_asset_prices", DailyAssetPrice.Type).
			Annotations(entgql.MapsTo("dailyAssetPrices")),
	}
}

func (Asset) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
