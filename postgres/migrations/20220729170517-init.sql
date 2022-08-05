-- +migrate Up
CREATE TABLE IF NOT EXISTS account (
    id uuid,
    nickname TEXT NOT NULL,
    email TEXT NOT NULL,
    CONSTRAINT pk_account PRIMARY KEY (id),
    CONSTRAINT uq_account_nickname UNIQUE (nickname),
    CONSTRAINT uq_account_nickname UNIQUE (email),
);

CREATE TABLE IF NOT EXISTS portfolio (
    id uuid,
    name TEXT NOT NULL,
    account_id uuid NOT NULL,
    CONSTRAINT pk_portfolio PRIMARY KEY (id),
    CONSTRAINT fk_portfolio_account FOREIGN KEY (account_id) REFERENCES account(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT uq_portfolio_account_id_name UNIQUE (account_id, name),
);

-- +migrate Down
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS portfolio;