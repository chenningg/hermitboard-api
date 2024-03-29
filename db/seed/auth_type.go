package seed

import (
	"context"
	"fmt"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/authtype"
)

func seedAuthTypes(ctx context.Context, client *ent.Client) error {
	seedAuthTypes := []authtype.Value{
		authtype.ValueLocal, authtype.ValueGoogle, authtype.ValueFacebook, authtype.ValueApple,
	}
	seedDescriptions := []string{
		"Local authentication with passwords stored in the database.", "OAuth authentication through Google.",
		"OAuth authentication through Facebook.", "OAuth authentication through Apple.",
	}

	bulk := make([]*ent.AuthTypeCreate, len(seedAuthTypes))
	for i := 0; i < len(seedAuthTypes); i++ {
		bulk[i] = client.AuthType.Create().SetValue(seedAuthTypes[i]).SetDescription(seedDescriptions[i])
	}

	// Create in bulk
	_, err := client.AuthType.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedAuthTypes(): %v", err)
	}

	return nil
}
