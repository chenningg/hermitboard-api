package seed

import (
	"context"
	"fmt"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/assetclass"
)

func seedAssetClasses(ctx context.Context, client *ent.Client) error {
	seedAssetClasses := []assetclass.Value{
		assetclass.ValueCashOrCashEquivalent, assetclass.ValueCommodity, assetclass.ValueCryptocurrency,
		assetclass.ValueEquity, assetclass.ValueFixedIncome, assetclass.ValueFuture,
		assetclass.ValueRealEstate,
	}
	seedDescriptions := []string{
		"Value of assets that are cash or can be converted to cash immediately.",
		"A raw material or primary agricultural product that can be bought and sold.",
		"A digital or virtual currency secured by cryptography.",
		"The ownership of assets with financial value that may have debts or other liabilities attached to them.",
		"Assets and securities that pay out a set level of cash flows to investors.",
		"Derivative financial contracts that obligate parties to buy or sell an asset at a predetermined future date and price.",
		"Property consisting of land and buildings.",
	}

	bulk := make([]*ent.AssetClassCreate, len(seedAssetClasses))
	for i := 0; i < len(seedAssetClasses); i++ {
		bulk[i] = client.AssetClass.Create().SetValue(seedAssetClasses[i]).SetDescription(seedDescriptions[i])
	}

	// Create in bulk
	_, err := client.AssetClass.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedAssetClasses(): %v", err)
	}

	return nil
}
