-- create "exchanges" table
CREATE TABLE "exchanges"
(
    "id"         character varying NOT NULL,
    "created_at" timestamptz       NOT NULL,
    "updated_at" timestamptz       NOT NULL,
    "deleted_at" timestamptz       NULL,
    "name"       character varying NOT NULL,
    "icon"       character varying NULL,
    "url"        character varying NULL,
    PRIMARY KEY ("id")
);
-- create "accounts" table
CREATE TABLE "accounts"
(
    "id"                  character varying NOT NULL,
    "created_at"          timestamptz       NOT NULL,
    "updated_at"          timestamptz       NOT NULL,
    "deleted_at"          timestamptz       NULL,
    "auth_type"           character varying NOT NULL DEFAULT 'LOCAL',
    "nickname"            character varying NOT NULL,
    "email"               character varying NOT NULL,
    "password"            character varying NULL,
    "password_updated_at" timestamptz       NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "account_chk_if_auth_type_local_then_password_not_null" CHECK ((auth_type <> 'LOCAL') OR (password IS NOT NULL))
);
-- create index "accounts_nickname_key" to table: "accounts"
CREATE UNIQUE INDEX "accounts_nickname_key" ON "accounts" ("nickname");
-- create index "accounts_email_key" to table: "accounts"
CREATE UNIQUE INDEX "accounts_email_key" ON "accounts" ("email");
-- create "portfolios" table
CREATE TABLE "portfolios"
(
    "id"         character varying NOT NULL,
    "created_at" timestamptz       NOT NULL,
    "updated_at" timestamptz       NOT NULL,
    "deleted_at" timestamptz       NULL,
    "name"       character varying NOT NULL,
    "is_public"  boolean           NOT NULL DEFAULT false,
    "is_visible" boolean           NOT NULL DEFAULT true,
    "account_id" character varying NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "portfolios_accounts_portfolios" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE NO ACTION
);
-- create index "portfolio_account_id_name" to table: "portfolios"
CREATE UNIQUE INDEX "portfolio_account_id_name" ON "portfolios" ("account_id", "name");
-- create "transaction_types" table
CREATE TABLE "transaction_types"
(
    "id"               character varying NOT NULL,
    "created_at"       timestamptz       NOT NULL,
    "updated_at"       timestamptz       NOT NULL,
    "deleted_at"       timestamptz       NULL,
    "transaction_type" character varying NOT NULL,
    "description"      character varying NULL,
    PRIMARY KEY ("id")
);
-- create "asset_classes" table
CREATE TABLE "asset_classes"
(
    "id"          character varying NOT NULL,
    "created_at"  timestamptz       NOT NULL,
    "updated_at"  timestamptz       NOT NULL,
    "deleted_at"  timestamptz       NULL,
    "asset_class" character varying NOT NULL,
    "description" character varying NULL,
    PRIMARY KEY ("id")
);
-- create "assets" table
CREATE TABLE "assets"
(
    "id"             character varying NOT NULL,
    "created_at"     timestamptz       NOT NULL,
    "updated_at"     timestamptz       NOT NULL,
    "deleted_at"     timestamptz       NULL,
    "asset_class_id" character varying NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "assets_asset_classes_asset_class" FOREIGN KEY ("asset_class_id") REFERENCES "asset_classes" ("id") ON DELETE NO ACTION
);
-- create "transactions" table
CREATE TABLE "transactions"
(
    "id"                  character varying NOT NULL,
    "created_at"          timestamptz       NOT NULL,
    "updated_at"          timestamptz       NOT NULL,
    "deleted_at"          timestamptz       NULL,
    "time"                timestamptz       NOT NULL,
    "units"               bigint            NOT NULL,
    "price_per_unit"      double precision  NOT NULL,
    "exchange_id"         character varying NOT NULL,
    "portfolio_id"        character varying NOT NULL,
    "transaction_type_id" character varying NOT NULL,
    "base_asset_id"       character varying NOT NULL,
    "quote_asset_id"      character varying NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "transactions_exchanges_transactions" FOREIGN KEY ("exchange_id") REFERENCES "exchanges" ("id") ON DELETE NO ACTION,
    CONSTRAINT "transactions_portfolios_transactions" FOREIGN KEY ("portfolio_id") REFERENCES "portfolios" ("id") ON DELETE NO ACTION,
    CONSTRAINT "transactions_transaction_types_transaction_type" FOREIGN KEY ("transaction_type_id") REFERENCES "transaction_types" ("id") ON DELETE NO ACTION,
    CONSTRAINT "transactions_assets_base_asset" FOREIGN KEY ("base_asset_id") REFERENCES "assets" ("id") ON DELETE NO ACTION,
    CONSTRAINT "transactions_assets_quote_asset" FOREIGN KEY ("quote_asset_id") REFERENCES "assets" ("id") ON DELETE SET NULL
);
-- create "auth_roles" table
CREATE TABLE "auth_roles"
(
    "id"          character varying NOT NULL,
    "created_at"  timestamptz       NOT NULL,
    "updated_at"  timestamptz       NOT NULL,
    "deleted_at"  timestamptz       NULL,
    "auth_role"   character varying NOT NULL,
    "description" character varying NULL,
    PRIMARY KEY ("id")
);
-- create "account_auth_roles" table
CREATE TABLE "account_auth_roles"
(
    "id"           character varying NOT NULL,
    "created_at"   timestamptz       NOT NULL,
    "updated_at"   timestamptz       NOT NULL,
    "deleted_at"   timestamptz       NULL,
    "account_id"   character varying NOT NULL,
    "auth_role_id" character varying NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "account_auth_roles_accounts_account" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE NO ACTION,
    CONSTRAINT "account_auth_roles_auth_roles_auth_role" FOREIGN KEY ("auth_role_id") REFERENCES "auth_roles" ("id") ON DELETE NO ACTION
);
-- create index "accountauthrole_account_id_auth_role_id" to table: "account_auth_roles"
CREATE UNIQUE INDEX "accountauthrole_account_id_auth_role_id" ON "account_auth_roles" ("account_id", "auth_role_id");
-- create "blockchains" table
CREATE TABLE "blockchains"
(
    "id"         character varying NOT NULL,
    "created_at" timestamptz       NOT NULL,
    "updated_at" timestamptz       NOT NULL,
    "deleted_at" timestamptz       NULL,
    "name"       character varying NOT NULL,
    "symbol"     character varying NOT NULL,
    "icon"       character varying NULL,
    "chain_id"   bigint            NULL,
    PRIMARY KEY ("id")
);
-- create "cryptocurrencies" table
CREATE TABLE "cryptocurrencies"
(
    "id"         character varying NOT NULL,
    "created_at" timestamptz       NOT NULL,
    "updated_at" timestamptz       NOT NULL,
    "deleted_at" timestamptz       NULL,
    "symbol"     character varying NOT NULL,
    "icon"       character varying NULL,
    "name"       character varying NOT NULL,
    "asset_id"   character varying NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "cryptocurrencies_assets_cryptocurrency" FOREIGN KEY ("asset_id") REFERENCES "assets" ("id") ON DELETE NO ACTION
);
-- create index "cryptocurrencies_asset_id_key" to table: "cryptocurrencies"
CREATE UNIQUE INDEX "cryptocurrencies_asset_id_key" ON "cryptocurrencies" ("asset_id");
-- create "blockchain_cryptocurrencies" table
CREATE TABLE "blockchain_cryptocurrencies"
(
    "id"                character varying NOT NULL,
    "created_at"        timestamptz       NOT NULL,
    "updated_at"        timestamptz       NOT NULL,
    "deleted_at"        timestamptz       NULL,
    "blockchain_id"     character varying NOT NULL,
    "cryptocurrency_id" character varying NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "blockchain_cryptocurrencies_blockchains_blockchain" FOREIGN KEY ("blockchain_id") REFERENCES "blockchains" ("id") ON DELETE NO ACTION,
    CONSTRAINT "blockchain_cryptocurrencies_cryptocurrencies_cryptocurrency" FOREIGN KEY ("cryptocurrency_id") REFERENCES "cryptocurrencies" ("id") ON DELETE NO ACTION
);
-- create index "blockchaincryptocurrency_blockchain_id_cryptocurrency_id" to table: "blockchain_cryptocurrencies"
CREATE UNIQUE INDEX "blockchaincryptocurrency_blockchain_id_cryptocurrency_id" ON "blockchain_cryptocurrencies" ("blockchain_id", "cryptocurrency_id");
-- create "daily_asset_prices" table
CREATE TABLE "daily_asset_prices"
(
    "id"             character varying NOT NULL,
    "created_at"     timestamptz       NOT NULL,
    "updated_at"     timestamptz       NOT NULL,
    "deleted_at"     timestamptz       NULL,
    "time"           timestamptz       NOT NULL,
    "open"           double precision  NULL,
    "high"           double precision  NULL,
    "low"            double precision  NULL,
    "close"          double precision  NULL,
    "adjusted_close" double precision  NOT NULL,
    "base_asset_id"  character varying NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "daily_asset_prices_assets_daily_asset_price" FOREIGN KEY ("base_asset_id") REFERENCES "assets" ("id") ON DELETE NO ACTION
);
-- create index "dailyassetprice_base_asset_id_time" to table: "daily_asset_prices"
CREATE UNIQUE INDEX "dailyassetprice_base_asset_id_time" ON "daily_asset_prices" ("base_asset_id", "time");
