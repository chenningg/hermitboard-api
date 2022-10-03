package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
			NotEmpty().
			Annotations(entgql.OrderField("NAME")),
		field.Bool("is_public").
			Default(false).
			Comment("Whether this portfolio is visible to others.").
			Annotations(entgql.OrderField("IS_PUBLIC")),
		field.Bool("is_visible").
			Default(true).
			Comment("Whether this portfolio is visible to the owner.").
			Annotations(
				entgql.OrderField("IS_VISIBLE"),
			),
		field.String("account_id").
			GoType(pulid.PULID("")).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the Portfolio.
func (Portfolio) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("portfolios").
			Field("account_id").
			Unique().
			Required().
			Annotations(
				entgql.MapsTo("account"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
		edge.To("transactions", Transaction.Type).
			Annotations(
				entgql.RelayConnection(), entgql.MapsTo("transactions"),
				entgql.Skip(entgql.SkipWhereInput|entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
		edge.To("connections", Connection.Type).
			Annotations(entgql.RelayConnection(), entgql.MapsTo("connections")),
	}
}

func (Portfolio) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("portfolios"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Portfolio) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_id", "name").
			Unique(),
	}
}
