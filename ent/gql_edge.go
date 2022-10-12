// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Account) Friends(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *AccountOrder, where *AccountWhereInput,
) (*AccountConnection, error) {
	opts := []AccountPaginateOption{
		WithAccountOrder(orderBy),
		WithAccountFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := a.Edges.totalCount[0][alias]
	if nodes, err := a.NamedFriends(alias); err == nil || hasTotalCount {
		pager, err := newAccountPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &AccountConnection{Edges: []*AccountEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return a.QueryFriends().Paginate(ctx, after, first, before, last, opts...)
}

func (a *Account) AuthRoles(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *AuthRoleOrder, where *AuthRoleWhereInput,
) (*AuthRoleConnection, error) {
	opts := []AuthRolePaginateOption{
		WithAuthRoleOrder(orderBy),
		WithAuthRoleFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := a.Edges.totalCount[1][alias]
	if nodes, err := a.NamedAuthRoles(alias); err == nil || hasTotalCount {
		pager, err := newAuthRolePager(opts)
		if err != nil {
			return nil, err
		}
		conn := &AuthRoleConnection{Edges: []*AuthRoleEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return a.QueryAuthRoles().Paginate(ctx, after, first, before, last, opts...)
}

func (a *Account) Portfolios(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *PortfolioOrder, where *PortfolioWhereInput,
) (*PortfolioConnection, error) {
	opts := []PortfolioPaginateOption{
		WithPortfolioOrder(orderBy),
		WithPortfolioFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := a.Edges.totalCount[2][alias]
	if nodes, err := a.NamedPortfolios(alias); err == nil || hasTotalCount {
		pager, err := newPortfolioPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &PortfolioConnection{Edges: []*PortfolioEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return a.QueryPortfolios().Paginate(ctx, after, first, before, last, opts...)
}

func (a *Account) AuthType(ctx context.Context) (*AuthType, error) {
	result, err := a.Edges.AuthTypeOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryAuthType().Only(ctx)
	}
	return result, err
}

func (a *Account) Connections(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ConnectionOrder,
) (*ConnectionConnection, error) {
	opts := []ConnectionPaginateOption{
		WithConnectionOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := a.Edges.totalCount[4][alias]
	if nodes, err := a.NamedConnections(alias); err == nil || hasTotalCount {
		pager, err := newConnectionPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &ConnectionConnection{Edges: []*ConnectionEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return a.QueryConnections().Paginate(ctx, after, first, before, last, opts...)
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

func (a *Asset) DailyAssetPrices(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *DailyAssetPriceOrder,
) (*DailyAssetPriceConnection, error) {
	opts := []DailyAssetPricePaginateOption{
		WithDailyAssetPriceOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := a.Edges.totalCount[2][alias]
	if nodes, err := a.NamedDailyAssetPrices(alias); err == nil || hasTotalCount {
		pager, err := newDailyAssetPricePager(opts)
		if err != nil {
			return nil, err
		}
		conn := &DailyAssetPriceConnection{Edges: []*DailyAssetPriceEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return a.QueryDailyAssetPrices().Paginate(ctx, after, first, before, last, opts...)
}

func (ar *AuthRole) Accounts(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *AccountOrder,
) (*AccountConnection, error) {
	opts := []AccountPaginateOption{
		WithAccountOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := ar.Edges.totalCount[0][alias]
	if nodes, err := ar.NamedAccounts(alias); err == nil || hasTotalCount {
		pager, err := newAccountPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &AccountConnection{Edges: []*AccountEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return ar.QueryAccounts().Paginate(ctx, after, first, before, last, opts...)
}

func (ar *AuthRole) StaffAccounts(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *StaffAccountOrder,
) (*StaffAccountConnection, error) {
	opts := []StaffAccountPaginateOption{
		WithStaffAccountOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := ar.Edges.totalCount[1][alias]
	if nodes, err := ar.NamedStaffAccounts(alias); err == nil || hasTotalCount {
		pager, err := newStaffAccountPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &StaffAccountConnection{Edges: []*StaffAccountEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return ar.QueryStaffAccounts().Paginate(ctx, after, first, before, last, opts...)
}

func (b *Blockchain) Cryptocurrencies(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *CryptocurrencyOrder,
) (*CryptocurrencyConnection, error) {
	opts := []CryptocurrencyPaginateOption{
		WithCryptocurrencyOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := b.Edges.totalCount[0][alias]
	if nodes, err := b.NamedCryptocurrencies(alias); err == nil || hasTotalCount {
		pager, err := newCryptocurrencyPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &CryptocurrencyConnection{Edges: []*CryptocurrencyEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return b.QueryCryptocurrencies().Paginate(ctx, after, first, before, last, opts...)
}

func (b *Blockchain) Transactions(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *TransactionOrder,
) (*TransactionConnection, error) {
	opts := []TransactionPaginateOption{
		WithTransactionOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := b.Edges.totalCount[1][alias]
	if nodes, err := b.NamedTransactions(alias); err == nil || hasTotalCount {
		pager, err := newTransactionPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &TransactionConnection{Edges: []*TransactionEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return b.QueryTransactions().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Connection) Source(ctx context.Context) (*Source, error) {
	result, err := c.Edges.SourceOrErr()
	if IsNotLoaded(err) {
		result, err = c.QuerySource().Only(ctx)
	}
	return result, err
}

func (c *Connection) Account(ctx context.Context) (*Account, error) {
	result, err := c.Edges.AccountOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryAccount().Only(ctx)
	}
	return result, err
}

func (c *Connection) Portfolios(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *PortfolioOrder,
) (*PortfolioConnection, error) {
	opts := []PortfolioPaginateOption{
		WithPortfolioOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[2][alias]
	if nodes, err := c.NamedPortfolios(alias); err == nil || hasTotalCount {
		pager, err := newPortfolioPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &PortfolioConnection{Edges: []*PortfolioEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryPortfolios().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Cryptocurrency) Asset(ctx context.Context) (*Asset, error) {
	result, err := c.Edges.AssetOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryAsset().Only(ctx)
	}
	return result, err
}

func (c *Cryptocurrency) Blockchains(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *BlockchainOrder, where *BlockchainWhereInput,
) (*BlockchainConnection, error) {
	opts := []BlockchainPaginateOption{
		WithBlockchainOrder(orderBy),
		WithBlockchainFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[1][alias]
	if nodes, err := c.NamedBlockchains(alias); err == nil || hasTotalCount {
		pager, err := newBlockchainPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &BlockchainConnection{Edges: []*BlockchainEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryBlockchains().Paginate(ctx, after, first, before, last, opts...)
}

func (dap *DailyAssetPrice) Asset(ctx context.Context) (*Asset, error) {
	result, err := dap.Edges.AssetOrErr()
	if IsNotLoaded(err) {
		result, err = dap.QueryAsset().Only(ctx)
	}
	return result, err
}

func (e *Exchange) Transactions(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *TransactionOrder, where *TransactionWhereInput,
) (*TransactionConnection, error) {
	opts := []TransactionPaginateOption{
		WithTransactionOrder(orderBy),
		WithTransactionFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := e.Edges.totalCount[0][alias]
	if nodes, err := e.NamedTransactions(alias); err == nil || hasTotalCount {
		pager, err := newTransactionPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &TransactionConnection{Edges: []*TransactionEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return e.QueryTransactions().Paginate(ctx, after, first, before, last, opts...)
}

func (po *Portfolio) Account(ctx context.Context) (*Account, error) {
	result, err := po.Edges.AccountOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryAccount().Only(ctx)
	}
	return result, err
}

func (po *Portfolio) Transactions(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *TransactionOrder,
) (*TransactionConnection, error) {
	opts := []TransactionPaginateOption{
		WithTransactionOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := po.Edges.totalCount[1][alias]
	if nodes, err := po.NamedTransactions(alias); err == nil || hasTotalCount {
		pager, err := newTransactionPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &TransactionConnection{Edges: []*TransactionEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return po.QueryTransactions().Paginate(ctx, after, first, before, last, opts...)
}

func (po *Portfolio) Connections(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ConnectionOrder, where *ConnectionWhereInput,
) (*ConnectionConnection, error) {
	opts := []ConnectionPaginateOption{
		WithConnectionOrder(orderBy),
		WithConnectionFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := po.Edges.totalCount[2][alias]
	if nodes, err := po.NamedConnections(alias); err == nil || hasTotalCount {
		pager, err := newConnectionPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &ConnectionConnection{Edges: []*ConnectionEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return po.QueryConnections().Paginate(ctx, after, first, before, last, opts...)
}

func (s *Source) Connections(ctx context.Context) ([]*Connection, error) {
	result, err := s.NamedConnections(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = s.QueryConnections().All(ctx)
	}
	return result, err
}

func (s *Source) SourceType(ctx context.Context) (*SourceType, error) {
	result, err := s.Edges.SourceTypeOrErr()
	if IsNotLoaded(err) {
		result, err = s.QuerySourceType().Only(ctx)
	}
	return result, err
}

func (st *SourceType) Sources(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *SourceOrder,
) (*SourceConnection, error) {
	opts := []SourcePaginateOption{
		WithSourceOrder(orderBy),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := st.Edges.totalCount[0][alias]
	if nodes, err := st.NamedSources(alias); err == nil || hasTotalCount {
		pager, err := newSourcePager(opts)
		if err != nil {
			return nil, err
		}
		conn := &SourceConnection{Edges: []*SourceEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return st.QuerySources().Paginate(ctx, after, first, before, last, opts...)
}

func (sa *StaffAccount) AuthRoles(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *AuthRoleOrder, where *AuthRoleWhereInput,
) (*AuthRoleConnection, error) {
	opts := []AuthRolePaginateOption{
		WithAuthRoleOrder(orderBy),
		WithAuthRoleFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := sa.Edges.totalCount[0][alias]
	if nodes, err := sa.NamedAuthRoles(alias); err == nil || hasTotalCount {
		pager, err := newAuthRolePager(opts)
		if err != nil {
			return nil, err
		}
		conn := &AuthRoleConnection{Edges: []*AuthRoleEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return sa.QueryAuthRoles().Paginate(ctx, after, first, before, last, opts...)
}

func (sa *StaffAccount) AuthType(ctx context.Context) (*AuthType, error) {
	result, err := sa.Edges.AuthTypeOrErr()
	if IsNotLoaded(err) {
		result, err = sa.QueryAuthType().Only(ctx)
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

func (t *Transaction) Blockchain(ctx context.Context) (*Blockchain, error) {
	result, err := t.Edges.BlockchainOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryBlockchain().Only(ctx)
	}
	return result, MaskNotFound(err)
}
