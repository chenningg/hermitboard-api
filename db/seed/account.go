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

func seedAccounts(ctx context.Context, client *ent.Client) error {
	// Create the first user account
	accPassword, err := bcrypt.GenerateFromPassword([]byte("password"), 12)
	if err != nil {
		return fmt.Errorf("seedAccounts(): failed to generate password for user account")
	}

	// Retrieve pro reference from database
	proRole, err := client.AuthRole.Query().Where(authrole.ValueEQ(authrole.ValuePro)).Only(ctx)
	if err != nil {
		return fmt.Errorf("seedAccounts(): could not find pro role while seeding")
	}

	// Retrieve local authentication type reference from database
	localAuthType, err := client.AuthType.Query().Where(authtype.ValueEQ(authtype.ValueLocal)).Only(ctx)
	if err != nil {
		return fmt.Errorf("seedAccounts(): could not find local auth type while seeding")
	}

	_, err = client.Account.Create().
		SetID("ACC_01GEGJGGJHXB9FWZ84SPKCVFG8").
		SetAuthType(localAuthType).
		SetNickname("user").
		SetEmail("user@hermitboard.com").
		SetPassword(string(accPassword)).
		SetPasswordUpdatedAt(time.Now()).
		SetEmailConfirmed(true).
		AddAuthRoles(proRole).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("seedAccounts(): failed to create user account")
	}

	return nil
}
