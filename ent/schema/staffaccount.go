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

// StaffAccount holds the schema definition for the StaffAccount entity.
type StaffAccount struct {
	ent.Schema
}

// Mixin adds id, created_at and updated_at fields.
func (StaffAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("SAC"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the StaffAccount.
func (StaffAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname").
			Unique().
			NotEmpty(),
		field.String("email").
			Unique().
			NotEmpty(),
		field.Bool("email_confirmed").
			Default(false).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		field.String("password").
			Sensitive().
			Optional().
			Nillable().
			NotEmpty().
			Annotations(entgql.Skip(entgql.SkipType | entgql.SkipWhereInput)),
		field.Time("password_updated_at").
			UpdateDefault(time.Now).
			Optional().
			Nillable().
			Annotations(
				entgql.OrderField("PASSWORD_UPDATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
	}
}

// Edges of the StaffAccount.
func (StaffAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("auth_roles", AuthRole.Type).
			Annotations(
				entgql.RelayConnection(), entgql.MapsTo("authRoles"),
			),
		edge.To("auth_type", AuthType.Type).
			Required().
			Unique().
			Annotations(
				entgql.MapsTo("authType"),
			),
	}
}

func (StaffAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("staffAccounts"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
