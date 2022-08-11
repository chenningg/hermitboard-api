-- +migrate Up
CREATE TABLE IF NOT EXISTS account(
    id uuid,
    auth_id TEXT NOT NULL,
    nickname TEXT NOT NULL,
    email TEXT NOT NULL,
    
    CONSTRAINT pk_account PRIMARY KEY (id),
    CONSTRAINT uq_account_auth_id UNIQUE (auth_id),
    CONSTRAINT uq_account_nickname UNIQUE (nickname),
    CONSTRAINT uq_account_email UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS portfolio(
    id uuid,
    name TEXT NOT NULL,
    account_id uuid NOT NULL,
    is_public boolean NOT NULL,
    is_visible boolean NOT NULL, 
    CONSTRAINT pk_portfolio PRIMARY KEY(id),
    CONSTRAINT fk_portfolio_account FOREIGN KEY(account_id) REFERENCES account(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT uq_portfolio_account_id_name UNIQUE(account_id, name)
);

CREATE TABLE IF NOT EXISTS asset_class(
    id TEXT,
    description TEXT,
    CONSTRAINT pk_asset_class PRIMARY KEY(id)
);

INSERT INTO asset_class(id, description)
VALUES
    ('cash or cash equivalents', 'Value of assets that are cash or can be converted to cash immediately.'),
    ('commodity', 'A raw material or primary agricultural product that can be bought and sold.'),
    ('cryptocurrency', 'A digital or virtual currency secured by cryptography.'),
    ('equity', 'The ownership of assets with financial value that may have debts or other liabilities attached to them.'),
    ('fixed income', 'Assets and securities that pay out a set level of cash flows to investors.'),
    ('futures', 'Derivative financial contracts that obligate parties to buy or sell an asset at a predetermined future date and price.'),
    ('real estate', 'Property consisting of land and buildings.');

CREATE TABLE IF NOT EXISTS asset (
    id uuid,
    asset_class_id TEXT NOT NULL,
    CONSTRAINT fk_asset_asset_class FOREIGN KEY(asset_class_id) REFERENCES asset_class(id) ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS portfolio;
DROP TABLE IF EXISTS asset;
DROP TABLE IF EXISTS asset_class;