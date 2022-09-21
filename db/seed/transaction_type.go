package seed

import (
	"context"
	"fmt"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/transactiontype"
)

func seedTransactionTypes(ctx context.Context, client *ent.Client) error {
	seedTransactionTypes := []transactiontype.TransactionType{
		transactiontype.TransactionTypeBuy, transactiontype.TransactionTypeSell, transactiontype.TransactionTypeStake,
		transactiontype.TransactionTypeDividendIncome, transactiontype.TransactionTypeRentPayment,
		transactiontype.TransactionTypeRentIncome, transactiontype.TransactionTypeStockDividend,
	}
	seedDescriptions := []string{
		"A buy transaction.",
		"A sell transaction.",
		"Staking of a cryptocurrency.",
		"Income from dividends.",
		"Outflow of cash due to rent payment.",
		"Inflow of cash due to rent collection.",
		"Dividends paid out in stock.",
	}

	bulk := make([]*ent.TransactionTypeCreate, len(seedTransactionTypes))
	for i := 0; i < len(seedTransactionTypes); i++ {
		bulk[i] = client.TransactionType.Create().SetTransactionType(seedTransactionTypes[i]).SetDescription(seedDescriptions[i])
	}

	// Create in bulk
	_, err := client.TransactionType.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedTransactionTypes(): %v", err)
	}

	return nil
}
