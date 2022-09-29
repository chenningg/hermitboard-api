package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
			NotEmpty().
			Annotations(entgql.OrderField("NAME")),
		field.String("symbol").
			NotEmpty().
			Annotations(entgql.OrderField("SYMBOL")),
		field.String("icon").
			Optional().
			Nillable().
			NotEmpty().
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Int64("chain_id").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("CHAIN_ID")),
	}
}

// Edges of the Blockchain.
func (Blockchain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cryptocurrencies", Cryptocurrency.Type).
			Annotations(
				entgql.RelayConnection(), entgql.MapsTo("cryptocurrencies"),
				entgql.Skip(entgql.SkipWhereInput|entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
		edge.To("transactions", Transaction.Type).
			Annotations(
				entgql.RelayConnection(), entgql.MapsTo("transactions"),
				entgql.Skip(entgql.SkipWhereInput|entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
	}
}

func (Blockchain) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
