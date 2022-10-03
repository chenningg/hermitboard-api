package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/graph"
	"github.com/chenningg/hermitboard-api/pulid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id pulid.PULID) (ent.Noder, error) {
	return r.dbService.Client().Noder(ctx, id, ent.WithNodeType(ent.IDToType))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []pulid.PULID) ([]ent.Noder, error) {
	return r.dbService.Client().Noders(ctx, ids, ent.WithNodeType(ent.IDToType))
}

// Accounts is the resolver for the accounts field.
func (r *queryResolver) Accounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AccountOrder, where *ent.AccountWhereInput) (*ent.AccountConnection, error) {
	return r.dbService.Client().Account.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithAccountOrder(orderBy),
			ent.WithAccountFilter(where.Filter),
		)
}

// Assets is the resolver for the assets field.
func (r *queryResolver) Assets(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssetOrder, where *ent.AssetWhereInput) (*ent.AssetConnection, error) {
	return r.dbService.Client().Asset.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithAssetOrder(orderBy),
			ent.WithAssetFilter(where.Filter),
		)
}

// AssetClasses is the resolver for the assetClasses field.
func (r *queryResolver) AssetClasses(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssetClassOrder, where *ent.AssetClassWhereInput) (*ent.AssetClassConnection, error) {
	return r.dbService.Client().AssetClass.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithAssetClassOrder(orderBy),
			ent.WithAssetClassFilter(where.Filter),
		)
}

// AuthRoles is the resolver for the authRoles field.
func (r *queryResolver) AuthRoles(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AuthRoleOrder, where *ent.AuthRoleWhereInput) (*ent.AuthRoleConnection, error) {
	return r.dbService.Client().AuthRole.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithAuthRoleOrder(orderBy),
			ent.WithAuthRoleFilter(where.Filter),
		)
}

// AuthTypes is the resolver for the authTypes field.
func (r *queryResolver) AuthTypes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AuthTypeOrder, where *ent.AuthTypeWhereInput) (*ent.AuthTypeConnection, error) {
	return r.dbService.Client().AuthType.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithAuthTypeOrder(orderBy),
			ent.WithAuthTypeFilter(where.Filter),
		)
}

// Blockchains is the resolver for the blockchains field.
func (r *queryResolver) Blockchains(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.BlockchainOrder, where *ent.BlockchainWhereInput) (*ent.BlockchainConnection, error) {
	return r.dbService.Client().Blockchain.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithBlockchainOrder(orderBy),
			ent.WithBlockchainFilter(where.Filter),
		)
}

// Connections is the resolver for the connections field.
func (r *queryResolver) Connections(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ConnectionOrder, where *ent.ConnectionWhereInput) (*ent.ConnectionConnection, error) {
	return r.dbService.Client().Connection.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithConnectionOrder(orderBy),
			ent.WithConnectionFilter(where.Filter),
		)
}

// Cryptocurrencies is the resolver for the cryptocurrencies field.
func (r *queryResolver) Cryptocurrencies(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CryptocurrencyOrder, where *ent.CryptocurrencyWhereInput) (*ent.CryptocurrencyConnection, error) {
	return r.dbService.Client().Cryptocurrency.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithCryptocurrencyOrder(orderBy),
			ent.WithCryptocurrencyFilter(where.Filter),
		)
}

// DailyAssetPrices is the resolver for the dailyAssetPrices field.
func (r *queryResolver) DailyAssetPrices(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.DailyAssetPriceOrder, where *ent.DailyAssetPriceWhereInput) (*ent.DailyAssetPriceConnection, error) {
	return r.dbService.Client().DailyAssetPrice.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithDailyAssetPriceOrder(orderBy),
			ent.WithDailyAssetPriceFilter(where.Filter),
		)
}

// Exchanges is the resolver for the exchanges field.
func (r *queryResolver) Exchanges(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ExchangeOrder, where *ent.ExchangeWhereInput) (*ent.ExchangeConnection, error) {
	return r.dbService.Client().Exchange.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithExchangeOrder(orderBy),
			ent.WithExchangeFilter(where.Filter),
		)
}

// Portfolios is the resolver for the portfolios field.
func (r *queryResolver) Portfolios(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PortfolioOrder, where *ent.PortfolioWhereInput) (*ent.PortfolioConnection, error) {
	return r.dbService.Client().Portfolio.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithPortfolioOrder(orderBy),
			ent.WithPortfolioFilter(where.Filter),
		)
}

// Sources is the resolver for the sources field.
func (r *queryResolver) Sources(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SourceOrder, where *ent.SourceWhereInput) (*ent.SourceConnection, error) {
	return r.dbService.Client().Source.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithSourceOrder(orderBy),
			ent.WithSourceFilter(where.Filter),
		)
}

// SourceTypes is the resolver for the sourceTypes field.
func (r *queryResolver) SourceTypes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SourceTypeOrder, where *ent.SourceTypeWhereInput) (*ent.SourceTypeConnection, error) {
	return r.dbService.Client().SourceType.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithSourceTypeOrder(orderBy),
			ent.WithSourceTypeFilter(where.Filter),
		)
}

// StaffAccounts is the resolver for the staffAccounts field.
func (r *queryResolver) StaffAccounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.StaffAccountOrder, where *ent.StaffAccountWhereInput) (*ent.StaffAccountConnection, error) {
	return r.dbService.Client().StaffAccount.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithStaffAccountOrder(orderBy),
			ent.WithStaffAccountFilter(where.Filter),
		)
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TransactionOrder, where *ent.TransactionWhereInput) (*ent.TransactionConnection, error) {
	return r.dbService.Client().Transaction.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithTransactionOrder(orderBy),
			ent.WithTransactionFilter(where.Filter),
		)
}

// TransactionTypes is the resolver for the transactionTypes field.
func (r *queryResolver) TransactionTypes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TransactionTypeOrder, where *ent.TransactionTypeWhereInput) (*ent.TransactionTypeConnection, error) {
	return r.dbService.Client().TransactionType.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithTransactionTypeOrder(orderBy),
			ent.WithTransactionTypeFilter(where.Filter),
		)
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
