// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Account) AuthRoles(ctx context.Context) ([]*AuthRole, error) {
	result, err := a.NamedAuthRoles(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = a.QueryAuthRoles().All(ctx)
	}
	return result, err
}

func (a *Account) Portfolios(ctx context.Context) ([]*Portfolio, error) {
	result, err := a.NamedPortfolios(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = a.QueryPortfolios().All(ctx)
	}
	return result, err
}

func (a *Account) AuthType(ctx context.Context) (*AuthType, error) {
	result, err := a.Edges.AuthTypeOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryAuthType().Only(ctx)
	}
	return result, err
}

func (a *Account) AccountAuthRoles(ctx context.Context) ([]*AccountAuthRole, error) {
	result, err := a.NamedAccountAuthRoles(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = a.QueryAccountAuthRoles().All(ctx)
	}
	return result, err
}

func (aar *AccountAuthRole) Account(ctx context.Context) (*Account, error) {
	result, err := aar.Edges.AccountOrErr()
	if IsNotLoaded(err) {
		result, err = aar.QueryAccount().Only(ctx)
	}
	return result, err
}

func (aar *AccountAuthRole) AuthRole(ctx context.Context) (*AuthRole, error) {
	result, err := aar.Edges.AuthRoleOrErr()
	if IsNotLoaded(err) {
		result, err = aar.QueryAuthRole().Only(ctx)
	}
	return result, err
}

func (a *Asset) AssetClass(ctx context.Context) (*AssetClass, error) {
	result, err := a.Edges.AssetClassOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryAssetClass().Only(ctx)
	}
	return result, err
}

func (a *Asset) Cryptocurrency(ctx context.Context) (*Cryptocurrency, error) {
	result, err := a.Edges.CryptocurrencyOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryCryptocurrency().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Asset) TransactionBase(ctx context.Context) ([]*Transaction, error) {
	result, err := a.NamedTransactionBase(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = a.QueryTransactionBase().All(ctx)
	}
	return result, err
}

func (a *Asset) TransactionQuote(ctx context.Context) ([]*Transaction, error) {
	result, err := a.NamedTransactionQuote(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = a.QueryTransactionQuote().All(ctx)
	}
	return result, err
}

func (a *Asset) DailyAssetPrice(ctx context.Context) ([]*DailyAssetPrice, error) {
	result, err := a.NamedDailyAssetPrice(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = a.QueryDailyAssetPrice().All(ctx)
	}
	return result, err
}

func (ac *AssetClass) Assets(ctx context.Context) ([]*Asset, error) {
	result, err := ac.NamedAssets(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = ac.QueryAssets().All(ctx)
	}
	return result, err
}

func (ar *AuthRole) Accounts(ctx context.Context) ([]*Account, error) {
	result, err := ar.NamedAccounts(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = ar.QueryAccounts().All(ctx)
	}
	return result, err
}

func (ar *AuthRole) StaffAccounts(ctx context.Context) ([]*StaffAccount, error) {
	result, err := ar.NamedStaffAccounts(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = ar.QueryStaffAccounts().All(ctx)
	}
	return result, err
}

func (ar *AuthRole) AccountAuthRoles(ctx context.Context) ([]*AccountAuthRole, error) {
	result, err := ar.NamedAccountAuthRoles(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = ar.QueryAccountAuthRoles().All(ctx)
	}
	return result, err
}

func (ar *AuthRole) StaffAccountAuthRoles(ctx context.Context) ([]*StaffAccountAuthRole, error) {
	result, err := ar.NamedStaffAccountAuthRoles(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = ar.QueryStaffAccountAuthRoles().All(ctx)
	}
	return result, err
}

func (at *AuthType) Account(ctx context.Context) (*Account, error) {
	result, err := at.Edges.AccountOrErr()
	if IsNotLoaded(err) {
		result, err = at.QueryAccount().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (at *AuthType) StaffAccount(ctx context.Context) (*StaffAccount, error) {
	result, err := at.Edges.StaffAccountOrErr()
	if IsNotLoaded(err) {
		result, err = at.QueryStaffAccount().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (b *Blockchain) Cryptocurrencies(ctx context.Context) ([]*Cryptocurrency, error) {
	result, err := b.NamedCryptocurrencies(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = b.QueryCryptocurrencies().All(ctx)
	}
	return result, err
}

func (b *Blockchain) BlockchainCryptocurrencies(ctx context.Context) ([]*BlockchainCryptocurrency, error) {
	result, err := b.NamedBlockchainCryptocurrencies(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = b.QueryBlockchainCryptocurrencies().All(ctx)
	}
	return result, err
}

func (bc *BlockchainCryptocurrency) Blockchain(ctx context.Context) (*Blockchain, error) {
	result, err := bc.Edges.BlockchainOrErr()
	if IsNotLoaded(err) {
		result, err = bc.QueryBlockchain().Only(ctx)
	}
	return result, err
}

func (bc *BlockchainCryptocurrency) Cryptocurrency(ctx context.Context) (*Cryptocurrency, error) {
	result, err := bc.Edges.CryptocurrencyOrErr()
	if IsNotLoaded(err) {
		result, err = bc.QueryCryptocurrency().Only(ctx)
	}
	return result, err
}

func (c *Cryptocurrency) Asset(ctx context.Context) (*Asset, error) {
	result, err := c.Edges.AssetOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryAsset().Only(ctx)
	}
	return result, err
}

func (c *Cryptocurrency) Blockchains(ctx context.Context) ([]*Blockchain, error) {
	result, err := c.NamedBlockchains(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = c.QueryBlockchains().All(ctx)
	}
	return result, err
}

func (c *Cryptocurrency) BlockchainCryptocurrencies(ctx context.Context) ([]*BlockchainCryptocurrency, error) {
	result, err := c.NamedBlockchainCryptocurrencies(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = c.QueryBlockchainCryptocurrencies().All(ctx)
	}
	return result, err
}

func (dap *DailyAssetPrice) BaseAsset(ctx context.Context) (*Asset, error) {
	result, err := dap.Edges.BaseAssetOrErr()
	if IsNotLoaded(err) {
		result, err = dap.QueryBaseAsset().Only(ctx)
	}
	return result, err
}

func (e *Exchange) Transactions(ctx context.Context) ([]*Transaction, error) {
	result, err := e.NamedTransactions(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = e.QueryTransactions().All(ctx)
	}
	return result, err
}

func (po *Portfolio) Account(ctx context.Context) (*Account, error) {
	result, err := po.Edges.AccountOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryAccount().Only(ctx)
	}
	return result, err
}

func (po *Portfolio) Transactions(ctx context.Context) ([]*Transaction, error) {
	result, err := po.NamedTransactions(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = po.QueryTransactions().All(ctx)
	}
	return result, err
}

func (sa *StaffAccount) AuthRoles(ctx context.Context) ([]*AuthRole, error) {
	result, err := sa.NamedAuthRoles(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = sa.QueryAuthRoles().All(ctx)
	}
	return result, err
}

func (sa *StaffAccount) AuthType(ctx context.Context) (*AuthType, error) {
	result, err := sa.Edges.AuthTypeOrErr()
	if IsNotLoaded(err) {
		result, err = sa.QueryAuthType().Only(ctx)
	}
	return result, err
}

func (sa *StaffAccount) StaffAccountAuthRoles(ctx context.Context) ([]*StaffAccountAuthRole, error) {
	result, err := sa.NamedStaffAccountAuthRoles(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = sa.QueryStaffAccountAuthRoles().All(ctx)
	}
	return result, err
}

func (saar *StaffAccountAuthRole) StaffAccount(ctx context.Context) (*StaffAccount, error) {
	result, err := saar.Edges.StaffAccountOrErr()
	if IsNotLoaded(err) {
		result, err = saar.QueryStaffAccount().Only(ctx)
	}
	return result, err
}

func (saar *StaffAccountAuthRole) AuthRole(ctx context.Context) (*AuthRole, error) {
	result, err := saar.Edges.AuthRoleOrErr()
	if IsNotLoaded(err) {
		result, err = saar.QueryAuthRole().Only(ctx)
	}
	return result, err
}

func (t *Transaction) TransactionType(ctx context.Context) (*TransactionType, error) {
	result, err := t.Edges.TransactionTypeOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTransactionType().Only(ctx)
	}
	return result, err
}

func (t *Transaction) BaseAsset(ctx context.Context) (*Asset, error) {
	result, err := t.Edges.BaseAssetOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryBaseAsset().Only(ctx)
	}
	return result, err
}

func (t *Transaction) QuoteAsset(ctx context.Context) (*Asset, error) {
	result, err := t.Edges.QuoteAssetOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryQuoteAsset().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Transaction) Portfolio(ctx context.Context) (*Portfolio, error) {
	result, err := t.Edges.PortfolioOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryPortfolio().Only(ctx)
	}
	return result, err
}

func (t *Transaction) Exchange(ctx context.Context) (*Exchange, error) {
	result, err := t.Edges.ExchangeOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryExchange().Only(ctx)
	}
	return result, err
}

func (tt *TransactionType) Transactions(ctx context.Context) ([]*Transaction, error) {
	result, err := tt.NamedTransactions(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = tt.QueryTransactions().All(ctx)
	}
	return result, err
}
