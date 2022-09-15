// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/accountauthrole"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/schema/pulid"
)

// AccountAuthRole is the model entity for the AccountAuthRole schema.
type AccountAuthRole struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.PULID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID pulid.PULID `json:"account_id,omitempty"`
	// AuthRoleID holds the value of the "auth_role_id" field.
	AuthRoleID pulid.PULID `json:"auth_role_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountAuthRoleQuery when eager-loading is set.
	Edges AccountAuthRoleEdges `json:"edges"`
}

// AccountAuthRoleEdges holds the relations/edges for other nodes in the graph.
type AccountAuthRoleEdges struct {
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// AuthRole holds the value of the auth_role edge.
	AuthRole *AuthRole `json:"auth_role,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountAuthRoleEdges) AccountOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Account == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// AuthRoleOrErr returns the AuthRole value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountAuthRoleEdges) AuthRoleOrErr() (*AuthRole, error) {
	if e.loadedTypes[1] {
		if e.AuthRole == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: authrole.Label}
		}
		return e.AuthRole, nil
	}
	return nil, &NotLoadedError{edge: "auth_role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccountAuthRole) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case accountauthrole.FieldID, accountauthrole.FieldAccountID, accountauthrole.FieldAuthRoleID:
			values[i] = new(pulid.PULID)
		case accountauthrole.FieldCreatedAt, accountauthrole.FieldUpdatedAt, accountauthrole.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AccountAuthRole", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccountAuthRole fields.
func (aar *AccountAuthRole) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accountauthrole.FieldID:
			if value, ok := values[i].(*pulid.PULID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				aar.ID = *value
			}
		case accountauthrole.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				aar.CreatedAt = value.Time
			}
		case accountauthrole.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				aar.UpdatedAt = value.Time
			}
		case accountauthrole.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				aar.DeletedAt = new(time.Time)
				*aar.DeletedAt = value.Time
			}
		case accountauthrole.FieldAccountID:
			if value, ok := values[i].(*pulid.PULID); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value != nil {
				aar.AccountID = *value
			}
		case accountauthrole.FieldAuthRoleID:
			if value, ok := values[i].(*pulid.PULID); !ok {
				return fmt.Errorf("unexpected type %T for field auth_role_id", values[i])
			} else if value != nil {
				aar.AuthRoleID = *value
			}
		}
	}
	return nil
}

// QueryAccount queries the "account" edge of the AccountAuthRole entity.
func (aar *AccountAuthRole) QueryAccount() *AccountQuery {
	return (&AccountAuthRoleClient{config: aar.config}).QueryAccount(aar)
}

// QueryAuthRole queries the "auth_role" edge of the AccountAuthRole entity.
func (aar *AccountAuthRole) QueryAuthRole() *AuthRoleQuery {
	return (&AccountAuthRoleClient{config: aar.config}).QueryAuthRole(aar)
}

// Update returns a builder for updating this AccountAuthRole.
// Note that you need to call AccountAuthRole.Unwrap() before calling this method if this AccountAuthRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (aar *AccountAuthRole) Update() *AccountAuthRoleUpdateOne {
	return (&AccountAuthRoleClient{config: aar.config}).UpdateOne(aar)
}

// Unwrap unwraps the AccountAuthRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aar *AccountAuthRole) Unwrap() *AccountAuthRole {
	_tx, ok := aar.config.driver.(*txDriver)
	if !ok {
		panic("ent: AccountAuthRole is not a transactional entity")
	}
	aar.config.driver = _tx.drv
	return aar
}

// String implements the fmt.Stringer.
func (aar *AccountAuthRole) String() string {
	var builder strings.Builder
	builder.WriteString("AccountAuthRole(")
	builder.WriteString(fmt.Sprintf("id=%v, ", aar.ID))
	builder.WriteString("created_at=")
	builder.WriteString(aar.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(aar.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := aar.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", aar.AccountID))
	builder.WriteString(", ")
	builder.WriteString("auth_role_id=")
	builder.WriteString(fmt.Sprintf("%v", aar.AuthRoleID))
	builder.WriteByte(')')
	return builder.String()
}

// AccountAuthRoles is a parsable slice of AccountAuthRole.
type AccountAuthRoles []*AccountAuthRole

func (aar AccountAuthRoles) config(cfg config) {
	for _i := range aar {
		aar[_i].config = cfg
	}
}
