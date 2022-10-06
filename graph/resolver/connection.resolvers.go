package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/pulid"
)

// CreateConnection is the resolver for the createConnection field.
func (r *mutationResolver) CreateConnection(ctx context.Context, input ent.CreateConnectionInput) (*ent.Connection, error) {
	return r.connectionService.CreateConnection(ctx, input)
}

// UpdateConnection is the resolver for the updateConnection field.
func (r *mutationResolver) UpdateConnection(ctx context.Context, id pulid.PULID, input ent.UpdateConnectionInput) (*ent.Connection, error) {
	return r.connectionService.UpdateConnection(ctx, id, input)
}
