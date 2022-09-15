// Code generated by ent, DO NOT EDIT.

package dailyassetprice

import (
	"time"

	"github.com/chenningg/hermitboard-api/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the dailyassetprice type in the database.
	Label = "daily_asset_price"
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
	// FieldOpen holds the string denoting the open field in the database.
	FieldOpen = "open"
	// FieldHigh holds the string denoting the high field in the database.
	FieldHigh = "high"
	// FieldLow holds the string denoting the low field in the database.
	FieldLow = "low"
	// FieldClose holds the string denoting the close field in the database.
	FieldClose = "close"
	// FieldAdjustedClose holds the string denoting the adjusted_close field in the database.
	FieldAdjustedClose = "adjusted_close"
	// FieldBaseAssetID holds the string denoting the base_asset_id field in the database.
	FieldBaseAssetID = "base_asset_id"
	// EdgeBaseAsset holds the string denoting the base_asset edge name in mutations.
	EdgeBaseAsset = "base_asset"
	// Table holds the table name of the dailyassetprice in the database.
	Table = "daily_asset_prices"
	// BaseAssetTable is the table that holds the base_asset relation/edge.
	BaseAssetTable = "daily_asset_prices"
	// BaseAssetInverseTable is the table name for the Asset entity.
	// It exists in this package in order to avoid circular dependency with the "asset" package.
	BaseAssetInverseTable = "assets"
	// BaseAssetColumn is the table column denoting the base_asset relation/edge.
	BaseAssetColumn = "base_asset_id"
)

// Columns holds all SQL columns for dailyassetprice fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTime,
	FieldOpen,
	FieldHigh,
	FieldLow,
	FieldClose,
	FieldAdjustedClose,
	FieldBaseAssetID,
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
	// DefaultTime holds the default value on creation for the "time" field.
	DefaultTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.PULID
)
