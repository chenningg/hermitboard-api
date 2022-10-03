package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// SourceType holds the schema definition for the SourceType entity.
type SourceType struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (SourceType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("SCT"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the SourceType.
func (SourceType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("value").
			NamedValues(
				"Exchange", "EXCHANGE",
				"Bank", "BANK",
				"DecentralizedExchange", "DECENTRALIZED_EXCHANGE",
			).
			Annotations(
				entgql.Type("SourceTypeValue"),
				entgql.OrderField("SOURCE_TYPE_VALUE"),
			),
		field.String("description").
			Optional().
			Nillable().
			NotEmpty(),
	}
}

// Edges of the SourceType.
func (SourceType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sources", Source.Type).
			Annotations(
				entgql.RelayConnection(), entgql.MapsTo("sources"),
				entgql.Skip(entgql.SkipWhereInput|entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
	}
}

func (SourceType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("sourceTypes"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
