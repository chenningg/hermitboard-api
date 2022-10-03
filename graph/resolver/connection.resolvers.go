package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/connection"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/graph/model"
)

// CreateConnection is the resolver for the createConnection field.
func (r *mutationResolver) CreateConnection(ctx context.Context, input ent.CreateConnectionInput) (*ent.Connection, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateConnection is the resolver for the updateConnection field.
func (r *mutationResolver) UpdateConnection(ctx context.Context, input ent.UpdateConnectionInput) (*ent.Connection, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteConnection is the resolver for the deleteConnection field.
func (r *mutationResolver) DeleteConnection(ctx context.Context, input connection.DeleteConnectionInput) (*model.Void, error) {
	panic(fmt.Errorf("not implemented"))
}
