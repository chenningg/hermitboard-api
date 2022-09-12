package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// DailyAssetPrice holds the schema definition for the DailyAssetPrice entity.
type DailyAssetPrice struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (DailyAssetPrice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the DailyAssetPrice.
func (DailyAssetPrice) Fields() []ent.Field {
	return []ent.Field{
		field.Time("time").
			Default(time.Now),
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
		edge.From("base_asset", Asset.Type).
			Ref("daily_asset_price").
			Unique().
			Required(),
	}
}

func (DailyAssetPrice) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("time").
			Edges("base_asset").
			Unique(),
	}
}
