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

// Connection holds the schema definition for the Connection entity.
type Connection struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (Connection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("CNN"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the Connection.
func (Connection) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MinLen(3).
			Annotations(entgql.OrderField("NAME")),
		field.String("access_token").
			NotEmpty().
			StructTag("json:\"accessToken,omitempty\"").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.String("refresh_token").
			Optional().
			Nillable().
			NotEmpty().
			StructTag("json:\"refreshToken,omitempty\"").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.String("account_id").
			GoType(pulid.PULID("")).
			StructTag("json:\"accountID,omitempty\"").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the Connection.
func (Connection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("connections").
			Field("account_id").
			Unique().
			Required().
			Annotations(
				entgql.MapsTo("account"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
		edge.From("portfolios", Portfolio.Type).
			Ref("connections").
			Annotations(
				entgql.RelayConnection(), entgql.MapsTo("portfolios"),
				entgql.Skip(entgql.SkipWhereInput),
			),
	}
}

func (Connection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("connections"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Connection) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_id", "name").
			Unique(),
	}
}
