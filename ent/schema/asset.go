package schema

import (
	"entgo.io/ent"
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
		mixin.TimeMixin{},
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
			Required(),
		edge.To("cryptocurrency", Cryptocurrency.Type).
			Unique(),
		edge.From("transaction_base", Transaction.Type).
			Ref("base_asset"),
		edge.From("transaction_quote", Transaction.Type).
			Ref("quote_asset"),
		edge.To("daily_asset_price", DailyAssetPrice.Type),
	}
}
