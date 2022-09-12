// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chenningg/hermitboard-api/ent/exchange"
	"github.com/oklog/ulid/v2"
	ulid "github.com/oklog/ulid/v2"
)

// Exchange is the model entity for the Exchange schema.
type Exchange struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ULID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon *string `json:"icon,omitempty"`
	// A url to the exchange site.
	URL *string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExchangeQuery when eager-loading is set.
	Edges ExchangeEdges `json:"edges"`
}

// ExchangeEdges holds the relations/edges for other nodes in the graph.
type ExchangeEdges struct {
	// Transactions holds the value of the transactions edge.
	Transactions []*Transaction `json:"transactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TransactionsOrErr returns the Transactions value or an error if the edge
// was not loaded in eager-loading.
func (e ExchangeEdges) TransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[0] {
		return e.Transactions, nil
	}
	return nil, &NotLoadedError{edge: "transactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Exchange) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case exchange.FieldName, exchange.FieldIcon, exchange.FieldURL:
			values[i] = new(sql.NullString)
		case exchange.FieldCreatedAt, exchange.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case exchange.FieldID:
			values[i] = new(ulid.ULID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Exchange", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Exchange fields.
func (e *Exchange) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case exchange.FieldID:
			if value, ok := values[i].(*ulid.ULID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case exchange.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case exchange.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case exchange.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case exchange.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				e.Icon = new(string)
				*e.Icon = value.String
			}
		case exchange.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				e.URL = new(string)
				*e.URL = value.String
			}
		}
	}
	return nil
}

// QueryTransactions queries the "transactions" edge of the Exchange entity.
func (e *Exchange) QueryTransactions() *TransactionQuery {
	return (&ExchangeClient{config: e.config}).QueryTransactions(e)
}

// Update returns a builder for updating this Exchange.
// Note that you need to call Exchange.Unwrap() before calling this method if this Exchange
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Exchange) Update() *ExchangeUpdateOne {
	return (&ExchangeClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Exchange entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Exchange) Unwrap() *Exchange {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Exchange is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Exchange) String() string {
	var builder strings.Builder
	builder.WriteString("Exchange(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	if v := e.Icon; v != nil {
		builder.WriteString("icon=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := e.URL; v != nil {
		builder.WriteString("url=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Exchanges is a parsable slice of Exchange.
type Exchanges []*Exchange

func (e Exchanges) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
