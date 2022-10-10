// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"github.com/chenningg/hermitboard-api/pulid"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.PULID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// EmailConfirmed holds the value of the "email_confirmed" field.
	EmailConfirmed bool `json:"emailConfirmed,omitempty"`
	// Password holds the value of the "password" field.
	Password *string `json:"-"`
	// PasswordUpdatedAt holds the value of the "password_updated_at" field.
	PasswordUpdatedAt *time.Time `json:"passwordUpdatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges             AccountEdges `json:"edges"`
	account_auth_type *pulid.PULID
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Friends holds the value of the friends edge.
	Friends []*Account `json:"friends,omitempty"`
	// AuthRoles holds the value of the auth_roles edge.
	AuthRoles []*AuthRole `json:"authRoles,omitempty"`
	// Portfolios holds the value of the portfolios edge.
	Portfolios []*Portfolio `json:"portfolios,omitempty"`
	// AuthType holds the value of the auth_type edge.
	AuthType *AuthType `json:"authType,omitempty"`
	// Connections holds the value of the connections edge.
	Connections []*Connection `json:"connections,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
	// totalCount holds the count of the edges above.
	totalCount [5]map[string]int

	namedFriends     map[string][]*Account
	namedAuthRoles   map[string][]*AuthRole
	namedPortfolios  map[string][]*Portfolio
	namedConnections map[string][]*Connection
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) FriendsOrErr() ([]*Account, error) {
	if e.loadedTypes[0] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// AuthRolesOrErr returns the AuthRoles value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) AuthRolesOrErr() ([]*AuthRole, error) {
	if e.loadedTypes[1] {
		return e.AuthRoles, nil
	}
	return nil, &NotLoadedError{edge: "auth_roles"}
}

// PortfoliosOrErr returns the Portfolios value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) PortfoliosOrErr() ([]*Portfolio, error) {
	if e.loadedTypes[2] {
		return e.Portfolios, nil
	}
	return nil, &NotLoadedError{edge: "portfolios"}
}

// AuthTypeOrErr returns the AuthType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) AuthTypeOrErr() (*AuthType, error) {
	if e.loadedTypes[3] {
		if e.AuthType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: authtype.Label}
		}
		return e.AuthType, nil
	}
	return nil, &NotLoadedError{edge: "auth_type"}
}

// ConnectionsOrErr returns the Connections value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) ConnectionsOrErr() ([]*Connection, error) {
	if e.loadedTypes[4] {
		return e.Connections, nil
	}
	return nil, &NotLoadedError{edge: "connections"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			values[i] = new(pulid.PULID)
		case account.FieldEmailConfirmed:
			values[i] = new(sql.NullBool)
		case account.FieldNickname, account.FieldEmail, account.FieldPassword:
			values[i] = new(sql.NullString)
		case account.FieldCreatedAt, account.FieldUpdatedAt, account.FieldDeletedAt, account.FieldPasswordUpdatedAt:
			values[i] = new(sql.NullTime)
		case account.ForeignKeys[0]: // account_auth_type
			values[i] = &sql.NullScanner{S: new(pulid.PULID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Account", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			if value, ok := values[i].(*pulid.PULID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case account.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case account.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case account.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = new(time.Time)
				*a.DeletedAt = value.Time
			}
		case account.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				a.Nickname = value.String
			}
		case account.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				a.Email = value.String
			}
		case account.FieldEmailConfirmed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_confirmed", values[i])
			} else if value.Valid {
				a.EmailConfirmed = value.Bool
			}
		case account.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				a.Password = new(string)
				*a.Password = value.String
			}
		case account.FieldPasswordUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field password_updated_at", values[i])
			} else if value.Valid {
				a.PasswordUpdatedAt = new(time.Time)
				*a.PasswordUpdatedAt = value.Time
			}
		case account.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field account_auth_type", values[i])
			} else if value.Valid {
				a.account_auth_type = new(pulid.PULID)
				*a.account_auth_type = *value.S.(*pulid.PULID)
			}
		}
	}
	return nil
}

// QueryFriends queries the "friends" edge of the Account entity.
func (a *Account) QueryFriends() *AccountQuery {
	return (&AccountClient{config: a.config}).QueryFriends(a)
}

// QueryAuthRoles queries the "auth_roles" edge of the Account entity.
func (a *Account) QueryAuthRoles() *AuthRoleQuery {
	return (&AccountClient{config: a.config}).QueryAuthRoles(a)
}

// QueryPortfolios queries the "portfolios" edge of the Account entity.
func (a *Account) QueryPortfolios() *PortfolioQuery {
	return (&AccountClient{config: a.config}).QueryPortfolios(a)
}

// QueryAuthType queries the "auth_type" edge of the Account entity.
func (a *Account) QueryAuthType() *AuthTypeQuery {
	return (&AccountClient{config: a.config}).QueryAuthType(a)
}

// QueryConnections queries the "connections" edge of the Account entity.
func (a *Account) QueryConnections() *ConnectionQuery {
	return (&AccountClient{config: a.config}).QueryConnections(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := a.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(a.Nickname)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(a.Email)
	builder.WriteString(", ")
	builder.WriteString("email_confirmed=")
	builder.WriteString(fmt.Sprintf("%v", a.EmailConfirmed))
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	if v := a.PasswordUpdatedAt; v != nil {
		builder.WriteString("password_updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedFriends returns the Friends named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Account) NamedFriends(name string) ([]*Account, error) {
	if a.Edges.namedFriends == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedFriends[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Account) appendNamedFriends(name string, edges ...*Account) {
	if a.Edges.namedFriends == nil {
		a.Edges.namedFriends = make(map[string][]*Account)
	}
	if len(edges) == 0 {
		a.Edges.namedFriends[name] = []*Account{}
	} else {
		a.Edges.namedFriends[name] = append(a.Edges.namedFriends[name], edges...)
	}
}

// NamedAuthRoles returns the AuthRoles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Account) NamedAuthRoles(name string) ([]*AuthRole, error) {
	if a.Edges.namedAuthRoles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedAuthRoles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Account) appendNamedAuthRoles(name string, edges ...*AuthRole) {
	if a.Edges.namedAuthRoles == nil {
		a.Edges.namedAuthRoles = make(map[string][]*AuthRole)
	}
	if len(edges) == 0 {
		a.Edges.namedAuthRoles[name] = []*AuthRole{}
	} else {
		a.Edges.namedAuthRoles[name] = append(a.Edges.namedAuthRoles[name], edges...)
	}
}

// NamedPortfolios returns the Portfolios named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Account) NamedPortfolios(name string) ([]*Portfolio, error) {
	if a.Edges.namedPortfolios == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedPortfolios[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Account) appendNamedPortfolios(name string, edges ...*Portfolio) {
	if a.Edges.namedPortfolios == nil {
		a.Edges.namedPortfolios = make(map[string][]*Portfolio)
	}
	if len(edges) == 0 {
		a.Edges.namedPortfolios[name] = []*Portfolio{}
	} else {
		a.Edges.namedPortfolios[name] = append(a.Edges.namedPortfolios[name], edges...)
	}
}

// NamedConnections returns the Connections named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Account) NamedConnections(name string) ([]*Connection, error) {
	if a.Edges.namedConnections == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedConnections[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Account) appendNamedConnections(name string, edges ...*Connection) {
	if a.Edges.namedConnections == nil {
		a.Edges.namedConnections = make(map[string][]*Connection)
	}
	if len(edges) == 0 {
		a.Edges.namedConnections[name] = []*Connection{}
	} else {
		a.Edges.namedConnections[name] = append(a.Edges.namedConnections[name], edges...)
	}
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
