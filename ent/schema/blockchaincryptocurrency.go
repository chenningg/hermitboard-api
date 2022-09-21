package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
	"github.com/chenningg/hermitboard-api/pulid"
)

// BlockchainCryptocurrency holds the schema definition for the BlockchainCryptocurrency entity.
type BlockchainCryptocurrency struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (BlockchainCryptocurrency) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("AAR"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the BlockchainCryptocurrency.
func (BlockchainCryptocurrency) Fields() []ent.Field {
	return []ent.Field{
		field.String("blockchain_id").
			GoType(pulid.PULID("")),
		field.String("cryptocurrency_id").
			GoType(pulid.PULID("")),
	}
}

// Edges of the BlockchainCryptocurrency.
func (BlockchainCryptocurrency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blockchain", Blockchain.Type).
			Required().
			Unique().
			Field("blockchain_id"),
		edge.To("cryptocurrency", Cryptocurrency.Type).
			Required().
			Unique().
			Field("cryptocurrency_id"),
	}
}
