// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chenningg/hermitboard-api/ent/transactiontype"
	"github.com/chenningg/hermitboard-api/pulid"
)

// TransactionType is the model entity for the TransactionType schema.
type TransactionType struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.PULID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// Value holds the value of the "value" field.
	Value transactiontype.Value `json:"value,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TransactionType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transactiontype.FieldID:
			values[i] = new(pulid.PULID)
		case transactiontype.FieldValue, transactiontype.FieldDescription:
			values[i] = new(sql.NullString)
		case transactiontype.FieldCreatedAt, transactiontype.FieldUpdatedAt, transactiontype.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TransactionType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TransactionType fields.
func (tt *TransactionType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transactiontype.FieldID:
			if value, ok := values[i].(*pulid.PULID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tt.ID = *value
			}
		case transactiontype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tt.CreatedAt = value.Time
			}
		case transactiontype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tt.UpdatedAt = value.Time
			}
		case transactiontype.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				tt.DeletedAt = new(time.Time)
				*tt.DeletedAt = value.Time
			}
		case transactiontype.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				tt.Value = transactiontype.Value(value.String)
			}
		case transactiontype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				tt.Description = new(string)
				*tt.Description = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TransactionType.
// Note that you need to call TransactionType.Unwrap() before calling this method if this TransactionType
// was returned from a transaction, and the transaction was committed or rolled back.
func (tt *TransactionType) Update() *TransactionTypeUpdateOne {
	return (&TransactionTypeClient{config: tt.config}).UpdateOne(tt)
}

// Unwrap unwraps the TransactionType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tt *TransactionType) Unwrap() *TransactionType {
	_tx, ok := tt.config.driver.(*txDriver)
	if !ok {
		panic("ent: TransactionType is not a transactional entity")
	}
	tt.config.driver = _tx.drv
	return tt
}

// String implements the fmt.Stringer.
func (tt *TransactionType) String() string {
	var builder strings.Builder
	builder.WriteString("TransactionType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tt.ID))
	builder.WriteString("created_at=")
	builder.WriteString(tt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(tt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := tt.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", tt.Value))
	builder.WriteString(", ")
	if v := tt.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// TransactionTypes is a parsable slice of TransactionType.
type TransactionTypes []*TransactionType

func (tt TransactionTypes) config(cfg config) {
	for _i := range tt {
		tt[_i].config = cfg
	}
}
