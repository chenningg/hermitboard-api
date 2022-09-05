//go:generate go run github.com/dmarkham/enumer -type=AuthType -json -text -sql

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
