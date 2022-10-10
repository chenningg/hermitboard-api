// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "email_confirmed", Type: field.TypeBool, Default: false},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "password_updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "account_auth_type", Type: field.TypeString},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_auth_types_auth_type",
				Columns:    []*schema.Column{AccountsColumns[9]},
				RefColumns: []*schema.Column{AuthTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AssetsColumns holds the columns for the "assets" table.
	AssetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "asset_asset_class", Type: field.TypeString},
	}
	// AssetsTable holds the schema information for the "assets" table.
	AssetsTable = &schema.Table{
		Name:       "assets",
		Columns:    AssetsColumns,
		PrimaryKey: []*schema.Column{AssetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "assets_asset_classes_asset_class",
				Columns:    []*schema.Column{AssetsColumns[4]},
				RefColumns: []*schema.Column{AssetClassesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AssetClassesColumns holds the columns for the "asset_classes" table.
	AssetClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "value", Type: field.TypeEnum, Enums: []string{"CASH_OR_CASH_EQUIVALENT", "COMMODITY", "CRYPTOCURRENCY", "EQUITY", "FIXED_INCOME", "FUTURE", "REAL_ESTATE"}},
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
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "value", Type: field.TypeEnum, Enums: []string{"DEMO", "FREE", "PLUS", "PRO", "ENTERPRISE", "SUPPORT", "ADMIN", "SUPER_ADMIN"}},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// AuthRolesTable holds the schema information for the "auth_roles" table.
	AuthRolesTable = &schema.Table{
		Name:       "auth_roles",
		Columns:    AuthRolesColumns,
		PrimaryKey: []*schema.Column{AuthRolesColumns[0]},
	}
	// AuthTypesColumns holds the columns for the "auth_types" table.
	AuthTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "value", Type: field.TypeEnum, Enums: []string{"LOCAL", "GOOGLE", "APPLE", "FACEBOOK"}},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// AuthTypesTable holds the schema information for the "auth_types" table.
	AuthTypesTable = &schema.Table{
		Name:       "auth_types",
		Columns:    AuthTypesColumns,
		PrimaryKey: []*schema.Column{AuthTypesColumns[0]},
	}
	// BlockchainsColumns holds the columns for the "blockchains" table.
	BlockchainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
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
	// ConnectionsColumns holds the columns for the "connections" table.
	ConnectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "access_token", Type: field.TypeString},
		{Name: "refresh_token", Type: field.TypeString, Nullable: true},
		{Name: "account_id", Type: field.TypeString},
	}
	// ConnectionsTable holds the schema information for the "connections" table.
	ConnectionsTable = &schema.Table{
		Name:       "connections",
		Columns:    ConnectionsColumns,
		PrimaryKey: []*schema.Column{ConnectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "connections_accounts_connections",
				Columns:    []*schema.Column{ConnectionsColumns[7]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "connection_account_id_name",
				Unique:  true,
				Columns: []*schema.Column{ConnectionsColumns[7], ConnectionsColumns[4]},
			},
		},
	}
	// CryptocurrenciesColumns holds the columns for the "cryptocurrencies" table.
	CryptocurrenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "symbol", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "asset_id", Type: field.TypeString, Unique: true},
	}
	// CryptocurrenciesTable holds the schema information for the "cryptocurrencies" table.
	CryptocurrenciesTable = &schema.Table{
		Name:       "cryptocurrencies",
		Columns:    CryptocurrenciesColumns,
		PrimaryKey: []*schema.Column{CryptocurrenciesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cryptocurrencies_assets_cryptocurrency",
				Columns:    []*schema.Column{CryptocurrenciesColumns[7]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DailyAssetPricesColumns holds the columns for the "daily_asset_prices" table.
	DailyAssetPricesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "time", Type: field.TypeTime},
		{Name: "open", Type: field.TypeFloat64, Nullable: true},
		{Name: "high", Type: field.TypeFloat64, Nullable: true},
		{Name: "low", Type: field.TypeFloat64, Nullable: true},
		{Name: "close", Type: field.TypeFloat64, Nullable: true},
		{Name: "adjusted_close", Type: field.TypeFloat64},
		{Name: "asset_id", Type: field.TypeString},
	}
	// DailyAssetPricesTable holds the schema information for the "daily_asset_prices" table.
	DailyAssetPricesTable = &schema.Table{
		Name:       "daily_asset_prices",
		Columns:    DailyAssetPricesColumns,
		PrimaryKey: []*schema.Column{DailyAssetPricesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "daily_asset_prices_assets_daily_asset_prices",
				Columns:    []*schema.Column{DailyAssetPricesColumns[10]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ExchangesColumns holds the columns for the "exchanges" table.
	ExchangesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
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
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_public", Type: field.TypeBool, Default: false},
		{Name: "is_visible", Type: field.TypeBool, Default: true},
		{Name: "account_id", Type: field.TypeString},
	}
	// PortfoliosTable holds the schema information for the "portfolios" table.
	PortfoliosTable = &schema.Table{
		Name:       "portfolios",
		Columns:    PortfoliosColumns,
		PrimaryKey: []*schema.Column{PortfoliosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "portfolios_accounts_portfolios",
				Columns:    []*schema.Column{PortfoliosColumns[7]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "portfolio_account_id_name",
				Unique:  true,
				Columns: []*schema.Column{PortfoliosColumns[7], PortfoliosColumns[4]},
			},
		},
	}
	// SourcesColumns holds the columns for the "sources" table.
	SourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "icon", Type: field.TypeString, Nullable: true},
		{Name: "source_type_sources", Type: field.TypeString},
	}
	// SourcesTable holds the schema information for the "sources" table.
	SourcesTable = &schema.Table{
		Name:       "sources",
		Columns:    SourcesColumns,
		PrimaryKey: []*schema.Column{SourcesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sources_source_types_sources",
				Columns:    []*schema.Column{SourcesColumns[6]},
				RefColumns: []*schema.Column{SourceTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SourceTypesColumns holds the columns for the "source_types" table.
	SourceTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "value", Type: field.TypeEnum, Enums: []string{"CRYPTOCURRENCY_WALLET", "EXCHANGE", "BANK", "DECENTRALIZED_EXCHANGE"}},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// SourceTypesTable holds the schema information for the "source_types" table.
	SourceTypesTable = &schema.Table{
		Name:       "source_types",
		Columns:    SourceTypesColumns,
		PrimaryKey: []*schema.Column{SourceTypesColumns[0]},
	}
	// StaffAccountsColumns holds the columns for the "staff_accounts" table.
	StaffAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "email_confirmed", Type: field.TypeBool, Default: false},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "password_updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "staff_account_auth_type", Type: field.TypeString},
	}
	// StaffAccountsTable holds the schema information for the "staff_accounts" table.
	StaffAccountsTable = &schema.Table{
		Name:       "staff_accounts",
		Columns:    StaffAccountsColumns,
		PrimaryKey: []*schema.Column{StaffAccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "staff_accounts_auth_types_auth_type",
				Columns:    []*schema.Column{StaffAccountsColumns[9]},
				RefColumns: []*schema.Column{AuthTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "time", Type: field.TypeTime},
		{Name: "units", Type: field.TypeInt},
		{Name: "price_per_unit", Type: field.TypeFloat64},
		{Name: "blockchain_id", Type: field.TypeString, Nullable: true},
		{Name: "exchange_id", Type: field.TypeString},
		{Name: "portfolio_id", Type: field.TypeString},
		{Name: "transaction_transaction_type", Type: field.TypeString},
		{Name: "base_asset_id", Type: field.TypeString},
		{Name: "quote_asset_id", Type: field.TypeString, Nullable: true},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_blockchains_transactions",
				Columns:    []*schema.Column{TransactionsColumns[7]},
				RefColumns: []*schema.Column{BlockchainsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "transactions_exchanges_transactions",
				Columns:    []*schema.Column{TransactionsColumns[8]},
				RefColumns: []*schema.Column{ExchangesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "transactions_portfolios_transactions",
				Columns:    []*schema.Column{TransactionsColumns[9]},
				RefColumns: []*schema.Column{PortfoliosColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "transactions_transaction_types_transaction_type",
				Columns:    []*schema.Column{TransactionsColumns[10]},
				RefColumns: []*schema.Column{TransactionTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "transactions_assets_base_asset",
				Columns:    []*schema.Column{TransactionsColumns[11]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "transactions_assets_quote_asset",
				Columns:    []*schema.Column{TransactionsColumns[12]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TransactionTypesColumns holds the columns for the "transaction_types" table.
	TransactionTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "value", Type: field.TypeEnum, Enums: []string{"BUY", "SELL", "SWAP", "STAKE", "DIVIDEND_INCOME", "RENT_PAYMENT", "RENT_INCOME", "STOCK_DIVIDEND"}},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// TransactionTypesTable holds the schema information for the "transaction_types" table.
	TransactionTypesTable = &schema.Table{
		Name:       "transaction_types",
		Columns:    TransactionTypesColumns,
		PrimaryKey: []*schema.Column{TransactionTypesColumns[0]},
	}
	// AccountFriendsColumns holds the columns for the "account_friends" table.
	AccountFriendsColumns = []*schema.Column{
		{Name: "account_id", Type: field.TypeString},
		{Name: "friend_id", Type: field.TypeString},
	}
	// AccountFriendsTable holds the schema information for the "account_friends" table.
	AccountFriendsTable = &schema.Table{
		Name:       "account_friends",
		Columns:    AccountFriendsColumns,
		PrimaryKey: []*schema.Column{AccountFriendsColumns[0], AccountFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_friends_account_id",
				Columns:    []*schema.Column{AccountFriendsColumns[0]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "account_friends_friend_id",
				Columns:    []*schema.Column{AccountFriendsColumns[1]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// AccountAuthRolesColumns holds the columns for the "account_auth_roles" table.
	AccountAuthRolesColumns = []*schema.Column{
		{Name: "account_id", Type: field.TypeString},
		{Name: "auth_role_id", Type: field.TypeString},
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
		{Name: "blockchain_id", Type: field.TypeString},
		{Name: "cryptocurrency_id", Type: field.TypeString},
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
	// PortfolioConnectionsColumns holds the columns for the "portfolio_connections" table.
	PortfolioConnectionsColumns = []*schema.Column{
		{Name: "portfolio_id", Type: field.TypeString},
		{Name: "connection_id", Type: field.TypeString},
	}
	// PortfolioConnectionsTable holds the schema information for the "portfolio_connections" table.
	PortfolioConnectionsTable = &schema.Table{
		Name:       "portfolio_connections",
		Columns:    PortfolioConnectionsColumns,
		PrimaryKey: []*schema.Column{PortfolioConnectionsColumns[0], PortfolioConnectionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "portfolio_connections_portfolio_id",
				Columns:    []*schema.Column{PortfolioConnectionsColumns[0]},
				RefColumns: []*schema.Column{PortfoliosColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "portfolio_connections_connection_id",
				Columns:    []*schema.Column{PortfolioConnectionsColumns[1]},
				RefColumns: []*schema.Column{ConnectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// StaffAccountAuthRolesColumns holds the columns for the "staff_account_auth_roles" table.
	StaffAccountAuthRolesColumns = []*schema.Column{
		{Name: "staff_account_id", Type: field.TypeString},
		{Name: "auth_role_id", Type: field.TypeString},
	}
	// StaffAccountAuthRolesTable holds the schema information for the "staff_account_auth_roles" table.
	StaffAccountAuthRolesTable = &schema.Table{
		Name:       "staff_account_auth_roles",
		Columns:    StaffAccountAuthRolesColumns,
		PrimaryKey: []*schema.Column{StaffAccountAuthRolesColumns[0], StaffAccountAuthRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "staff_account_auth_roles_staff_account_id",
				Columns:    []*schema.Column{StaffAccountAuthRolesColumns[0]},
				RefColumns: []*schema.Column{StaffAccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "staff_account_auth_roles_auth_role_id",
				Columns:    []*schema.Column{StaffAccountAuthRolesColumns[1]},
				RefColumns: []*schema.Column{AuthRolesColumns[0]},
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
		AuthTypesTable,
		BlockchainsTable,
		ConnectionsTable,
		CryptocurrenciesTable,
		DailyAssetPricesTable,
		ExchangesTable,
		PortfoliosTable,
		SourcesTable,
		SourceTypesTable,
		StaffAccountsTable,
		TransactionsTable,
		TransactionTypesTable,
		AccountFriendsTable,
		AccountAuthRolesTable,
		BlockchainCryptocurrenciesTable,
		PortfolioConnectionsTable,
		StaffAccountAuthRolesTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = AuthTypesTable
	AssetsTable.ForeignKeys[0].RefTable = AssetClassesTable
	ConnectionsTable.ForeignKeys[0].RefTable = AccountsTable
	CryptocurrenciesTable.ForeignKeys[0].RefTable = AssetsTable
	DailyAssetPricesTable.ForeignKeys[0].RefTable = AssetsTable
	PortfoliosTable.ForeignKeys[0].RefTable = AccountsTable
	SourcesTable.ForeignKeys[0].RefTable = SourceTypesTable
	StaffAccountsTable.ForeignKeys[0].RefTable = AuthTypesTable
	TransactionsTable.ForeignKeys[0].RefTable = BlockchainsTable
	TransactionsTable.ForeignKeys[1].RefTable = ExchangesTable
	TransactionsTable.ForeignKeys[2].RefTable = PortfoliosTable
	TransactionsTable.ForeignKeys[3].RefTable = TransactionTypesTable
	TransactionsTable.ForeignKeys[4].RefTable = AssetsTable
	TransactionsTable.ForeignKeys[5].RefTable = AssetsTable
	AccountFriendsTable.ForeignKeys[0].RefTable = AccountsTable
	AccountFriendsTable.ForeignKeys[1].RefTable = AccountsTable
	AccountAuthRolesTable.ForeignKeys[0].RefTable = AccountsTable
	AccountAuthRolesTable.ForeignKeys[1].RefTable = AuthRolesTable
	BlockchainCryptocurrenciesTable.ForeignKeys[0].RefTable = BlockchainsTable
	BlockchainCryptocurrenciesTable.ForeignKeys[1].RefTable = CryptocurrenciesTable
	PortfolioConnectionsTable.ForeignKeys[0].RefTable = PortfoliosTable
	PortfolioConnectionsTable.ForeignKeys[1].RefTable = ConnectionsTable
	StaffAccountAuthRolesTable.ForeignKeys[0].RefTable = StaffAccountsTable
	StaffAccountAuthRolesTable.ForeignKeys[1].RefTable = AuthRolesTable
}
