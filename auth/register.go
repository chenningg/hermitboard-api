package auth

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/ent"
)

func (authService AuthService) CreateAccount(ctx context.Context, input ent.CreateAccountInput) (
	*ent.Account, error,
) {
	dbClient := ent.FromContext(ctx)

	acc, err := dbClient.Account.Create().SetEmail(input.Email).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not register account: %v", err)
	}

	return acc, nil
}

func (authService AuthService) CreateStaffAccount(
	ctx context.Context, input ent.CreateStaffAccountInput,
) (*ent.StaffAccount, error) {
	return nil, nil
}
