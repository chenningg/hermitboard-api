package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
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
			Optional().
			Nillable().
			NotEmpty().
			Annotations(entgql.Skip(entgql.SkipType | entgql.SkipWhereInput)),
		field.Time("password_updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entgql.OrderField("PASSWORD_UPDATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("auth_roles", AuthRole.Type).
			Required().
			Annotations(
				entgql.RelayConnection(), entgql.MapsTo("authRoles"),
			),
		edge.To("portfolios", Portfolio.Type).
			Annotations(
				entgql.RelayConnection(), entgql.MapsTo("portfolios"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
		edge.To("auth_type", AuthType.Type).
			Required().
			Unique().
			Annotations(
				entgql.MapsTo("authType"),
				entgql.Skip(entgql.SkipMutationCreateInput),
			),
		edge.To("connections", Connection.Type).
			Annotations(
				entgql.RelayConnection(), entgql.MapsTo("connections"),
				entgql.Skip(entgql.SkipWhereInput|entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
	}
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
