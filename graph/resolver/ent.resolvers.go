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
	return r.db.Client().Noder(ctx, id, ent.WithNodeType(ent.IDToType))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []pulid.PULID) ([]ent.Noder, error) {
	return r.db.Client().Noders(ctx, ids, ent.WithNodeType(ent.IDToType))
}

// Accounts is the resolver for the accounts field.
func (r *queryResolver) Accounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AccountOrder, where *ent.AccountWhereInput) (*ent.AccountConnection, error) {
	return r.db.Client().Account.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithAccountOrder(orderBy),
			ent.WithAccountFilter(where.Filter),
		)
}

// Assets is the resolver for the assets field.
func (r *queryResolver) Assets(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssetOrder, where *ent.AssetWhereInput) (*ent.AssetConnection, error) {
	return r.db.Client().Asset.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithAssetOrder(orderBy),
			ent.WithAssetFilter(where.Filter),
		)
}

// Assetclasses is the resolver for the assetclasses field.
func (r *queryResolver) Assetclasses(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssetClassOrder, where *ent.AssetClassWhereInput) (*ent.AssetClassConnection, error) {
	return r.db.Client().AssetClass.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithAssetClassOrder(orderBy),
			ent.WithAssetClassFilter(where.Filter),
		)
}

// Authroles is the resolver for the authroles field.
func (r *queryResolver) Authroles(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AuthRoleOrder, where *ent.AuthRoleWhereInput) (*ent.AuthRoleConnection, error) {
	return r.db.Client().AuthRole.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithAuthRoleOrder(orderBy),
			ent.WithAuthRoleFilter(where.Filter),
		)
}

// Authtypes is the resolver for the authtypes field.
func (r *queryResolver) Authtypes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AuthTypeOrder, where *ent.AuthTypeWhereInput) (*ent.AuthTypeConnection, error) {
	return r.db.Client().AuthType.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithAuthTypeOrder(orderBy),
			ent.WithAuthTypeFilter(where.Filter),
		)
}

// Blockchains is the resolver for the blockchains field.
func (r *queryResolver) Blockchains(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.BlockchainOrder, where *ent.BlockchainWhereInput) (*ent.BlockchainConnection, error) {
	return r.db.Client().Blockchain.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithBlockchainOrder(orderBy),
			ent.WithBlockchainFilter(where.Filter),
		)
}

// Connections is the resolver for the connections field.
func (r *queryResolver) Connections(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ConnectionOrder, where *ent.ConnectionWhereInput) (*ent.ConnectionConnection, error) {
	return r.db.Client().Connection.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithConnectionOrder(orderBy),
			ent.WithConnectionFilter(where.Filter),
		)
}

// Cryptocurrencies is the resolver for the cryptocurrencies field.
func (r *queryResolver) Cryptocurrencies(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CryptocurrencyOrder, where *ent.CryptocurrencyWhereInput) (*ent.CryptocurrencyConnection, error) {
	return r.db.Client().Cryptocurrency.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithCryptocurrencyOrder(orderBy),
			ent.WithCryptocurrencyFilter(where.Filter),
		)
}

// Dailyassetprices is the resolver for the dailyassetprices field.
func (r *queryResolver) Dailyassetprices(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.DailyAssetPriceOrder, where *ent.DailyAssetPriceWhereInput) (*ent.DailyAssetPriceConnection, error) {
	return r.db.Client().DailyAssetPrice.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithDailyAssetPriceOrder(orderBy),
			ent.WithDailyAssetPriceFilter(where.Filter),
		)
}

// Exchanges is the resolver for the exchanges field.
func (r *queryResolver) Exchanges(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ExchangeOrder, where *ent.ExchangeWhereInput) (*ent.ExchangeConnection, error) {
	return r.db.Client().Exchange.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithExchangeOrder(orderBy),
			ent.WithExchangeFilter(where.Filter),
		)
}

// Portfolios is the resolver for the portfolios field.
func (r *queryResolver) Portfolios(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PortfolioOrder, where *ent.PortfolioWhereInput) (*ent.PortfolioConnection, error) {
	return r.db.Client().Portfolio.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithPortfolioOrder(orderBy),
			ent.WithPortfolioFilter(where.Filter),
		)
}

// Sources is the resolver for the sources field.
func (r *queryResolver) Sources(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SourceOrder, where *ent.SourceWhereInput) (*ent.SourceConnection, error) {
	return r.db.Client().Source.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithSourceOrder(orderBy),
			ent.WithSourceFilter(where.Filter),
		)
}

// Sourcetypes is the resolver for the sourcetypes field.
func (r *queryResolver) Sourcetypes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SourceTypeOrder, where *ent.SourceTypeWhereInput) (*ent.SourceTypeConnection, error) {
	return r.db.Client().SourceType.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithSourceTypeOrder(orderBy),
			ent.WithSourceTypeFilter(where.Filter),
		)
}

// Staffaccounts is the resolver for the staffaccounts field.
func (r *queryResolver) Staffaccounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.StaffAccountOrder, where *ent.StaffAccountWhereInput) (*ent.StaffAccountConnection, error) {
	return r.db.Client().StaffAccount.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithStaffAccountOrder(orderBy),
			ent.WithStaffAccountFilter(where.Filter),
		)
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TransactionOrder, where *ent.TransactionWhereInput) (*ent.TransactionConnection, error) {
	return r.db.Client().Transaction.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithTransactionOrder(orderBy),
			ent.WithTransactionFilter(where.Filter),
		)
}

// Transactiontypes is the resolver for the transactiontypes field.
func (r *queryResolver) Transactiontypes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TransactionTypeOrder, where *ent.TransactionTypeWhereInput) (*ent.TransactionTypeConnection, error) {
	return r.db.Client().TransactionType.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithTransactionTypeOrder(orderBy),
			ent.WithTransactionTypeFilter(where.Filter),
		)
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
