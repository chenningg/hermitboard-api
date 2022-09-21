// Code generated by ent, DO NOT EDIT.

package blockchaincryptocurrency

import (
	"time"

	"github.com/chenningg/hermitboard-api/pulid"
)

const (
	// Label holds the string label denoting the blockchaincryptocurrency type in the database.
	Label = "blockchain_cryptocurrency"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldBlockchainID holds the string denoting the blockchain_id field in the database.
	FieldBlockchainID = "blockchain_id"
	// FieldCryptocurrencyID holds the string denoting the cryptocurrency_id field in the database.
	FieldCryptocurrencyID = "cryptocurrency_id"
	// EdgeBlockchain holds the string denoting the blockchain edge name in mutations.
	EdgeBlockchain = "blockchain"
	// EdgeCryptocurrency holds the string denoting the cryptocurrency edge name in mutations.
	EdgeCryptocurrency = "cryptocurrency"
	// Table holds the table name of the blockchaincryptocurrency in the database.
	Table = "blockchain_cryptocurrencies"
	// BlockchainTable is the table that holds the blockchain relation/edge.
	BlockchainTable = "blockchain_cryptocurrencies"
	// BlockchainInverseTable is the table name for the Blockchain entity.
	// It exists in this package in order to avoid circular dependency with the "blockchain" package.
	BlockchainInverseTable = "blockchains"
	// BlockchainColumn is the table column denoting the blockchain relation/edge.
	BlockchainColumn = "blockchain_id"
	// CryptocurrencyTable is the table that holds the cryptocurrency relation/edge.
	CryptocurrencyTable = "blockchain_cryptocurrencies"
	// CryptocurrencyInverseTable is the table name for the Cryptocurrency entity.
	// It exists in this package in order to avoid circular dependency with the "cryptocurrency" package.
	CryptocurrencyInverseTable = "cryptocurrencies"
	// CryptocurrencyColumn is the table column denoting the cryptocurrency relation/edge.
	CryptocurrencyColumn = "cryptocurrency_id"
)

// Columns holds all SQL columns for blockchaincryptocurrency fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldBlockchainID,
	FieldCryptocurrencyID,
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