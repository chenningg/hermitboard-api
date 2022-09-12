// Code generated by ent, DO NOT EDIT.

package transactiontype

import (
	"fmt"
	"time"

	"github.com/chenningg/hermitboard-api/hbtype"
	"github.com/oklog/ulid/v2"
	ulid "github.com/oklog/ulid/v2"
)

const (
	// Label holds the string label denoting the transactiontype type in the database.
	Label = "transaction_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeTransactions holds the string denoting the transactions edge name in mutations.
	EdgeTransactions = "transactions"
	// Table holds the table name of the transactiontype in the database.
	Table = "transaction_types"
	// TransactionsTable is the table that holds the transactions relation/edge.
	TransactionsTable = "transactions"
	// TransactionsInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionsInverseTable = "transactions"
	// TransactionsColumn is the table column denoting the transactions relation/edge.
	TransactionsColumn = "transaction_transaction_type"
)

// Columns holds all SQL columns for transactiontype fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldType,
	FieldDescription,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ULID
)

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type hbtype.TransactionType) error {
	switch _type.String() {
	case "Buy", "Sell", "Stake", "DividendIncome", "RentPayment", "RentIncome", "StockDividend":
		return nil
	default:
		return fmt.Errorf("transactiontype: invalid enum value for type field: %q", _type)
	}
}
