package auth

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/staffaccount"
	"golang.org/x/crypto/bcrypt"
)

func (authService AuthService) LoginToLocalAccount(ctx context.Context, username string, password string) (
	*ent.Account, error,
) {
	// Username can be either email or nickname.
	dbClient := ent.FromContext(ctx)

	accounts, err := dbClient.Account.Query().Where(
		account.Or(
			account.NicknameEQ(username), account.EmailEQ(username),
		),
	).All(ctx)
	if err != nil {
		return nil, err
	}

	for _, acc := range accounts {
		// Compare passwords.
		hashedPassword := acc.Password
		if hashedPassword == nil {
			continue
		}
		err = bcrypt.CompareHashAndPassword([]byte(*acc.Password), []byte(password))
		if err == nil {
			return acc, nil
		}
	}

	// If we reach here, no staff accounts match the given username and password. Return an error.
	return nil, fmt.Errorf("incorrect username or password provided")
}

func (authService AuthService) LoginToLocalStaffAccount(
	ctx context.Context, username string, password string,
) (*ent.StaffAccount, error) {
	// Username can be either email or nickname.
	dbClient := ent.FromContext(ctx)

	staffAccounts, err := dbClient.StaffAccount.Query().Where(
		staffaccount.Or(
			staffaccount.NicknameEQ(username), staffaccount.EmailEQ(username),
		),
	).All(ctx)
	if err != nil {
		return nil, err
	}

	for _, staffAccount := range staffAccounts {
		// Compare passwords.
		hashedPassword := staffAccount.Password
		if hashedPassword == nil {
			continue
		}
		err = bcrypt.CompareHashAndPassword([]byte(*staffAccount.Password), []byte(password))
		if err == nil {
			return staffAccount, nil
		}
	}

	// If we reach here, no staff accounts match the given username and password. Return an error.
	return nil, fmt.Errorf("incorrect username or password provided")
}
