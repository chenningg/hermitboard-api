// Code generated by ent, DO NOT EDIT.

package transactiontype

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/chenningg/hermitboard-api/pulid"
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
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldTransactionType holds the string denoting the transaction_type field in the database.
	FieldTransactionType = "transaction_type"
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
	FieldDeletedAt,
	FieldTransactionType,
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
	// UpdateDefaultDeletedAt holds the default value on update for the "deleted_at" field.
	UpdateDefaultDeletedAt func() time.Time
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.PULID
)

// TransactionType defines the type for the "transaction_type" enum field.
type TransactionType string

// TransactionType values.
const (
	TransactionTypeBuy            TransactionType = "BUY"
	TransactionTypeSell           TransactionType = "SELL"
	TransactionTypeStake          TransactionType = "STAKE"
	TransactionTypeDividendIncome TransactionType = "DIVIDEND_INCOME"
	TransactionTypeRentPayment    TransactionType = "RENT_PAYMENT"
	TransactionTypeRentIncome     TransactionType = "RENT_INCOME"
	TransactionTypeStockDividend  TransactionType = "STOCK_DIVIDEND"
)

func (tt TransactionType) String() string {
	return string(tt)
}

// TransactionTypeValidator is a validator for the "transaction_type" field enum values. It is called by the builders before save.
func TransactionTypeValidator(tt TransactionType) error {
	switch tt {
	case TransactionTypeBuy, TransactionTypeSell, TransactionTypeStake, TransactionTypeDividendIncome, TransactionTypeRentPayment, TransactionTypeRentIncome, TransactionTypeStockDividend:
		return nil
	default:
		return fmt.Errorf("transactiontype: invalid enum value for transaction_type field: %q", tt)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e TransactionType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *TransactionType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = TransactionType(str)
	if err := TransactionTypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid TransactionType", str)
	}
	return nil
}
