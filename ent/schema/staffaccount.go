package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Enum("auth_type").
			NamedValues(
				"Local", "LOCAL",
				"Google", "GOOGLE",
				"Facebook", "FACEBOOK",
				"Apple", "APPLE",
			).
			Default("LOCAL").
			Annotations(
				entgql.OrderField("AUTH_TYPE"),
			),
		field.String("nickname").
			Unique().
			NotEmpty(),
		field.String("email").
			Unique().
			NotEmpty(),
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

// Edges of the StaffAccount.
func (StaffAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("auth_roles", AuthRole.Type).
			Through("staff_account_auth_roles", StaffAccountAuthRole.Type),
	}
}

func (StaffAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		&entsql.Annotation{
			Checks: map[string]string{
				// Checks that if auth_type is LOCAL, password field cannot be NULL.
				"staff_account_chk_if_auth_type_local_then_password_not_null": "(auth_type <> 'LOCAL') OR (password IS NOT NULL)",
			},
		},
	}
}