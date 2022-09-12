package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("auth_id").
			Optional().
			Nillable().
			Unique().
			NotEmpty(),
		field.String("nickname").
			Unique().
			NotEmpty(),
		field.String("email").
			Unique().
			NotEmpty(),
		field.String("password").
			Optional().
			Nillable().
			NotEmpty(),
		field.Time("password_updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("auth_roles", AuthRole.Type).
			Required(),
		edge.To("portfolios", Portfolio.Type),
	}
}

// Check that either auth_id or password is not null.
func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		&entsql.Annotation{
			Checks: map[string]string{
				"account_chk_auth_id_password_not_null": "auth_id IS NOT NULL OR password IS NOT NULL",
			},
		},
	}
}
