// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"

	hbtype "github.com/chenningg/hermitboard-api/hbtype"
)

type Account struct {
	ID       hbtype.ULID
	AuthID   string
	Nickname string
	Email    string
}

type Asset struct {
	ID hbtype.ULID
}

type AssetClass struct {
	ID          string
	Description sql.NullString
}

type Blockchain struct {
	ID                          hbtype.ULID
	Name                        string
	Symbol                      string
	NativeTokenCryptocurrencyID hbtype.NullULID
	ChainID                     sql.NullInt64
}

type Cryptocurrency struct {
	ID           hbtype.ULID
	AssetID      hbtype.ULID
	AssetClassID hbtype.ULID
	Symbol       string
	ShortName    string
	LongName     string
}

type Portfolio struct {
	ID        hbtype.ULID
	Name      string
	AccountID hbtype.ULID
	IsPublic  bool
	IsVisible bool
}
