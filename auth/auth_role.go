package auth

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type AuthRole string

const (
	Demo       AuthRole = "DEMO"
	Free       AuthRole = "FREE"
	Plus       AuthRole = "PLUS"
	Pro        AuthRole = "PRO"
	Enterprise AuthRole = "ENTERPRISE"
	Support    AuthRole = "SUPPORT"
	Admin      AuthRole = "ADMIN"
	SuperAdmin AuthRole = "SUPER_ADMIN"
)

// Validates that the enum is a valid value, returns an error if not.
func (authRole AuthRole) Validate() error {
	switch authRole {
	case Demo, Free, Plus, Pro, Enterprise, Support, Admin, SuperAdmin:
		return nil
	}

	return fmt.Errorf("invalid value for AuthRole, expected %v, got %v", authRole.Values(), authRole.String())
}

// Returns a slice of all values of this enum.
func (authRole AuthRole) Values() []string {
	return []string{"DEMO", "FREE", "PLUS", "PRO", "ENTERPRISE", "SUPPORT", "ADMIN", "SUPER_ADMIN"}
}

// Converts a string value to the enum.
func (authRole *AuthRole) StringToEnum(src string) error {
	switch s := strings.ToUpper(src); s {
	case "DEMO":
		*authRole = AuthRole("DEMO")
	case "FREE":
		*authRole
	}
}

// Converts the enum to a string value.
func (authRole AuthRole) String() string {
	return string(authRole)
}

// Marshals into a graphql scalar string.
func (authRole AuthRole) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(authRole.String()))
}

// Unmarshals a graphql scalar into a Go type.
func (authRole *AuthRole) UnmarshalGQL(v interface{}) error {
	return authRole.Scan(v)
}

// Checks whether this value is a valid enum.
func (authRole AuthRole) IsValid() bool {
	err := authRole.Validate()
	if err != nil {
		return false
	}
	return true
}

// MarshalJSON implements the json.Marshaler interface
func (authRole AuthRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(authRole.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (authRole *AuthRole) UnmarshalJSON(src []byte) error {
	var s string
	if err := json.Unmarshal(src, &s); err != nil {
		return fmt.Errorf("AuthRole should be a string, got %s", src)
	}

	var err error
	*authRole, err = AuthTypeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface to turn a Go type to a []byte.
func (i AuthType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface to turn a []byte to a Go type.
func (i *AuthType) UnmarshalText(text []byte) error {
	var err error
	*i, err = AuthTypeString(string(text))
	return err
}

func (i AuthType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *AuthType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of AuthType: %[1]T(%[1]v)", value)
	}

	val, err := AuthTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
