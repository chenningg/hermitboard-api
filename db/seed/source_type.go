package seed

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/sourcetype"
)

func seedSourceTypes(ctx context.Context, client *ent.Client) error {
	seedSourceTypes := []sourcetype.Value{
		sourcetype.ValueExchange, sourcetype.ValueDecentralizedExchange, sourcetype.ValueBank,
	}
	seedDescriptions := []string{
		"Centralized exchanges with various investment assets.",
		"Decentralized exchanges that primarily offer cryptocurrencies.",
		"A bank that holds cash and possibly investments.",
	}

	bulk := make([]*ent.SourceTypeCreate, len(seedSourceTypes))
	for i := 0; i < len(seedSourceTypes); i++ {
		bulk[i] = client.SourceType.Create().SetValue(seedSourceTypes[i]).SetDescription(seedDescriptions[i])
	}

	// Create in bulk
	_, err := client.SourceType.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedSourceTypes(): %v", err)
	}

	return nil
}
