package connection

import (
	"context"
	"errors"
	"fmt"

	"github.com/chenningg/hermitboard-api/auth"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/connection"
	"github.com/chenningg/hermitboard-api/ent/portfolio"
	"github.com/chenningg/hermitboard-api/pulid"
	"github.com/chenningg/hermitboard-api/resperror"
	"github.com/go-logr/logr"
)

type ConnectionServicer interface {
	CreateConnection(
		ctx context.Context, input ent.CreateConnectionInput,
	) (*ent.Connection, error)
	UpdateConnection(ctx context.Context, id pulid.PULID, input ent.UpdateConnectionInput) (*ent.Connection, error)
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

func (connectionService ConnectionService) CreateConnection(
	ctx context.Context, input ent.CreateConnectionInput,
) (*ent.Connection, error) {
	// Check for empty fields.
	if input.Name == "" || input.AccessToken == "" {
		return nil, fmt.Errorf(
			"connection.CreateConnection(): input fields are invalid or missing: %w",
			resperror.NewGraphQLError("missing fields to create a connection", resperror.GQLBadUserInput),
		)
	}

	dbClient := ent.FromContext(ctx)
	session := auth.GetSessionFromContext(ctx)

	// Must be logged in and have non-staff auth role to create a connection.
	if !auth.IsLoggedIn(session) || !auth.HasAuthRoles(session, auth.NonStaffAuthRoles...) {
		return nil, fmt.Errorf(
			"connection.CreateConnection(): must be logged in as a non-staff user to create a connection: %w",
			resperror.NewGraphQLError("unauthorized to create a connection", resperror.GQLForbidden),
		)
	}

	// Check what portfolios exist on the account.
	if len(input.PortfolioIDs) > 0 {
		portfoliosExisting, err := dbClient.Portfolio.Query().Where(portfolio.AccountIDEQ(session.UserID)).All(ctx)
		if err != nil {
			if errors.Is(err, &ent.NotFoundError{}) {
				return nil, fmt.Errorf(
					"connection.CreateConnection(): could not add connection as no portfolios were found on this account: %w",
					resperror.NewGraphQLError("missing portfolio to add connection to", resperror.GQLBadUserInput),
				)

			}
			return nil, fmt.Errorf(
				"connection.CreateConnection(): could not query account's portfolios: %w",
				resperror.NewGraphQLError(
					fmt.Sprintf("unable to retrieve portfolios %d", len(input.PortfolioIDs)),
					resperror.GQLInternalServerError,
				),
			)
		}

		var portfoliosMap = map[pulid.PULID]struct{}{}
		for _, portfolio := range portfoliosExisting {
			portfoliosMap[portfolio.ID] = struct{}{}
		}

		for _, portfolioID := range input.PortfolioIDs {
			if _, ok := portfoliosMap[portfolioID]; !ok {
				return nil, fmt.Errorf(
					"connection.CreateConnection(): invalid portfolio IDs provided: %w",
					resperror.NewGraphQLError("invalid portfolio(s) to add connection to", resperror.GQLBadUserInput),
				)
			}
		}
	}

	// Create the new connection
	if input.RefreshToken != nil && *input.RefreshToken == "" {
		input.RefreshToken = nil
	}

	newConnection, err := dbClient.Connection.Create().
		SetInput(input).
		SetAccountID(session.UserID).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"connection.CreateConnection(): could not create connection: %w",
			resperror.NewGraphQLError("could not create connection", resperror.GQLInternalServerError),
		)
	}

	return newConnection, nil
}

func (connectionService ConnectionService) UpdateConnection(
	ctx context.Context, id pulid.PULID, input ent.UpdateConnectionInput,
) (*ent.Connection, error) {
	dbClient := ent.FromContext(ctx)
	session := auth.GetSessionFromContext(ctx)

	// Must be logged in and have non-staff auth role to update a connection.
	if !auth.IsLoggedIn(session) || !auth.HasAuthRoles(session, auth.NonStaffAuthRoles...) {
		return nil, fmt.Errorf(
			"connection.UpdateConnection(): must be logged in as a non-staff user to update a connection: %w",
			resperror.NewGraphQLError("unauthorized", resperror.GQLForbidden),
		)
	}

	// Check if connection ID even exists on this account.
	_, err := dbClient.Connection.Query().Where(
		connection.And(
			connection.AccountIDEQ(session.UserID),
			connection.IDEQ(id),
		),
	).Only(ctx)
	if err != nil {
		if errors.Is(err, &ent.NotFoundError{}) {
			return nil, fmt.Errorf(
				"connection.UpdateConnection(): could not update connection as no matching connection was found on this account: %w",
				resperror.NewGraphQLError("no connection found", resperror.GQLBadUserInput),
			)
		}
		return nil, fmt.Errorf(
			"connection.UpdateConnection(): could not query account's connections: %w",
			resperror.NewGraphQLError("could not retrieve connections", resperror.GQLInternalServerError),
		)
	}

	// Check what portfolios exist on the account.
	if len(input.AddPortfolioIDs) > 0 || len(input.RemovePortfolioIDs) > 0 {
		portfoliosExisting, err := dbClient.Portfolio.Query().Where(portfolio.AccountIDEQ(session.UserID)).All(ctx)
		if err != nil {
			if errors.Is(err, &ent.NotFoundError{}) {
				return nil, fmt.Errorf(
					"connection.UpdateConnection(): could not update connection as no portfolios were found on this account: %w",
					resperror.NewGraphQLError("no portfolios found", resperror.GQLBadUserInput),
				)
			}
			return nil, fmt.Errorf(
				"connection.UpdateConnection(): could not query account's portfolios: %w",
				resperror.NewGraphQLError("could not retrieve portfolios", resperror.GQLInternalServerError),
			)
		}

		var portfoliosMap = map[pulid.PULID]struct{}{}
		for _, portfolio := range portfoliosExisting {
			portfoliosMap[portfolio.ID] = struct{}{}
		}

		for _, portfolioID := range input.AddPortfolioIDs {
			if _, ok := portfoliosMap[portfolioID]; !ok {
				return nil, fmt.Errorf(
					"connection.UpdateConnection(): invalid portfolio IDs to add provided: %w",
					resperror.NewGraphQLError("invalid portfolio(s) to add connection to", resperror.GQLBadUserInput),
				)
			}
		}

		for _, portfolioID := range input.RemovePortfolioIDs {
			if _, ok := portfoliosMap[portfolioID]; !ok {
				return nil, fmt.Errorf(
					"connection.UpdateConnection(): invalid portfolio IDs to remove provided: %w",
					resperror.NewGraphQLError(
						"invalid portfolio(s) to remove connection from", resperror.GQLBadUserInput,
					),
				)
			}
		}
	}

	// Update the connection.
	updatedConnection, err := dbClient.Connection.UpdateOneID(id).
		SetInput(input).
		SetAccountID(session.UserID).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"connection.CreateConnection(): could not update connection: %w",
			resperror.NewGraphQLError("could not update connection", resperror.GQLInternalServerError),
		)
	}

	return updatedConnection, nil
}
