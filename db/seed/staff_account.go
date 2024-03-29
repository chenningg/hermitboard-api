package seed

import (
	"context"
	"fmt"
	"time"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"golang.org/x/crypto/bcrypt"
)

func seedStaffAccounts(ctx context.Context, client *ent.Client) error {
	// Create the first admin account
	ownerPassword, err := bcrypt.GenerateFromPassword([]byte("password"), 12)
	if err != nil {
		return fmt.Errorf("seedStaffAccounts(): failed to generate password for owner account")
	}

	// Retrieve super admin role reference from database
	superAdminRole, err := client.AuthRole.Query().Where(authrole.ValueEQ(authrole.ValueSuperAdmin)).Only(ctx)
	if err != nil {
		return fmt.Errorf("seedStaffAccounts(): could not find super admin role while seeding")
	}

	// Retrieve local authentication type reference from database
	localAuthType, err := client.AuthType.Query().Where(authtype.ValueEQ(authtype.ValueLocal)).Only(ctx)
	if err != nil {
		return fmt.Errorf("seedStaffAccounts(): could not find local auth type while seeding")
	}

	_, err = client.StaffAccount.Create().
		SetID("SAC_01GEGJGGJHXB9FWZ84SPKCVFG8").
		SetAuthType(localAuthType).
		SetNickname("owner").
		SetEmail("owner@hermitboard.com").
		SetPassword(string(ownerPassword)).
		SetPasswordUpdatedAt(time.Now()).
		SetEmailConfirmed(true).
		AddAuthRoles(superAdminRole).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("seedStaffAccounts(): failed to create owner account")
	}

	return nil
}
