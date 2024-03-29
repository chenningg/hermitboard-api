// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"time"

	"github.com/chenningg/hermitboard-api/pulid"
)

const (
	// Label holds the string label denoting the transaction type in the database.
	Label = "transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldUnits holds the string denoting the units field in the database.
	FieldUnits = "units"
	// FieldPricePerUnit holds the string denoting the price_per_unit field in the database.
	FieldPricePerUnit = "price_per_unit"
	// FieldBlockchainID holds the string denoting the blockchain_id field in the database.
	FieldBlockchainID = "blockchain_id"
	// FieldExchangeID holds the string denoting the exchange_id field in the database.
	FieldExchangeID = "exchange_id"
	// FieldPortfolioID holds the string denoting the portfolio_id field in the database.
	FieldPortfolioID = "portfolio_id"
	// FieldBaseAssetID holds the string denoting the base_asset_id field in the database.
	FieldBaseAssetID = "base_asset_id"
	// FieldQuoteAssetID holds the string denoting the quote_asset_id field in the database.
	FieldQuoteAssetID = "quote_asset_id"
	// EdgeTransactionType holds the string denoting the transaction_type edge name in mutations.
	EdgeTransactionType = "transaction_type"
	// EdgeBaseAsset holds the string denoting the base_asset edge name in mutations.
	EdgeBaseAsset = "base_asset"
	// EdgeQuoteAsset holds the string denoting the quote_asset edge name in mutations.
	EdgeQuoteAsset = "quote_asset"
	// EdgePortfolio holds the string denoting the portfolio edge name in mutations.
	EdgePortfolio = "portfolio"
	// EdgeExchange holds the string denoting the exchange edge name in mutations.
	EdgeExchange = "exchange"
	// EdgeBlockchain holds the string denoting the blockchain edge name in mutations.
	EdgeBlockchain = "blockchain"
	// Table holds the table name of the transaction in the database.
	Table = "transactions"
	// TransactionTypeTable is the table that holds the transaction_type relation/edge.
	TransactionTypeTable = "transactions"
	// TransactionTypeInverseTable is the table name for the TransactionType entity.
	// It exists in this package in order to avoid circular dependency with the "transactiontype" package.
	TransactionTypeInverseTable = "transaction_types"
	// TransactionTypeColumn is the table column denoting the transaction_type relation/edge.
	TransactionTypeColumn = "transaction_transaction_type"
	// BaseAssetTable is the table that holds the base_asset relation/edge.
	BaseAssetTable = "transactions"
	// BaseAssetInverseTable is the table name for the Asset entity.
	// It exists in this package in order to avoid circular dependency with the "asset" package.
	BaseAssetInverseTable = "assets"
	// BaseAssetColumn is the table column denoting the base_asset relation/edge.
	BaseAssetColumn = "base_asset_id"
	// QuoteAssetTable is the table that holds the quote_asset relation/edge.
	QuoteAssetTable = "transactions"
	// QuoteAssetInverseTable is the table name for the Asset entity.
	// It exists in this package in order to avoid circular dependency with the "asset" package.
	QuoteAssetInverseTable = "assets"
	// QuoteAssetColumn is the table column denoting the quote_asset relation/edge.
	QuoteAssetColumn = "quote_asset_id"
	// PortfolioTable is the table that holds the portfolio relation/edge.
	PortfolioTable = "transactions"
	// PortfolioInverseTable is the table name for the Portfolio entity.
	// It exists in this package in order to avoid circular dependency with the "portfolio" package.
	PortfolioInverseTable = "portfolios"
	// PortfolioColumn is the table column denoting the portfolio relation/edge.
	PortfolioColumn = "portfolio_id"
	// ExchangeTable is the table that holds the exchange relation/edge.
	ExchangeTable = "transactions"
	// ExchangeInverseTable is the table name for the Exchange entity.
	// It exists in this package in order to avoid circular dependency with the "exchange" package.
	ExchangeInverseTable = "exchanges"
	// ExchangeColumn is the table column denoting the exchange relation/edge.
	ExchangeColumn = "exchange_id"
	// BlockchainTable is the table that holds the blockchain relation/edge.
	BlockchainTable = "transactions"
	// BlockchainInverseTable is the table name for the Blockchain entity.
	// It exists in this package in order to avoid circular dependency with the "blockchain" package.
	BlockchainInverseTable = "blockchains"
	// BlockchainColumn is the table column denoting the blockchain relation/edge.
	BlockchainColumn = "blockchain_id"
)

// Columns holds all SQL columns for transaction fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTime,
	FieldUnits,
	FieldPricePerUnit,
	FieldBlockchainID,
	FieldExchangeID,
	FieldPortfolioID,
	FieldBaseAssetID,
	FieldQuoteAssetID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "transactions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"transaction_transaction_type",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
