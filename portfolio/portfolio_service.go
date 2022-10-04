package portfolio

import (
	"context"
	"errors"
	"fmt"

	"github.com/chenningg/hermitboard-api/auth"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/connection"
	"github.com/chenningg/hermitboard-api/ent/portfolio"
	"github.com/chenningg/hermitboard-api/pulid"
	"github.com/go-logr/logr"
)

type PortfolioServicer interface {
	CreatePortfolio(
		ctx context.Context, input ent.CreatePortfolioInput,
	) (*ent.Portfolio, error)
	UpdatePortfolio(ctx context.Context, id pulid.PULID, input ent.UpdatePortfolioInput) (*ent.Portfolio, error)
}

type PortfolioService struct {
	logger logr.Logger
}

func NewPortfolioService(logger logr.Logger) *PortfolioService {
	var portfolioService = new(PortfolioService)

	// Initialize logger.
	portfolioService.logger = logger
	portfolioService.logger.V(2).Info("NewPortfolioService(): initializing portfolio service")

	return portfolioService
}

func (portfolioService PortfolioService) CreatePortfolio(
	ctx context.Context, input ent.CreatePortfolioInput,
) (*ent.Portfolio, error) {
	dbClient := ent.FromContext(ctx)
	session := auth.GetSessionFromContext(ctx)

	// Check if logged in.
	if !auth.IsLoggedIn(session) {
		return nil, fmt.Errorf(
			"%w: portfolio.CreatePortfolio(): not logged in, cannot create portfolio", ErrUnauthorized,
		)
	}

	// Check for empty fields.
	if input.Name == "" {
		return nil, fmt.Errorf("%w: portfolio.CreatePortfolio(): portfolio name cannot be empty", ErrBadInput)
	}

	// Only user accounts and non-demo accounts can create portfolios.
	if !auth.HasAuthRoles(session, authrole.ValueDemo) || !auth.HasAuthRoles(session, auth.NonStaffAuthRoles...) {
		return nil, fmt.Errorf(
			"%w: portfolio.CreatePortfolio(): only non-staff and non-demo accounts can create portfolios",
			ErrUnauthorized,
		)
	}

	// If there are connections attached to the portfolio creation, try to retrieve them.
	if len(input.ConnectionIDs) > 0 {
		connections, err := dbClient.Connection.Query().Where(connection.AccountIDEQ(session.UserID)).All(ctx)
		if err != nil {
			return nil, fmt.Errorf(
				"%w: portfolio.CreatePortfolio(): could not retrieve account connections", ErrInternal,
			)
		}

		// Create hash map to facilitate O(1) checking.
		connectionsMap := map[pulid.PULID]struct{}{}
		for _, connectionAvailable := range connections {
			connectionsMap[connectionAvailable.ID] = struct{}{}
		}

		for _, connectionToAddToPortfolio := range input.ConnectionIDs {
			// If we can't find the connection in the account, we cannot create this portfolio.
			if _, ok := connectionsMap[connectionToAddToPortfolio]; !ok {
				return nil, fmt.Errorf(
					"%w: portfolio.CreatePortfolio(): a connection supplied does not exist on this account",
					ErrBadInput,
				)
			}
		}
	}

	// Add the account ID to the portfolio.
	newPortfolio, err := dbClient.Portfolio.Create().
		SetInput(input).
		SetAccountID(session.UserID).
		Save(ctx)
	if err != nil {
		// TODO: May also have constraint error such as if portfolio name conflicts with an existing portfolio.
		return nil, fmt.Errorf(
			"%w: portfolio.CreatePortfolio(): could not create a new portfolio", ErrInternal,
		)
	}

	return newPortfolio, nil
}

func (portfolioService PortfolioService) UpdatePortfolio(
	ctx context.Context, id pulid.PULID, input ent.UpdatePortfolioInput,
) (*ent.Portfolio, error) {
	dbClient := ent.FromContext(ctx)
	session := auth.GetSessionFromContext(ctx)

	// Check if logged in.
	if !auth.IsLoggedIn(session) {
		return nil, fmt.Errorf(
			"%w: portfolio.UpdatePortfolio(): not logged in, cannot update portfolio", ErrUnauthorized,
		)
	}

	// Only user accounts and non-demo accounts can update portfolios.
	if !auth.HasAuthRoles(session, authrole.ValueDemo) || !auth.HasAuthRoles(session, auth.NonStaffAuthRoles...) {
		return nil, fmt.Errorf(
			"%w: portfolio.UpdatePortfolio(): only non-staff and non-demo accounts can update portfolios",
			ErrUnauthorized,
		)
	}

	// Check that portfolio being updated is owned by this session's user.
	_, err := dbClient.Portfolio.Query().Where(
		portfolio.And(
			portfolio.IDEQ(id), portfolio.AccountIDEQ(session.UserID),
		),
	).Only(ctx)
	if err != nil {
		if errors.Is(err, &ent.NotFoundError{}) {
			return nil, fmt.Errorf(
				"%w: portfolio.UpdatePortfolio(): could not find a portfolio with a matching ID under this account",
				ErrNotFound,
			)
		}
		return nil, fmt.Errorf(
			"%w: portfolio.UpdatePortfolio(): could not query account's portfolios",
			ErrInternal,
		)
	}

	if len(input.AddConnectionIDs) > 0 || len(input.RemoveConnectionIDs) > 0 {
		// Check that the connection IDs belong to the user.
		connectionsExisting, err := dbClient.Connection.Query().Where(connection.AccountIDEQ(session.UserID)).All(ctx)
		if err != nil {
			if errors.Is(err, &ent.NotFoundError{}) {
				return nil, fmt.Errorf(
					"%w: portfolio.UpdatePortfolio(): no connections found on this account",
					ErrNotFound,
				)
			}
			return nil, fmt.Errorf(
				"%w: portfolio.UpdatePortfolio(): could not query account's connections",
				ErrInternal,
			)
		}

		var connectionsMap = map[pulid.PULID]struct{}{}
		for _, connection := range connectionsExisting {
			connectionsMap[connection.ID] = struct{}{}
		}

		for _, connectionToAdd := range input.AddConnectionIDs {
			if _, ok := connectionsMap[connectionToAdd]; !ok {
				return nil, fmt.Errorf(
					"%w: portfolio.UpdatePortfolio(): cannot add a connection that doesn't exist on this account",
					ErrBadInput,
				)
			}
		}

		for _, connectionToRemove := range input.RemoveConnectionIDs {
			if _, ok := connectionsMap[connectionToRemove]; !ok {
				return nil, fmt.Errorf(
					"%w: portfolio.UpdatePortfolio(): cannot remove a connection that doesn't exist on this account",
					ErrBadInput,
				)
			}
		}
	}

	// Update the portfolio.
	updatedPortfolio, err := dbClient.Portfolio.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: portfolio.UpdatePortfolio(): could not update portfolio",
			ErrInternal,
		)
	}

	return updatedPortfolio, nil
}
