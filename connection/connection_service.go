package connection

import (
	"context"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/go-logr/logr"
)

type ConnectionServicer interface {
	CreatePortfolio(
		ctx context.Context, input ent.CreateConnectionInput,
	) (*ent.Connection, error)
	UpdatePortfolio(ctx context.Context, input ent.UpdateConnectionInput) (*ent.Connection, error)
	DeletePortfolio(ctx context.Context, input DeleteConnectionInput) error
}

type ConnectionService struct {
	logger logr.Logger
}

func NewConnectionService(logger logr.Logger) *ConnectionService {
	var connectionService = new(ConnectionService)

	// Initialize logger.
	connectionService.logger = logger
	connectionService.logger.V(2).Info("NewConnectionService(): initializing connection service")

	return connectionService
}
