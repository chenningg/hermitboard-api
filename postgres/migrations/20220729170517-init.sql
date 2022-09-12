-- +migrate Up

-- An account represents a user in the application.
CREATE TABLE IF NOT EXISTS account (
    id uuid,
    auth_id TEXT, -- For external auth (e.g. Firebase).
    nickname TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT, -- Hashed and salted password for internal auth.
    
    CONSTRAINT pk_account PRIMARY KEY (id),
    CONSTRAINT uq_account_auth_id UNIQUE (auth_id),
    CONSTRAINT uq_account_nickname UNIQUE (nickname),
    CONSTRAINT uq_account_email UNIQUE (email),
    CONSTRAINT chk_account_auth_id_password_not_null CHECK (auth_id is not null or password is not null)
);

-- An auth_role is an authorization role for an account.
CREATE TABLE IF NOT EXISTS auth_role (
    id text,
    description TEXT,

    CONSTRAINT pk_auth_role PRIMARY KEY(id)
);

INSERT INTO auth_role(id, description)
VALUES
    ('Demo', 'A demo role with restricted rights.'),
    ('Free', 'A free role with limited feature access.'),
    ('Plus', 'A paying Plus role.'),
    ('Pro', 'A paying Pro role.'),
    ('Enterprise', 'A paying Enterprise role.'),
    ('Support', 'A customer support role with limited permissions.'),
    ('Admin', 'An administrative role with elevated permissions.'),
    ('SuperAdmin', 'A super administrative role with full rights.');

-- Links an account to one or more auth_roles.
CREATE TABLE IF NOT EXISTS account_auth_role (
    id uuid,
    account_id uuid NOT NULL,
    auth_role_id TEXT NOT NULL,

    CONSTRAINT pk_account_auth_role PRIMARY KEY(id),
    CONSTRAINT fk_account_auth_role_account FOREIGN KEY(account_id) REFERENCES account(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_account_auth_role_auth_role FOREIGN KEY(auth_role_id) REFERENCES auth_role(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- An account may hold multiple portfolios (collections of assets).
CREATE TABLE IF NOT EXISTS portfolio (
    id uuid,
    name TEXT NOT NULL,
    account_id uuid NOT NULL,
    is_public boolean NOT NULL DEFAULT FALSE, -- Whether portfolio can be seen by others.
    is_visible boolean NOT NULL DEFAULT TRUE, -- Whether portfolio can be seen in one's account.

    CONSTRAINT pk_portfolio PRIMARY KEY(id),
    CONSTRAINT fk_portfolio_account FOREIGN KEY(account_id) REFERENCES account(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT uq_portfolio_account_id_name UNIQUE(account_id, name)
);

-- Denotes the various types of asset classes.
CREATE TABLE IF NOT EXISTS asset_class (
    id TEXT,
    description TEXT,

    CONSTRAINT pk_asset_class PRIMARY KEY(id)
);

INSERT INTO asset_class(id, description)
VALUES
    ('CashOrCashEquivalent', 'Value of assets that are cash or can be converted to cash immediately.'),
    ('Commodity', 'A raw material or primary agricultural product that can be bought and sold.'),
    ('Cryptocurrency', 'A digital or virtual currency secured by cryptography.'),
    ('Equity', 'The ownership of assets with financial value that may have debts or other liabilities attached to them.'),
    ('FixedIncome', 'Assets and securities that pay out a set level of cash flows to investors.'),
    ('Future', 'Derivative financial contracts that obligate parties to buy or sell an asset at a predetermined future date and price.'),
    ('RealEstate', 'Property consisting of land and buildings.');

-- A specific asset type derived from an asset.
CREATE TABLE IF NOT EXISTS cryptocurrency (
    id uuid,
    symbol TEXT NOT NULL,
    icon TEXT,
    short_name TEXT NOT NULL,
    long_name TEXT NOT NULL,
    
    CONSTRAINT pk_cryptocurrency PRIMARY KEY (id)
);

-- A generic entity representing an asset (polymorphic association).
-- Holds foreign keys in an exclusive relationship to its polymorphic children.
CREATE TABLE IF NOT EXISTS asset (
    id uuid,
    asset_class_id TEXT NOT NULL,
    cryptocurrency_id uuid,

    CONSTRAINT pk_asset PRIMARY KEY(id),
    CONSTRAINT fk_asset_asset_class FOREIGN KEY(asset_class_id) REFERENCES asset_class(id) ON UPDATE CASCADE,
    CONSTRAINT fk_asset_cryptocurrency FOREIGN KEY(cryptocurrency_id) REFERENCES cryptocurrency(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT chk_asset_polymorphic_children_not_null CHECK (
        (
            (cryptocurrency_id IS NOT NULL)::integer
        ) = 1
    )
);

-- Denotes a specific blockchain.
CREATE TABLE IF NOT EXISTS blockchain (
    id uuid,
    name TEXT NOT NULL,
    symbol TEXT,
    icon TEXT,
    cryptocurrency_id uuid, -- The native token of the blockchain.
    chain_id BIGINT,

    CONSTRAINT pk_blockchain PRIMARY KEY (id),
    CONSTRAINT fk_blockchain_cryptocurrency FOREIGN KEY(cryptocurrency_id) REFERENCES cryptocurrency(id) ON UPDATE CASCADE
);

-- Links cryptocurrencies to one or more blockchains.
CREATE TABLE IF NOT EXISTS blockchain_cryptocurrency (
    id uuid,
    blockchain_id uuid NOT NULL,
    cryptocurrency_id uuid NOT NULL,

    CONSTRAINT pk_blockchain_cryptocurrency PRIMARY KEY (id),
    CONSTRAINT fk_blockchain_cryptocurrency_blockchain FOREIGN KEY(blockchain_id) REFERENCES blockchain(id) ON UPDATE CASCADE,
    CONSTRAINT fk_blockchain_cryptocurrency_cryptocurrency FOREIGN KEY(cryptocurrency_id) REFERENCES cryptocurrency(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Represents an exchange.
CREATE TABLE IF NOT EXISTS exchange (
    id uuid,
    name TEXT NOT NULL,
    url TEXT,

    CONSTRAINT pk_exchange PRIMARY KEY (id)
);

-- Denotes the various types of transactions.
CREATE TABLE IF NOT EXISTS transaction_type (
    id TEXT,
    description TEXT,

    CONSTRAINT pk_transaction_type PRIMARY KEY(id)
);

INSERT INTO transaction_type(id, description)
VALUES
    ('Buy', 'A buy transaction.'),
    ('Sell', 'A sell transaction.'),
    ('Stake', 'Staking of a cryptocurrency.'),
    ('DividendIncome', 'Income from dividends.'),
    ('RentPayment', 'Outflow of cash due to rent payment.'),
    ('RentIncome', 'Inflow of cash due to rent collection.'),
    ('StockDividend', 'Dividends paid out in stock.');

-- Represents a transaction. A transaction may contain arbitrary data for additional information.
CREATE TABLE IF NOT EXISTS transaction (
    id uuid,
    time TIMESTAMPTZ NOT NULL,
    transaction_type_id TEXT NOT NULL,
    base_asset_id uuid NOT NULL,
    quote_asset_id uuid,
    portfolio_id uuid NOT NULL,
    exchange_id uuid,
    units NUMERIC NOT NULL,
    price_per_unit NUMERIC NOT NULL,
    data JSONB, -- May contain arbitrary data such as blockchain of the exchange.

    CONSTRAINT pk_transaction PRIMARY KEY (id),
    CONSTRAINT fk_transaction_transaction_type FOREIGN KEY(transaction_type_id) REFERENCES transaction_type(id) ON UPDATE CASCADE,
    CONSTRAINT fk_transaction_base_asset FOREIGN KEY(base_asset_id) REFERENCES asset(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_transaction_quote_asset FOREIGN KEY(quote_asset_id) REFERENCES asset(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_transaction_portfolio FOREIGN KEY(portfolio_id) REFERENCES portfolio(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_transaction_exchange FOREIGN KEY(exchange_id) REFERENCES exchange(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Daily asset price of an asset.
CREATE TABLE IF NOT EXISTS daily_asset_price (
    id uuid,
    base_asset_id uuid NOT NULL,
    time TIMESTAMPTZ NOT NULL,
    open NUMERIC,
    high NUMERIC,
    low NUMERIC,
    close NUMERIC,
    adjusted_close NUMERIC NOT NULL,

    CONSTRAINT pk_daily_asset_price PRIMARY KEY (id),
    CONSTRAINT fk_daily_asset_price_base_asset FOREIGN KEY(base_asset_id) REFERENCES asset(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT uq_daily_asset_price_base_asset_id_time UNIQUE (base_asset_id, time)
);

-- +migrate Down
DROP TABLE IF EXISTS account_auth_role;
DROP TABLE IF EXISTS auth_role;
DROP TABLE IF EXISTS blockchain_cryptocurrency;
DROP TABLE IF EXISTS transaction;
DROP TABLE IF EXISTS transaction_type;
DROP TABLE IF EXISTS daily_asset_price;
DROP TABLE IF EXISTS portfolio;
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS asset;
DROP TABLE IF EXISTS asset_class;
DROP TABLE IF EXISTS blockchain;
DROP TABLE IF EXISTS exchange;
DROP TABLE IF EXISTS cryptocurrency;