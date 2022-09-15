package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
	"github.com/chenningg/hermitboard-api/ent/schema/pulid"
)

// Cryptocurrency holds the schema definition for the Cryptocurrency entity.
type Cryptocurrency struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Cryptocurrency) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("CRC"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
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
		field.String("asset_id").
			GoType(pulid.PULID("")),
	}
}

// Edges of the Cryptocurrency.
func (Cryptocurrency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("asset", Asset.Type).
			Ref("cryptocurrency").
			Unique().
			Required().
			Field("asset_id"),
		edge.From("blockchains", Blockchain.Type).
			Ref("cryptocurrencies").
			Through("blockchain_cryptocurrencies", BlockchainCryptocurrency.Type),
	}
}
