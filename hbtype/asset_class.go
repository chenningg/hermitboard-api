//go:generate go run github.com/dmarkham/enumer -type=AssetClass -json -text -sql

package hbtype

import (
	"fmt"
)

type AssetClass int

const (
	CashOrCashEquivalent AssetClass = iota
	Commodity
	Cryptocurrency
	Equity
	FixedIncome
	Future
	RealEstate
)

func (assetClass *AssetClass) Validate() error {
	if assetClass.IsAAssetClass() {
		return nil
	}

	return fmt.Errorf("invalid value for AssetClass, expected %v, got %v", AssetClassValues(), assetClass.String())
}
