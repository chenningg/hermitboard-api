//go:generate go run github.com/dmarkham/enumer -type=TransactionType -json -text -sql

package hbtype

import (
	"fmt"
)

type TransactionType int

const (
	Buy TransactionType = iota
	Sell
	Stake
	DividendIncome
	RentPayment
	RentIncome
	StockDividend
)

func (transactionType *TransactionType) Validate() error {
	if transactionType.IsATransactionType() {
		return nil
	}

	return fmt.Errorf("invalid value for TransactionType, expected %v, got %v", TransactionTypeValues(), transactionType.String())
}
