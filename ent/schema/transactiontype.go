package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/schema/mixin"
)

// TransactionType holds the schema definition for the TransactionType entity.
type TransactionType struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (TransactionType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.PULIDMixinWithPrefix("TRT"),
		mixin.CreatedAtMixin{},
		mixin.UpdatedAtMixin{},
		mixin.DeletedAtMixin{},
	}
}

// Fields of the TransactionType.
func (TransactionType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("value").
			NamedValues(
				"Buy", "BUY",
				"Sell", "SELL",
				"Swap", "SWAP",
				"Stake", "STAKE",
				"DividendIncome", "DIVIDEND_INCOME",
				"RentPayment", "RENT_PAYMENT",
				"RentIncome", "RENT_INCOME",
				"StockDividend", "STOCK_DIVIDEND",
			).
			Annotations(
				entgql.Type("TransactionTypeValue"),
				entgql.OrderField("TRANSACTION_TYPE"),
			),
		field.String("description").
			Optional().
			Nillable().
			NotEmpty(),
	}
}

// Edges of the TransactionType.
func (TransactionType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transactions", Transaction.Type).
			Annotations(entgql.RelayConnection(), entgql.MapsTo("transactions")),
	}
}

func (TransactionType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
