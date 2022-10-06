package mixin

import (
	"time"

	"entgo.io/contrib/entgql"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type UpdatedAtMixin struct {
	mixin.Schema
}

func (UpdatedAtMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			StructTag("json:\"updatedAt,omitempty\"").
			Annotations(
				entgql.OrderField("UPDATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
	}
}
