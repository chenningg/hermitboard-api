package seed

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/sourcetype"
)

func seedSources(ctx context.Context, client *ent.Client) error {
	// Get source type.
	cryptoWalletSourceType, err := client.SourceType.Query().Where(sourcetype.ValueEQ(sourcetype.ValueCryptocurrencyWallet)).Only(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not get source type: %v", err)
	}

	// Create an EVM wallet source.
	_, err = client.Source.Create().
		SetName("EVM Wallet").
		SetIcon("https://upload.wikimedia.org/wikipedia/commons/thumb/6/6f/Ethereum-icon-purple.svg/240px-Ethereum-icon-purple.svg.png").
		SetSourceType(cryptoWalletSourceType).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not seed EVM wallet source: %v", err)
	}

	return nil
}
