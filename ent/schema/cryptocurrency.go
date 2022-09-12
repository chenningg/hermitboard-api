package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// Cryptocurrency holds the schema definition for the Cryptocurrency entity.
type Cryptocurrency struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Cryptocurrency) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the Cryptocurrency.
func (Cryptocurrency) Fields() []ent.Field {
	return []ent.Field{
		field.String("symbol").
			NotEmpty(),
		field.String("icon").
			Optional().
			Nillable().
			NotEmpty().
			Comment("A url to the image icon for this cryptocurrency."),
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the Cryptocurrency.
func (Cryptocurrency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("asset", Asset.Type).
			Ref("cryptocurrency").
			Unique().
			Required(),
		edge.From("blockchains", Blockchain.Type).
			Ref("cryptocurrencies"),
	}
}
