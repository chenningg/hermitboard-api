package mixin

import (
	"time"

	"entgo.io/contrib/entgql"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type DeletedAtMixin struct {
	mixin.Schema
}

func (DeletedAtMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Optional().
			Nillable().
			UpdateDefault(time.Now).
			StructTag("json:\"deletedAt,omitempty\"").
			Annotations(
				entgql.OrderField("DELETED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput),
			),
	}
}
