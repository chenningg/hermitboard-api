// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chenningg/hermitboard-api/ent/assetclass"
	"github.com/chenningg/hermitboard-api/pulid"
)

// AssetClass is the model entity for the AssetClass schema.
type AssetClass struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.PULID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Value holds the value of the "value" field.
	Value assetclass.Value `json:"value,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AssetClass) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case assetclass.FieldID:
			values[i] = new(pulid.PULID)
		case assetclass.FieldValue, assetclass.FieldDescription:
			values[i] = new(sql.NullString)
		case assetclass.FieldCreatedAt, assetclass.FieldUpdatedAt, assetclass.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AssetClass", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AssetClass fields.
func (ac *AssetClass) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case assetclass.FieldID:
			if value, ok := values[i].(*pulid.PULID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ac.ID = *value
			}
		case assetclass.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ac.CreatedAt = value.Time
			}
		case assetclass.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ac.UpdatedAt = value.Time
			}
		case assetclass.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ac.DeletedAt = new(time.Time)
				*ac.DeletedAt = value.Time
			}
		case assetclass.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				ac.Value = assetclass.Value(value.String)
			}
		case assetclass.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ac.Description = new(string)
				*ac.Description = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AssetClass.
// Note that you need to call AssetClass.Unwrap() before calling this method if this AssetClass
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AssetClass) Update() *AssetClassUpdateOne {
	return (&AssetClassClient{config: ac.config}).UpdateOne(ac)
}

// Unwrap unwraps the AssetClass entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AssetClass) Unwrap() *AssetClass {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AssetClass is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AssetClass) String() string {
	var builder strings.Builder
	builder.WriteString("AssetClass(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ac.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ac.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := ac.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", ac.Value))
	builder.WriteString(", ")
	if v := ac.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// AssetClasses is a parsable slice of AssetClass.
type AssetClasses []*AssetClass

func (ac AssetClasses) config(cfg config) {
	for _i := range ac {
		ac[_i].config = cfg
	}
}
