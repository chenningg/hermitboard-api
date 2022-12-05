package seed

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/sourcetype"
)

func seedSources(ctx context.Context, client *ent.Client) error {
	// Get source types.
	cryptoWalletSourceType, err := client.SourceType.Query().Where(sourcetype.ValueEQ(sourcetype.ValueCryptocurrencyWallet)).Only(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not get source type: %v", err)
	}

	cexSourceType, err := client.SourceType.Query().Where(sourcetype.ValueEQ(sourcetype.ValueExchange)).Only(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not get source type: %v", err)
	}

	dexSourcetype, err := client.SourceType.Query().Where(sourcetype.ValueEQ(sourcetype.ValueDecentralizedExchange)).Only(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not get source type: %v", err)
	}

	bankSourceType, err := client.SourceType.Query().Where(sourcetype.ValueEQ(sourcetype.ValueBank)).Only(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not get source type: %v", err)
	}

	// Create an EVM wallet source.
	_, err = client.Source.Create().
		SetName("Ethereum Wallet").
		SetIcon("https://upload.wikimedia.org/wikipedia/commons/thumb/6/6f/Ethereum-icon-purple.svg/240px-Ethereum-icon-purple.svg.png").
		SetSourceType(cryptoWalletSourceType).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not seed EVM wallet source: %v", err)
	}

	// Create an EVM test wallet Goerli source.
	_, err = client.Source.Create().
		SetName("Goerli Wallet").
		SetIcon("https://upload.wikimedia.org/wikipedia/commons/thumb/6/6f/Ethereum-icon-purple.svg/240px-Ethereum-icon-purple.svg.png").
		SetSourceType(cryptoWalletSourceType).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not seed Goerli wallet source: %v", err)
	}

	// Create a bank wallet source.
	_, err = client.Source.Create().
		SetName("DBS Bank").
		SetIcon("https://logos-download.com/wp-content/uploads/2016/12/DBS_Bank_logo_logotype.png").
		SetSourceType(bankSourceType).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not seed DBS bank source: %v", err)
	}

	_, err = client.Source.Create().
		SetName("Standard Chartered Bank").
		SetIcon("https://assets.stickpng.com/images/5a1d2dc54ac6b00ff574e292.png").
		SetSourceType(bankSourceType).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not seed SCB bank source: %v", err)
	}

	// Create a dex source.
	_, err = client.Source.Create().
		SetName("Uniswap").
		SetIcon("https://upload.wikimedia.org/wikipedia/commons/thumb/e/e7/Uniswap_Logo.svg/1026px-Uniswap_Logo.svg.png").
		SetSourceType(dexSourcetype).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not seed Uniswap DEX source: %v", err)
	}

	// Create a cex source.
	_, err = client.Source.Create().
		SetName("Binance").
		SetIcon("https://upload.wikimedia.org/wikipedia/commons/5/57/Binance_Logo.png").
		SetSourceType(cexSourceType).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not seed Binance CEX source: %v", err)
	}

	_, err = client.Source.Create().
		SetName("Interactive Brokers").
		SetIcon("https://upload.wikimedia.org/wikipedia/commons/thumb/c/ca/Interactive_Brokers_Logo_%282014%29.svg/2560px-Interactive_Brokers_Logo_%282014%29.svg.png").
		SetSourceType(cexSourceType).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedSources(): could not seed IB CEX source: %v", err)
	}

	return nil
}
