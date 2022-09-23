package mixin

import (
	"entgo.io/contrib/entgql"
	"time"

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
			Annotations(entgql.OrderField("UPDATED_AT")),
	}
}
