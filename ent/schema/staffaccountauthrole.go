package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
	"github.com/chenningg/hermitboard-api/pulid"
)

// StaffAccountAuthRole holds the schema definition for the StaffAccountAuthRole entity.
type StaffAccountAuthRole struct {
	ent.Schema
}

// Mixin adds id, created_at and updated_at fields.
func (StaffAccountAuthRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("SAA"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the StaffAccountAuthRole.
func (StaffAccountAuthRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("staff_account_id").
			GoType(pulid.PULID("")),
		field.String("auth_role_id").
			GoType(pulid.PULID("")),
	}
}

// Edges of the StaffAccountAuthRole.
func (StaffAccountAuthRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("staff_account", StaffAccount.Type).
			Required().
			Unique().
			Field("staff_account_id"),
		edge.To("auth_role", AuthRole.Type).
			Required().
			Unique().
			Field("auth_role_id"),
	}
}
