package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// AssetClass holds the schema definition for the AssetClass entity.
type AssetClass struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (AssetClass) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("ASC"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the AssetClass.
func (AssetClass) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("asset_class").
			NamedValues(
				"CashOrCashEquivalent", "CASH_OR_CASH_EQUIVALENT",
				"Commodity", "COMMODITY",
				"Cryptocurrency", "CRYPTOCURRENCY",
				"Equity", "EQUITY",
				"FixedIncome", "FIXED_INCOME",
				"Future", "FUTURE",
				"RealEstate", "REAL_ESTATE",
			).
			Annotations(
				entgql.OrderField("ASSET_CLASS"),
			),
		field.String("description").
			Optional().
			Nillable().
			NotEmpty(),
	}
}

// Edges of the AssetClass.
func (AssetClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("assets", Asset.Type).
			Ref("asset_class"),
	}
}
