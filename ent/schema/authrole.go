package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// AuthRole holds the schema definition for the AuthRole entity.
type AuthRole struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (AuthRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("ATR"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the AuthRole.
func (AuthRole) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("value").
			NamedValues(
				"Demo", "DEMO",
				"Free", "FREE",
				"Plus", "PLUS",
				"Pro", "PRO",
				"Enterprise", "ENTERPRISE",
				"Support", "SUPPORT",
				"Admin", "ADMIN",
				"SuperAdmin", "SUPER_ADMIN",
			).
			Annotations(
				entgql.Type("AuthRoleValue"),
				entgql.OrderField("AUTH_ROLE"),
			),
		field.String("description").
			Optional().
			Nillable().
			NotEmpty(),
	}
}

// Edges of the AuthRole.
func (AuthRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("accounts", Account.Type).
			Ref("auth_roles").
			Annotations(entgql.RelayConnection(), entgql.MapsTo("accounts")),
		edge.From("staff_accounts", StaffAccount.Type).
			Ref("auth_roles").
			Annotations(entgql.RelayConnection(), entgql.MapsTo("staffAccounts")),
	}
}

func (AuthRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
