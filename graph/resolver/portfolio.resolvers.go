package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/graph/model"
	"github.com/chenningg/hermitboard-api/portfolio"
)

// CreatePortfolio is the resolver for the createPortfolio field.
func (r *mutationResolver) CreatePortfolio(ctx context.Context, input ent.CreatePortfolioInput) (*ent.Portfolio, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdatePortfolio is the resolver for the updatePortfolio field.
func (r *mutationResolver) UpdatePortfolio(ctx context.Context, input ent.UpdatePortfolioInput) (*ent.Portfolio, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeletePortfolio is the resolver for the deletePortfolio field.
func (r *mutationResolver) DeletePortfolio(ctx context.Context, input portfolio.DeletePortfolioInput) (*model.Void, error) {
	panic(fmt.Errorf("not implemented"))
}
