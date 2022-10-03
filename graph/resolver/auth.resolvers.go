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
func (r *mutationResolver) CreateAccount(ctx context.Context, input ent.CreateAccountInput) (*ent.Account, error) {
	return r.authService.CreateAccount(ctx, input)
}

// CreateStaffAccount is the resolver for the createStaffAccount field.
func (r *mutationResolver) CreateStaffAccount(ctx context.Context, input ent.CreateStaffAccountInput) (*ent.StaffAccount, error) {
	return r.authService.CreateStaffAccount(ctx, input)
}

// LoginToAccount is the resolver for the loginToAccount field.
func (r *mutationResolver) LoginToAccount(ctx context.Context, input auth.LoginToAccountInput) (auth.SessionID, error) {
	return r.authService.LoginToAccount(ctx, input)
}

// LoginToStaffAccount is the resolver for the loginToStaffAccount field.
func (r *mutationResolver) LoginToStaffAccount(ctx context.Context, input auth.LoginToStaffAccountInput) (auth.SessionID, error) {
	return r.authService.LoginToStaffAccount(ctx, input)
}

// LogoutFromAccount is the resolver for the logoutFromAccount field.
func (r *mutationResolver) LogoutFromAccount(ctx context.Context) (*model.Void, error) {
	return nil, r.authService.LogoutFromAccount(ctx)
}

// LogoutFromStaffAccount is the resolver for the logoutFromStaffAccount field.
func (r *mutationResolver) LogoutFromStaffAccount(ctx context.Context) (*model.Void, error) {
	return nil, r.authService.LogoutFromStaffAccount(ctx)
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
