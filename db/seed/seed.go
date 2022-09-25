package seed

import (
	"context"
	"fmt"
	"github.com/chenningg/hermitboard-api/ent"
)

func Seed(ctx context.Context, client *ent.Client) error {
	// Seed enumerations first.
	err := seedAssetClasses(ctx, client)
	if err != nil {
		return fmt.Errorf("SeedAll(): %v", err)
	}

	err = seedAuthRoles(ctx, client)
	if err != nil {
		return fmt.Errorf("SeedAll(): %v", err)
	}

	err = seedTransactionTypes(ctx, client)
	if err != nil {
		return fmt.Errorf("SeedAll(): %v", err)
	}

	err = seedAuthTypes(ctx, client)
	if err != nil {
		return fmt.Errorf("SeedAll(): %v", err)
	}

	err = seedSourceTypes(ctx, client)
	if err != nil {
		return fmt.Errorf("SeedAll(): %v", err)
	}

	// Seed other entities.
	err = seedStaffAccounts(ctx, client)
	if err != nil {
		return fmt.Errorf("SeedAll(): %v", err)
	}

	return nil
}
