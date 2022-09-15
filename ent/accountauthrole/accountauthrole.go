// Code generated by ent, DO NOT EDIT.

package accountauthrole

import (
	"time"

	"github.com/chenningg/hermitboard-api/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the accountauthrole type in the database.
	Label = "account_auth_role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldAuthRoleID holds the string denoting the auth_role_id field in the database.
	FieldAuthRoleID = "auth_role_id"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// EdgeAuthRole holds the string denoting the auth_role edge name in mutations.
	EdgeAuthRole = "auth_role"
	// Table holds the table name of the accountauthrole in the database.
	Table = "account_auth_roles"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "account_auth_roles"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_id"
	// AuthRoleTable is the table that holds the auth_role relation/edge.
	AuthRoleTable = "account_auth_roles"
	// AuthRoleInverseTable is the table name for the AuthRole entity.
	// It exists in this package in order to avoid circular dependency with the "authrole" package.
	AuthRoleInverseTable = "auth_roles"
	// AuthRoleColumn is the table column denoting the auth_role relation/edge.
	AuthRoleColumn = "auth_role_id"
)

// Columns holds all SQL columns for accountauthrole fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAccountID,
	FieldAuthRoleID,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.PULID
)
