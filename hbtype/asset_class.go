//go:generate go run github.com/dmarkham/enumer -type=AssetClass -json -text -sql

package hbtype

import (
	"fmt"
	"io"
	"strconv"
)

type AssetClass int

const (
	CASH_OR_CASH_EQUIVALENT AssetClass = iota + 1
	COMMODITY
	CRYPTOCURRENCY
	EQUITY
	FIXED_INCOME
	FUTURE
	REAL_ESTATE
)

func (assetClass *AssetClass) Validate() error {
	if assetClass.IsAAssetClass() {
		return nil
	}

	return fmt.Errorf("invalid value for AssetClass, expected %v, got %v", AssetClassValues(), assetClass.String())
}

func (assetClass AssetClass) Values() []string {
	return AssetClassStrings()
}

// Marshals an AssetClass into a graphql scalar string.
func (assetClass AssetClass) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(assetClass.String()))
}

// Unmarshals a graphql scalar into an AssetClass Go type.
func (assetClass *AssetClass) UnmarshalGQL(v interface{}) error {
	return assetClass.Scan(v)
}
