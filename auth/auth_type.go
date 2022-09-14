//go:generate go run github.com/dmarkham/enumer -type=AuthType -json -text -sql

package auth

import (
	"fmt"
	"io"
	"strconv"
)

type AuthType int

const (
	Local AuthType = iota + 1
	Firebase
)

func (authType *AuthType) Validate() error {
	if authType.IsAAuthType() {
		return nil
	}

	return fmt.Errorf("invalid value for AuthType, expected %v, got %v", AuthTypeValues(), authType.String())
}

func (authType AuthType) Values() []string {
	return AuthTypeStrings()
}

// Marshals a AuthType into a graphql scalar string.
func (authType AuthType) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(pulid.String()))
}

// Unmarshals a graphql scalar into a AuthType Go type.
func (authType *AuthType) UnmarshalGQL(v interface{}) error {
	return pulid.Scan(v)
}
