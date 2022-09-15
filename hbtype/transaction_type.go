//go:generate go run github.com/dmarkham/enumer -type=TransactionType -json -text -sql

package hbtype

import (
	"fmt"
	"io"
	"strconv"
)

type TransactionType int

const (
	BUY TransactionType = iota + 1
	SELL
	STAKE
	DIVIDEND_INCOME
	RENT_PAYMENT
	RENT_INCOME
	STOCK_DIVIDEND
)

func (transactionType *TransactionType) Validate() error {
	if transactionType.IsATransactionType() {
		return nil
	}

	return fmt.Errorf("invalid value for TransactionType, expected %v, got %v", TransactionTypeValues(), transactionType.String())
}

func (transactionType TransactionType) Values() []string {
	return TransactionTypeStrings()
}

// Marshals a TransactionType into a graphql scalar string.
func (transactionType TransactionType) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(transactionType.String()))
}

// Unmarshals a graphql scalar into a TransactionType Go type.
func (transactionType *TransactionType) UnmarshalGQL(v interface{}) error {
	return transactionType.Scan(v)
}
