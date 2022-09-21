package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
	"github.com/chenningg/hermitboard-api/pulid"
)

// Portfolio holds the schema definition for the Portfolio entity.
type Portfolio struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Portfolio) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("PTF"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the Portfolio.
func (Portfolio) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Bool("is_public").
			Default(false).
			Comment("Whether this portfolio is visible to others."),
		field.Bool("is_visible").
			Default(true).
			Comment("Whether this portfolio is visible to the owner."),
		field.String("account_id").
			GoType(pulid.PULID("")),
	}
}

// Edges of the Portfolio.
func (Portfolio) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("portfolios").
			Unique().
			Required().
			Field("account_id"),
		edge.To("transactions", Transaction.Type),
	}
}

func (Portfolio) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_id", "name").
			Unique(),
	}
}
