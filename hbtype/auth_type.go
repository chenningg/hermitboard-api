//go:generate go run github.com/dmarkham/enumer -type=AuthType

package hbtype

import (
	"fmt"
)

type AuthType int

const (
	Local AuthType = iota
	Firebase
)

func (authType *AuthType) Validate() error {
	if authType.IsAAuthType() {
		return nil
	}

	return fmt.Errorf("invalid value for AuthType, expected %v, got %v", AuthTypeValues(), authType.String())
}

func (authType *AuthType) UnmarshalText(text []byte) error {
	input := string(text)
	marshalledAuthType, err := AuthTypeString(input)
	// If unable to parse, use a default value.
	if err != nil {
		return fmt.Errorf("invalid value for AuthType, expected one of %v, got '%v'", AuthTypeValues(), input)
	}

	*authType = marshalledAuthType

	return nil
}
