package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// Exchange holds the schema definition for the Exchange entity.
type Exchange struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Exchange) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("EXG"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the Exchange.
func (Exchange) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("icon").
			Optional().
			Nillable().
			NotEmpty(),
		field.String("url").
			Optional().
			Nillable().
			NotEmpty().
			Comment("A url to the exchange site."),
	}
}

// Edges of the Exchange.
func (Exchange) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transactions", Transaction.Type),
	}
}