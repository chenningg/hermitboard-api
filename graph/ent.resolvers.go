package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/graph/generated"
	"github.com/chenningg/hermitboard-api/pulid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id pulid.PULID) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []pulid.PULID) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Accounts is the resolver for the accounts field.
func (r *queryResolver) Accounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AccountOrder, where *ent.AccountWhereInput) (*ent.AccountConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Assets is the resolver for the assets field.
func (r *queryResolver) Assets(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssetOrder, where *ent.AssetWhereInput) (*ent.AssetConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Assetclasses is the resolver for the assetclasses field.
func (r *queryResolver) Assetclasses(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssetClassOrder, where *ent.AssetClassWhereInput) (*ent.AssetClassConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Authroles is the resolver for the authroles field.
func (r *queryResolver) Authroles(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AuthRoleOrder, where *ent.AuthRoleWhereInput) (*ent.AuthRoleConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Authtypes is the resolver for the authtypes field.
func (r *queryResolver) Authtypes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AuthTypeOrder, where *ent.AuthTypeWhereInput) (*ent.AuthTypeConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Blockchains is the resolver for the blockchains field.
func (r *queryResolver) Blockchains(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.BlockchainOrder, where *ent.BlockchainWhereInput) (*ent.BlockchainConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Connections is the resolver for the connections field.
func (r *queryResolver) Connections(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ConnectionOrder, where *ent.ConnectionWhereInput) (*ent.ConnectionConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Cryptocurrencies is the resolver for the cryptocurrencies field.
func (r *queryResolver) Cryptocurrencies(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CryptocurrencyOrder, where *ent.CryptocurrencyWhereInput) (*ent.CryptocurrencyConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Dailyassetprices is the resolver for the dailyassetprices field.
func (r *queryResolver) Dailyassetprices(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.DailyAssetPriceOrder, where *ent.DailyAssetPriceWhereInput) (*ent.DailyAssetPriceConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Exchanges is the resolver for the exchanges field.
func (r *queryResolver) Exchanges(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ExchangeOrder, where *ent.ExchangeWhereInput) (*ent.ExchangeConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Portfolios is the resolver for the portfolios field.
func (r *queryResolver) Portfolios(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PortfolioOrder, where *ent.PortfolioWhereInput) (*ent.PortfolioConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Sources is the resolver for the sources field.
func (r *queryResolver) Sources(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SourceOrder, where *ent.SourceWhereInput) (*ent.SourceConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Sourcetypes is the resolver for the sourcetypes field.
func (r *queryResolver) Sourcetypes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SourceTypeOrder, where *ent.SourceTypeWhereInput) (*ent.SourceTypeConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Staffaccounts is the resolver for the staffaccounts field.
func (r *queryResolver) Staffaccounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.StaffAccountOrder, where *ent.StaffAccountWhereInput) (*ent.StaffAccountConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TransactionOrder, where *ent.TransactionWhereInput) (*ent.TransactionConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Transactiontypes is the resolver for the transactiontypes field.
func (r *queryResolver) Transactiontypes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TransactionTypeOrder, where *ent.TransactionTypeWhereInput) (*ent.TransactionTypeConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
