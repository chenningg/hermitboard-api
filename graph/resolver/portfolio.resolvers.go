package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/pulid"
)

// CreatePortfolio is the resolver for the createPortfolio field.
func (r *mutationResolver) CreatePortfolio(ctx context.Context, input ent.CreatePortfolioInput) (
	*ent.Portfolio, error,
) {
	return r.portfolioService.CreatePortfolio(ctx, input)
}

// UpdatePortfolio is the resolver for the updatePortfolio field.
func (r *mutationResolver) UpdatePortfolio(
	ctx context.Context, id pulid.PULID, input ent.UpdatePortfolioInput,
) (*ent.Portfolio, error) {
	return r.portfolioService.UpdatePortfolio(ctx, id, input)
}
