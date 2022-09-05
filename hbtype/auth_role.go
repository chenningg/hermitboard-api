//go:generate go run github.com/dmarkham/enumer -type=AuthRole -json -text -sql

package hbtype

import (
	"fmt"
)

type AuthRole int

const (
	Demo AuthRole = iota
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
