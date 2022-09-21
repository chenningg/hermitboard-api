package schema

import (
	"github.com/chenningg/hermitboard-api/pulid"
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
		field.String("base_asset_id").
			GoType(pulid.PULID("")),
	}
}

// Edges of the DailyAssetPrice.
func (DailyAssetPrice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("base_asset", Asset.Type).
			Ref("daily_asset_price").
			Unique().
			Required().
			Field("base_asset_id"),
	}
}

func (DailyAssetPrice) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("base_asset_id", "time").
			Unique(),
	}
}