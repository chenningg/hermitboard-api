package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/chenningg/hermitboard-api/auth"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/graph"
	"github.com/chenningg/hermitboard-api/graph/model"
)

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input ent.CreateAccountInput) (*auth.CreateAccountPayload, error) {
	return r.authService.CreateAccount(ctx, input)
}

// CreateStaffAccount is the resolver for the createStaffAccount field.
func (r *mutationResolver) CreateStaffAccount(ctx context.Context, input ent.CreateStaffAccountInput) (*auth.CreateStaffAccountPayload, error) {
	return r.authService.CreateStaffAccount(ctx, input)
}

// LoginToAccount is the resolver for the loginToAccount field.
func (r *mutationResolver) LoginToAccount(ctx context.Context, input auth.LoginToAccountInput) (*auth.LoginToAccountPayload, error) {
	return r.authService.LoginToAccount(ctx, input)
}

// LoginToStaffAccount is the resolver for the loginToStaffAccount field.
func (r *mutationResolver) LoginToStaffAccount(ctx context.Context, input auth.LoginToStaffAccountInput) (*auth.LoginToStaffAccountPayload, error) {
	return r.authService.LoginToStaffAccount(ctx, input)
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context) (*model.Void, error) {
	if err := r.authService.Logout(ctx); err != nil {
		return nil, err
	}
	return nil, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
