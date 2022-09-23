package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// DailyAssetPrice holds the schema definition for the DailyAssetPrice entity.
type DailyAssetPrice struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (DailyAssetPrice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("DAP"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the DailyAssetPrice.
func (DailyAssetPrice) Fields() []ent.Field {
	return []ent.Field{
		field.Time("time").
			Default(time.Now).
			Annotations(entgql.OrderField("TIME")),
		field.Float("open").
			Optional().
			Nillable(),
		field.Float("high").
			Optional().
			Nillable(),
		field.Float("low").
			Optional().
			Nillable(),
		field.Float("close").
			Optional().
			Nillable(),
		field.Float("adjusted_close"),
	}
}

// Edges of the DailyAssetPrice.
func (DailyAssetPrice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("asset", Asset.Type).
			Unique().
			Ref("daily_asset_prices").
			Annotations(entgql.MapsTo("asset")),
	}
}

func (DailyAssetPrice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
