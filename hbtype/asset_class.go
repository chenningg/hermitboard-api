//go:generate go run github.com/dmarkham/enumer -type=AssetClass

package hbtype

import (
	"database/sql/driver"
	"fmt"
)

type AssetClass int

const (
	Commodity AssetClass = iota
	Cryptocurrency
	Equity
	FixedIncome
	RealEstate
)

func (assetClass *AssetClass) Validate() error {
	if assetClass.IsAAssetClass() {
		return nil
	}

	return fmt.Errorf("invalid value for AssetClass, expected %v, got %v", AssetClassValues(), assetClass.String())
}

// Encodes a string text to AssetClass type.
func (assetClass *AssetClass) UnmarshalText(text []byte) error {
	input := string(text)
	marshalledAppEnv, err := AssetClassString(input)
	// If unable to parse, use a default value.
	if err != nil {
		return fmt.Errorf("invalid value for AppEnv, expected one of %v, got '%v'", AssetClassValues(), input)
	}

	// Otherwise set AppEnv to be the supplied value.
	*assetClass = marshalledAppEnv

	return nil
}

// Implementation of Valuer to scan Go type to database value.
func (assetClass *AssetClass) Value() (driver.Value, error) {
	return assetClass.String(), nil
}

// Implementation of Scanner to scan database value to Go type.
func (assetClass *AssetClass) Scan(src any) error {
	var input string

	switch srcType := src.(type) {
	case string:
		input = src.(string)
	case []byte:
		input = string(src.([]byte))
	default:
		return fmt.Errorf("failed to scan AssetClass due to incompatible type %T", srcType)
	}

	marshalledAssetClass, err := AssetClassString(input)

	if err != nil {
		return fmt.Errorf("error converting database value to AssetClass")
	}

	*assetClass = marshalledAssetClass
	return nil
}
