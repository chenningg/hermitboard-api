// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "auth_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "password_updated_at", Type: field.TypeTime},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// AssetsColumns holds the columns for the "assets" table.
	AssetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "asset_asset_class", Type: field.TypeUUID},
	}
	// AssetsTable holds the schema information for the "assets" table.
	AssetsTable = &schema.Table{
		Name:       "assets",
		Columns:    AssetsColumns,
		PrimaryKey: []*schema.Column{AssetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "assets_asset_classes_asset_class",
				Columns:    []*schema.Column{AssetsColumns[3]},
				RefColumns: []*schema.Column{AssetClassesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AssetClassesColumns holds the columns for the "asset_classes" table.
	AssetClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "class", Type: field.TypeEnum, Enums: []string{"CashOrCashEquivalent", "Commodity", "Cryptocurrency", "Equity", "FixedIncome", "Future", "RealEstate"}},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// AssetClassesTable holds the schema information for the "asset_classes" table.
	AssetClassesTable = &schema.Table{
		Name:       "asset_classes",
		Columns:    AssetClassesColumns,
		PrimaryKey: []*schema.Column{AssetClassesColumns[0]},
	}
	// AuthRolesColumns holds the columns for the "auth_roles" table.
	AuthRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"Demo", "Free", "Plus", "Pro", "Enterprise", "Support", "Admin", "SuperAdmin"}},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// AuthRolesTable holds the schema information for the "auth_roles" table.
	AuthRolesTable = &schema.Table{
		Name:       "auth_roles",
		Columns:    AuthRolesColumns,
		PrimaryKey: []*schema.Column{AuthRolesColumns[0]},
	}
	// BlockchainsColumns holds the columns for the "blockchains" table.
	BlockchainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "symbol", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString, Nullable: true},
		{Name: "chain_id", Type: field.TypeInt64, Nullable: true},
	}
	// BlockchainsTable holds the schema information for the "blockchains" table.
	BlockchainsTable = &schema.Table{
		Name:       "blockchains",
		Columns:    BlockchainsColumns,
		PrimaryKey: []*schema.Column{BlockchainsColumns[0]},
	}
	// CryptocurrenciesColumns holds the columns for the "cryptocurrencies" table.
	CryptocurrenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "symbol", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "asset_cryptocurrency", Type: field.TypeUUID, Unique: true},
	}
	// CryptocurrenciesTable holds the schema information for the "cryptocurrencies" table.
	CryptocurrenciesTable = &schema.Table{
		Name:       "cryptocurrencies",
		Columns:    CryptocurrenciesColumns,
		PrimaryKey: []*schema.Column{CryptocurrenciesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cryptocurrencies_assets_cryptocurrency",
				Columns:    []*schema.Column{CryptocurrenciesColumns[6]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DailyAssetPricesColumns holds the columns for the "daily_asset_prices" table.
	DailyAssetPricesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "time", Type: field.TypeTime},
		{Name: "open", Type: field.TypeFloat64, Nullable: true},
		{Name: "high", Type: field.TypeFloat64, Nullable: true},
		{Name: "low", Type: field.TypeFloat64, Nullable: true},
		{Name: "close", Type: field.TypeFloat64, Nullable: true},
		{Name: "adjusted_close", Type: field.TypeFloat64},
		{Name: "asset_daily_asset_price", Type: field.TypeUUID},
	}
	// DailyAssetPricesTable holds the schema information for the "daily_asset_prices" table.
	DailyAssetPricesTable = &schema.Table{
		Name:       "daily_asset_prices",
		Columns:    DailyAssetPricesColumns,
		PrimaryKey: []*schema.Column{DailyAssetPricesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "daily_asset_prices_assets_daily_asset_price",
				Columns:    []*schema.Column{DailyAssetPricesColumns[9]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "dailyassetprice_time_asset_daily_asset_price",
				Unique:  true,
				Columns: []*schema.Column{DailyAssetPricesColumns[3], DailyAssetPricesColumns[9]},
			},
		},
	}
	// ExchangesColumns holds the columns for the "exchanges" table.
	ExchangesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
	}
	// ExchangesTable holds the schema information for the "exchanges" table.
	ExchangesTable = &schema.Table{
		Name:       "exchanges",
		Columns:    ExchangesColumns,
		PrimaryKey: []*schema.Column{ExchangesColumns[0]},
	}
	// PortfoliosColumns holds the columns for the "portfolios" table.
	PortfoliosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "is_public", Type: field.TypeBool, Default: false},
		{Name: "is_visible", Type: field.TypeBool, Default: true},
		{Name: "account_portfolios", Type: field.TypeUUID},
	}
	// PortfoliosTable holds the schema information for the "portfolios" table.
	PortfoliosTable = &schema.Table{
		Name:       "portfolios",
		Columns:    PortfoliosColumns,
		PrimaryKey: []*schema.Column{PortfoliosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "portfolios_accounts_portfolios",
				Columns:    []*schema.Column{PortfoliosColumns[6]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "portfolio_name_account_portfolios",
				Unique:  true,
				Columns: []*schema.Column{PortfoliosColumns[3], PortfoliosColumns[6]},
			},
		},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "time", Type: field.TypeTime},
		{Name: "units", Type: field.TypeInt},
		{Name: "price_per_unit", Type: field.TypeFloat64},
		{Name: "portfolio_transactions", Type: field.TypeUUID},
		{Name: "transaction_transaction_type", Type: field.TypeUUID},
		{Name: "transaction_exchange", Type: field.TypeUUID, Nullable: true},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_portfolios_transactions",
				Columns:    []*schema.Column{TransactionsColumns[6]},
				RefColumns: []*schema.Column{PortfoliosColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "transactions_transaction_types_transaction_type",
				Columns:    []*schema.Column{TransactionsColumns[7]},
				RefColumns: []*schema.Column{TransactionTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "transactions_exchanges_exchange",
				Columns:    []*schema.Column{TransactionsColumns[8]},
				RefColumns: []*schema.Column{ExchangesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "transaction_time_portfolio_transactions",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[3], TransactionsColumns[6]},
			},
		},
	}
	// TransactionTypesColumns holds the columns for the "transaction_types" table.
	TransactionTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"Buy", "Sell", "Stake", "DividendIncome", "RentPayment", "RentIncome", "StockDividend"}},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// TransactionTypesTable holds the schema information for the "transaction_types" table.
	TransactionTypesTable = &schema.Table{
		Name:       "transaction_types",
		Columns:    TransactionTypesColumns,
		PrimaryKey: []*schema.Column{TransactionTypesColumns[0]},
	}
	// AccountAuthRolesColumns holds the columns for the "account_auth_roles" table.
	AccountAuthRolesColumns = []*schema.Column{
		{Name: "account_id", Type: field.TypeUUID},
		{Name: "auth_role_id", Type: field.TypeUUID},
	}
	// AccountAuthRolesTable holds the schema information for the "account_auth_roles" table.
	AccountAuthRolesTable = &schema.Table{
		Name:       "account_auth_roles",
		Columns:    AccountAuthRolesColumns,
		PrimaryKey: []*schema.Column{AccountAuthRolesColumns[0], AccountAuthRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_auth_roles_account_id",
				Columns:    []*schema.Column{AccountAuthRolesColumns[0]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "account_auth_roles_auth_role_id",
				Columns:    []*schema.Column{AccountAuthRolesColumns[1]},
				RefColumns: []*schema.Column{AuthRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// BlockchainCryptocurrenciesColumns holds the columns for the "blockchain_cryptocurrencies" table.
	BlockchainCryptocurrenciesColumns = []*schema.Column{
		{Name: "blockchain_id", Type: field.TypeUUID},
		{Name: "cryptocurrency_id", Type: field.TypeUUID},
	}
	// BlockchainCryptocurrenciesTable holds the schema information for the "blockchain_cryptocurrencies" table.
	BlockchainCryptocurrenciesTable = &schema.Table{
		Name:       "blockchain_cryptocurrencies",
		Columns:    BlockchainCryptocurrenciesColumns,
		PrimaryKey: []*schema.Column{BlockchainCryptocurrenciesColumns[0], BlockchainCryptocurrenciesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blockchain_cryptocurrencies_blockchain_id",
				Columns:    []*schema.Column{BlockchainCryptocurrenciesColumns[0]},
				RefColumns: []*schema.Column{BlockchainsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "blockchain_cryptocurrencies_cryptocurrency_id",
				Columns:    []*schema.Column{BlockchainCryptocurrenciesColumns[1]},
				RefColumns: []*schema.Column{CryptocurrenciesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TransactionBaseAssetColumns holds the columns for the "transaction_base_asset" table.
	TransactionBaseAssetColumns = []*schema.Column{
		{Name: "transaction_id", Type: field.TypeUUID},
		{Name: "asset_id", Type: field.TypeUUID},
	}
	// TransactionBaseAssetTable holds the schema information for the "transaction_base_asset" table.
	TransactionBaseAssetTable = &schema.Table{
		Name:       "transaction_base_asset",
		Columns:    TransactionBaseAssetColumns,
		PrimaryKey: []*schema.Column{TransactionBaseAssetColumns[0], TransactionBaseAssetColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transaction_base_asset_transaction_id",
				Columns:    []*schema.Column{TransactionBaseAssetColumns[0]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "transaction_base_asset_asset_id",
				Columns:    []*schema.Column{TransactionBaseAssetColumns[1]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TransactionQuoteAssetColumns holds the columns for the "transaction_quote_asset" table.
	TransactionQuoteAssetColumns = []*schema.Column{
		{Name: "transaction_id", Type: field.TypeUUID},
		{Name: "asset_id", Type: field.TypeUUID},
	}
	// TransactionQuoteAssetTable holds the schema information for the "transaction_quote_asset" table.
	TransactionQuoteAssetTable = &schema.Table{
		Name:       "transaction_quote_asset",
		Columns:    TransactionQuoteAssetColumns,
		PrimaryKey: []*schema.Column{TransactionQuoteAssetColumns[0], TransactionQuoteAssetColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transaction_quote_asset_transaction_id",
				Columns:    []*schema.Column{TransactionQuoteAssetColumns[0]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "transaction_quote_asset_asset_id",
				Columns:    []*schema.Column{TransactionQuoteAssetColumns[1]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AssetsTable,
		AssetClassesTable,
		AuthRolesTable,
		BlockchainsTable,
		CryptocurrenciesTable,
		DailyAssetPricesTable,
		ExchangesTable,
		PortfoliosTable,
		TransactionsTable,
		TransactionTypesTable,
		AccountAuthRolesTable,
		BlockchainCryptocurrenciesTable,
		TransactionBaseAssetTable,
		TransactionQuoteAssetTable,
	}
)

func init() {
	AccountsTable.Annotation = &entsql.Annotation{}
	AccountsTable.Annotation.Checks = map[string]string{
		"account_chk_auth_id_password_not_null": "auth_id IS NOT NULL OR password IS NOT NULL",
	}
	AssetsTable.ForeignKeys[0].RefTable = AssetClassesTable
	CryptocurrenciesTable.ForeignKeys[0].RefTable = AssetsTable
	DailyAssetPricesTable.ForeignKeys[0].RefTable = AssetsTable
	PortfoliosTable.ForeignKeys[0].RefTable = AccountsTable
	TransactionsTable.ForeignKeys[0].RefTable = PortfoliosTable
	TransactionsTable.ForeignKeys[1].RefTable = TransactionTypesTable
	TransactionsTable.ForeignKeys[2].RefTable = ExchangesTable
	AccountAuthRolesTable.ForeignKeys[0].RefTable = AccountsTable
	AccountAuthRolesTable.ForeignKeys[1].RefTable = AuthRolesTable
	BlockchainCryptocurrenciesTable.ForeignKeys[0].RefTable = BlockchainsTable
	BlockchainCryptocurrenciesTable.ForeignKeys[1].RefTable = CryptocurrenciesTable
	TransactionBaseAssetTable.ForeignKeys[0].RefTable = TransactionsTable
	TransactionBaseAssetTable.ForeignKeys[1].RefTable = AssetsTable
	TransactionQuoteAssetTable.ForeignKeys[0].RefTable = TransactionsTable
	TransactionQuoteAssetTable.ForeignKeys[1].RefTable = AssetsTable
}
