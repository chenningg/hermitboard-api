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
	return []ent.Field{
		field.String("asset_class_id").
			GoType(pulid.PULID("")),
	}
}

// Edges of the Asset.
func (Asset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("asset_class", AssetClass.Type).
			Ref("assets").
			Field("asset_class_id").
			Unique().
			Required().
			Annotations(entgql.MapsTo("assetClass")),
		edge.To("cryptocurrency", Cryptocurrency.Type).
			Unique().
			Annotations(entgql.MapsTo("cryptocurrency")),
		edge.To("transaction_bases", Transaction.Type).
			Annotations(entgql.RelayConnection(), entgql.MapsTo("transactionBases")),
		edge.To("transaction_quotes", Transaction.Type).
			Annotations(entgql.RelayConnection(), entgql.MapsTo("transactionQuotes")),
		edge.To("daily_asset_prices", DailyAssetPrice.Type).
			Annotations(entgql.RelayConnection(), entgql.MapsTo("dailyAssetPrices")),
	}
}

func (Asset) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
