package schema

import (
	"entgo.io/contrib/entgql"
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
		mixin.PULIDMixinWithPrefix("ACC"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname").
			Unique().
			NotEmpty().
			MinLen(3).
			Annotations(entgql.OrderField("NICKNAME")),
		field.String("email").
			Unique().
			NotEmpty().
			Annotations(entgql.OrderField("EMAIL")),
		field.String("password").
			Sensitive().
			Comment("Hashed and salted password using Bcrypt.").
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
			Required().
			Annotations(entgql.MapsTo("authRoles")),
		edge.To("portfolios", Portfolio.Type).
			Annotations(entgql.MapsTo("portfolios")),
		edge.To("auth_type", AuthType.Type).
			Required().
			Unique().
			Annotations(entgql.MapsTo("authType")),
	}
}

// Check that either auth_id or password is not null.
func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		&entsql.Annotation{
			Checks: map[string]string{
				// Checks that if auth_type is LOCAL, password field cannot be NULL.
				"account_chk_if_auth_type_local_then_password_not_null": "(auth_type <> 'LOCAL') OR (password IS NOT NULL)",
			},
		},
	}
}
