package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
	"github.com/chenningg/hermitboard-api/hbtype"
)

// AssetClass holds the schema definition for the AssetClass entity.
type AssetClass struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (AssetClass) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the AssetClass.
func (AssetClass) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("class").
			GoType(hbtype.AssetClass(1)),
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
