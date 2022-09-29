// Code generated by ent, DO NOT EDIT.

package assetclass

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/chenningg/hermitboard-api/pulid"
)

const (
	// Label holds the string label denoting the assetclass type in the database.
	Label = "asset_class"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// Table holds the table name of the assetclass in the database.
	Table = "asset_classes"
)

// Columns holds all SQL columns for assetclass fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldValue,
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

// Value defines the type for the "value" enum field.
type Value string

// Value values.
const (
	ValueCashOrCashEquivalent Value = "CASH_OR_CASH_EQUIVALENT"
	ValueCommodity            Value = "COMMODITY"
	ValueCryptocurrency       Value = "CRYPTOCURRENCY"
	ValueEquity               Value = "EQUITY"
	ValueFixedIncome          Value = "FIXED_INCOME"
	ValueFuture               Value = "FUTURE"
	ValueRealEstate           Value = "REAL_ESTATE"
)

func (v Value) String() string {
	return string(v)
}

// ValueValidator is a validator for the "value" field enum values. It is called by the builders before save.
func ValueValidator(v Value) error {
	switch v {
	case ValueCashOrCashEquivalent, ValueCommodity, ValueCryptocurrency, ValueEquity, ValueFixedIncome, ValueFuture, ValueRealEstate:
		return nil
	default:
		return fmt.Errorf("assetclass: invalid enum value for value field: %q", v)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Value) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Value) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Value(str)
	if err := ValueValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Value", str)
	}
	return nil
}
