package mixin

import (
	"entgo.io/contrib/entgql"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CreatedAtMixin struct {
	mixin.Schema
}

func (CreatedAtMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),
	}
}
