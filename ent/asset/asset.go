// Code generated by ent, DO NOT EDIT.

package asset

import (
	"time"

	"github.com/chenningg/hermitboard-api/pulid"
)

const (
	// Label holds the string label denoting the asset type in the database.
	Label = "asset"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAssetClassID holds the string denoting the asset_class_id field in the database.
	FieldAssetClassID = "asset_class_id"
	// EdgeAssetClass holds the string denoting the asset_class edge name in mutations.
	EdgeAssetClass = "asset_class"
	// EdgeCryptocurrency holds the string denoting the cryptocurrency edge name in mutations.
	EdgeCryptocurrency = "cryptocurrency"
	// EdgeTransactionBases holds the string denoting the transaction_bases edge name in mutations.
	EdgeTransactionBases = "transaction_bases"
	// EdgeTransactionQuotes holds the string denoting the transaction_quotes edge name in mutations.
	EdgeTransactionQuotes = "transaction_quotes"
	// EdgeDailyAssetPrices holds the string denoting the daily_asset_prices edge name in mutations.
	EdgeDailyAssetPrices = "daily_asset_prices"
	// Table holds the table name of the asset in the database.
	Table = "assets"
	// AssetClassTable is the table that holds the asset_class relation/edge.
	AssetClassTable = "assets"
	// AssetClassInverseTable is the table name for the AssetClass entity.
	// It exists in this package in order to avoid circular dependency with the "assetclass" package.
	AssetClassInverseTable = "asset_classes"
	// AssetClassColumn is the table column denoting the asset_class relation/edge.
	AssetClassColumn = "asset_class_id"
	// CryptocurrencyTable is the table that holds the cryptocurrency relation/edge.
	CryptocurrencyTable = "cryptocurrencies"
	// CryptocurrencyInverseTable is the table name for the Cryptocurrency entity.
	// It exists in this package in order to avoid circular dependency with the "cryptocurrency" package.
	CryptocurrencyInverseTable = "cryptocurrencies"
	// CryptocurrencyColumn is the table column denoting the cryptocurrency relation/edge.
	CryptocurrencyColumn = "asset_id"
	// TransactionBasesTable is the table that holds the transaction_bases relation/edge.
	TransactionBasesTable = "transactions"
	// TransactionBasesInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionBasesInverseTable = "transactions"
	// TransactionBasesColumn is the table column denoting the transaction_bases relation/edge.
	TransactionBasesColumn = "base_asset_id"
	// TransactionQuotesTable is the table that holds the transaction_quotes relation/edge.
	TransactionQuotesTable = "transactions"
	// TransactionQuotesInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionQuotesInverseTable = "transactions"
	// TransactionQuotesColumn is the table column denoting the transaction_quotes relation/edge.
	TransactionQuotesColumn = "quote_asset_id"
	// DailyAssetPricesTable is the table that holds the daily_asset_prices relation/edge.
	DailyAssetPricesTable = "daily_asset_prices"
	// DailyAssetPricesInverseTable is the table name for the DailyAssetPrice entity.
	// It exists in this package in order to avoid circular dependency with the "dailyassetprice" package.
	DailyAssetPricesInverseTable = "daily_asset_prices"
	// DailyAssetPricesColumn is the table column denoting the daily_asset_prices relation/edge.
	DailyAssetPricesColumn = "asset_id"
)

// Columns holds all SQL columns for asset fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAssetClassID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// UpdateDefaultDeletedAt holds the default value on update for the "deleted_at" field.
	UpdateDefaultDeletedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.PULID
)
