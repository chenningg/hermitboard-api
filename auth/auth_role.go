//go:generate go run github.com/dmarkham/enumer -type=AuthRole -json -text -sql

package auth

import (
	"fmt"
	"io"
	"strconv"
)

type AuthRole int

const (
	Demo AuthRole = iota + 1
	Free
	Plus
	Pro
	Enterprise
	Support
	Admin
	SuperAdmin
)

func (authRole *AuthRole) Validate() error {
	if authRole.IsAAuthRole() {
		return nil
	}

	return fmt.Errorf("invalid value for AuthRole, expected %v, got %v", AuthRoleValues(), authRole.String())
}

func (authRole AuthRole) Values() []string {
	return AuthRoleStrings()
}

// Marshals a AuthRole into a graphql scalar string.
func (authRole AuthRole) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(pulid.String()))
}

// Unmarshals a graphql scalar into a AuthRole Go type.
func (authRole *AuthRole) UnmarshalGQL(v interface{}) error {
	return pulid.Scan(v)
}
