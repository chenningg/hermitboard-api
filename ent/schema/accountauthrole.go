package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
	"github.com/chenningg/hermitboard-api/pulid"
)

// AccountAuthRole holds the schema definition for the AccountAuthRole entity.
type AccountAuthRole struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (AccountAuthRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("AAR"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the AccountAuthRole.
func (AccountAuthRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_id").
			GoType(pulid.PULID("")),
		field.String("auth_role_id").
			GoType(pulid.PULID("")),
	}
}

// Edges of the AccountAuthRole.
func (AccountAuthRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
			Required().
			Unique().
			Field("account_id"),
		edge.To("auth_role", AuthRole.Type).
			Required().
			Unique().
			Field("auth_role_id"),
	}
}
