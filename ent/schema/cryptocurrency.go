package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
	"github.com/chenningg/hermitboard-api/pulid"
)

// Cryptocurrency holds the schema definition for the Cryptocurrency entity.
type Cryptocurrency struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Cryptocurrency) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("CRP"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the Cryptocurrency.
func (Cryptocurrency) Fields() []ent.Field {
	return []ent.Field{
		field.String("symbol").
			NotEmpty().
			Annotations(entgql.OrderField("SYMBOL")),
		field.String("icon").
			Optional().
			Nillable().
			NotEmpty().
			Comment("A url to the image icon for this cryptocurrency.").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.String("name").
			NotEmpty().
			Annotations(entgql.OrderField("NAME")),
		field.String("asset_id").
			GoType(pulid.PULID("")),
	}
}

// Edges of the Cryptocurrency.
func (Cryptocurrency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("asset", Asset.Type).
			Ref("cryptocurrency").
			Field("asset_id").
			Unique().
			Required().
			Annotations(entgql.MapsTo("asset")),
		edge.From("blockchains", Blockchain.Type).
			Ref("cryptocurrencies").
			Required().
			Annotations(entgql.RelayConnection(), entgql.MapsTo("blockchains")),
	}
}

func (Cryptocurrency) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("cryptocurrencies"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
