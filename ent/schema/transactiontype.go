package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
	"github.com/chenningg/hermitboard-api/hbtype"
)

// TransactionType holds the schema definition for the TransactionType entity.
type TransactionType struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (TransactionType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the TransactionType.
func (TransactionType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			GoType(hbtype.TransactionType(1)),
		field.String("description").
			Optional().
			Nillable().
			NotEmpty(),
	}
}

// Edges of the TransactionType.
func (TransactionType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("transactions", Transaction.Type).
			Ref("transaction_type"),
	}
}
