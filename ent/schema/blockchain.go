package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// Blockchain holds the schema definition for the Blockchain entity.
type Blockchain struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Blockchain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("BKC"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the Blockchain.
func (Blockchain) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("symbol").
			NotEmpty(),
		field.String("icon").
			Optional().
			Nillable().
			NotEmpty(),
		field.Int64("chain_id").
			Optional().
			Nillable(),
	}
}

// Edges of the Blockchain.
func (Blockchain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cryptocurrencies", Cryptocurrency.Type).
			Through("blockchain_cryptocurrencies", BlockchainCryptocurrency.Type),
	}
}
