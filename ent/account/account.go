// Code generated by ent, DO NOT EDIT.

package account

import (
	"time"

	"github.com/chenningg/hermitboard-api/pulid"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldEmailConfirmed holds the string denoting the email_confirmed field in the database.
	FieldEmailConfirmed = "email_confirmed"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldPasswordUpdatedAt holds the string denoting the password_updated_at field in the database.
	FieldPasswordUpdatedAt = "password_updated_at"
	// FieldProfilePictureURL holds the string denoting the profile_picture_url field in the database.
	FieldProfilePictureURL = "profile_picture_url"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeAuthRoles holds the string denoting the auth_roles edge name in mutations.
	EdgeAuthRoles = "auth_roles"
	// EdgePortfolios holds the string denoting the portfolios edge name in mutations.
	EdgePortfolios = "portfolios"
	// EdgeAuthType holds the string denoting the auth_type edge name in mutations.
	EdgeAuthType = "auth_type"
	// EdgeConnections holds the string denoting the connections edge name in mutations.
	EdgeConnections = "connections"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "account_friends"
	// AuthRolesTable is the table that holds the auth_roles relation/edge. The primary key declared below.
	AuthRolesTable = "account_auth_roles"
	// AuthRolesInverseTable is the table name for the AuthRole entity.
	// It exists in this package in order to avoid circular dependency with the "authrole" package.
	AuthRolesInverseTable = "auth_roles"
	// PortfoliosTable is the table that holds the portfolios relation/edge.
	PortfoliosTable = "portfolios"
	// PortfoliosInverseTable is the table name for the Portfolio entity.
	// It exists in this package in order to avoid circular dependency with the "portfolio" package.
	PortfoliosInverseTable = "portfolios"
	// PortfoliosColumn is the table column denoting the portfolios relation/edge.
	PortfoliosColumn = "account_id"
	// AuthTypeTable is the table that holds the auth_type relation/edge.
	AuthTypeTable = "accounts"
	// AuthTypeInverseTable is the table name for the AuthType entity.
	// It exists in this package in order to avoid circular dependency with the "authtype" package.
	AuthTypeInverseTable = "auth_types"
	// AuthTypeColumn is the table column denoting the auth_type relation/edge.
	AuthTypeColumn = "account_auth_type"
	// ConnectionsTable is the table that holds the connections relation/edge.
	ConnectionsTable = "connections"
	// ConnectionsInverseTable is the table name for the Connection entity.
	// It exists in this package in order to avoid circular dependency with the "connection" package.
	ConnectionsInverseTable = "connections"
	// ConnectionsColumn is the table column denoting the connections relation/edge.
	ConnectionsColumn = "account_id"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldNickname,
	FieldEmail,
	FieldEmailConfirmed,
	FieldPassword,
	FieldPasswordUpdatedAt,
	FieldProfilePictureURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "accounts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_auth_type",
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"account_id", "friend_id"}
	// AuthRolesPrimaryKey and AuthRolesColumn2 are the table columns denoting the
	// primary key for the auth_roles relation (M2M).
	AuthRolesPrimaryKey = []string{"account_id", "auth_role_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultEmailConfirmed holds the default value on creation for the "email_confirmed" field.
	DefaultEmailConfirmed bool
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// UpdateDefaultPasswordUpdatedAt holds the default value on update for the "password_updated_at" field.
	UpdateDefaultPasswordUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.PULID
)
