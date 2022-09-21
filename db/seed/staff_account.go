package seed

import (
	"context"
	"fmt"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/staffaccount"
	"golang.org/x/crypto/bcrypt"
)

func seedStaffAccounts(ctx context.Context, client *ent.Client) error {
	// Create the first admin account
	ownerPassword, err := bcrypt.GenerateFromPassword([]byte("owner"), 10)
	if err != nil {
		return fmt.Errorf("seedStaffAccounts(): failed to generate password for owner account")
	}

	// Retrieve super admin role reference from database
	superAdminRole, err := client.AuthRole.Query().Where(authrole.AuthRoleEQ(authrole.AuthRoleSuperAdmin)).Only(ctx)
	if err != nil {
		return fmt.Errorf("seedStaffAccounts(): could not find super admin role for seeding")
	}

	_, err = client.StaffAccount.Create().
		SetAuthType(staffaccount.AuthTypeLocal).
		SetNickname("owner").
		SetEmail("owner@hermitboard.com").
		SetPassword(string(ownerPassword)).
		AddAuthRoles(superAdminRole).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("seedAuthRoles(): failed to create owner account")
	}

	return nil
}
