// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/connection"
	"github.com/chenningg/hermitboard-api/ent/source"
	"github.com/chenningg/hermitboard-api/pulid"
)

// Connection is the model entity for the Connection schema.
type Connection struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.PULID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// AccessToken holds the value of the "access_token" field.
	AccessToken string `json:"accessToken,omitempty"`
	// RefreshToken holds the value of the "refresh_token" field.
	RefreshToken *string `json:"refreshToken,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID pulid.PULID `json:"accountID,omitempty"`
	// SourceID holds the value of the "source_id" field.
	SourceID pulid.PULID `json:"sourceID,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConnectionQuery when eager-loading is set.
	Edges ConnectionEdges `json:"edges"`
}

// ConnectionEdges holds the relations/edges for other nodes in the graph.
type ConnectionEdges struct {
	// Source holds the value of the source edge.
	Source *Source `json:"source,omitempty"`
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// Portfolios holds the value of the portfolios edge.
	Portfolios []*Portfolio `json:"portfolios,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedPortfolios map[string][]*Portfolio
}

// SourceOrErr returns the Source value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConnectionEdges) SourceOrErr() (*Source, error) {
	if e.loadedTypes[0] {
		if e.Source == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: source.Label}
		}
		return e.Source, nil
	}
	return nil, &NotLoadedError{edge: "source"}
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConnectionEdges) AccountOrErr() (*Account, error) {
	if e.loadedTypes[1] {
		if e.Account == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// PortfoliosOrErr returns the Portfolios value or an error if the edge
// was not loaded in eager-loading.
func (e ConnectionEdges) PortfoliosOrErr() ([]*Portfolio, error) {
	if e.loadedTypes[2] {
		return e.Portfolios, nil
	}
	return nil, &NotLoadedError{edge: "portfolios"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Connection) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case connection.FieldID, connection.FieldAccountID, connection.FieldSourceID:
			values[i] = new(pulid.PULID)
		case connection.FieldName, connection.FieldAccessToken, connection.FieldRefreshToken:
			values[i] = new(sql.NullString)
		case connection.FieldCreatedAt, connection.FieldUpdatedAt, connection.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Connection", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Connection fields.
func (c *Connection) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case connection.FieldID:
			if value, ok := values[i].(*pulid.PULID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case connection.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case connection.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case connection.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = new(time.Time)
				*c.DeletedAt = value.Time
			}
		case connection.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case connection.FieldAccessToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_token", values[i])
			} else if value.Valid {
				c.AccessToken = value.String
			}
		case connection.FieldRefreshToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token", values[i])
			} else if value.Valid {
				c.RefreshToken = new(string)
				*c.RefreshToken = value.String
			}
		case connection.FieldAccountID:
			if value, ok := values[i].(*pulid.PULID); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value != nil {
				c.AccountID = *value
			}
		case connection.FieldSourceID:
			if value, ok := values[i].(*pulid.PULID); !ok {
				return fmt.Errorf("unexpected type %T for field source_id", values[i])
			} else if value != nil {
				c.SourceID = *value
			}
		}
	}
	return nil
}

// QuerySource queries the "source" edge of the Connection entity.
func (c *Connection) QuerySource() *SourceQuery {
	return (&ConnectionClient{config: c.config}).QuerySource(c)
}

// QueryAccount queries the "account" edge of the Connection entity.
func (c *Connection) QueryAccount() *AccountQuery {
	return (&ConnectionClient{config: c.config}).QueryAccount(c)
}

// QueryPortfolios queries the "portfolios" edge of the Connection entity.
func (c *Connection) QueryPortfolios() *PortfolioQuery {
	return (&ConnectionClient{config: c.config}).QueryPortfolios(c)
}

// Update returns a builder for updating this Connection.
// Note that you need to call Connection.Unwrap() before calling this method if this Connection
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Connection) Update() *ConnectionUpdateOne {
	return (&ConnectionClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Connection entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Connection) Unwrap() *Connection {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Connection is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Connection) String() string {
	var builder strings.Builder
	builder.WriteString("Connection(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := c.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("access_token=")
	builder.WriteString(c.AccessToken)
	builder.WriteString(", ")
	if v := c.RefreshToken; v != nil {
		builder.WriteString("refresh_token=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", c.AccountID))
	builder.WriteString(", ")
	builder.WriteString("source_id=")
	builder.WriteString(fmt.Sprintf("%v", c.SourceID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedPortfolios returns the Portfolios named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Connection) NamedPortfolios(name string) ([]*Portfolio, error) {
	if c.Edges.namedPortfolios == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedPortfolios[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Connection) appendNamedPortfolios(name string, edges ...*Portfolio) {
	if c.Edges.namedPortfolios == nil {
		c.Edges.namedPortfolios = make(map[string][]*Portfolio)
	}
	if len(edges) == 0 {
		c.Edges.namedPortfolios[name] = []*Portfolio{}
	} else {
		c.Edges.namedPortfolios[name] = append(c.Edges.namedPortfolios[name], edges...)
	}
}

// Connections is a parsable slice of Connection.
type Connections []*Connection

func (c Connections) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
