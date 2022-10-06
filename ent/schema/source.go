package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// Source holds the schema definition for the Source entity.
type Source struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Source) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("SRC"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the Source.
func (Source) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique().
			Annotations(entgql.OrderField("NAME")),
		field.String("icon").
			Optional().
			Nillable().
			NotEmpty().
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the Source.
func (Source) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("source_type", SourceType.Type).
			Ref("sources").
			StructTag("json:\"sourceType,omitempty\"").
			Required().
			Unique().
			Annotations(entgql.MapsTo("sourceType")),
	}
}

func (Source) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("sources"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
