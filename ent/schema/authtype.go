package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// AuthType holds the schema definition for the AuthType entity.
type AuthType struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (AuthType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("ASC"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the AuthType.
func (AuthType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("auth_type").
			NamedValues(
				"Local", "LOCAL",
				"Google", "GOOGLE",
				"Apple", "APPLE",
				"Facebook", "FACEBOOK",
			).
			Default("LOCAL").
			Annotations(
				entgql.OrderField("AUTH_TYPE"),
			),
		field.String("description").
			Optional().
			Nillable().
			NotEmpty(),
	}
}

// Edges of the AuthType.
func (AuthType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
			Unique(),
		edge.To("staff_account", StaffAccount.Type).
			Unique(),
	}
}
