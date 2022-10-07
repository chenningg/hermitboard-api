package auth

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"github.com/chenningg/hermitboard-api/ent/staffaccount"
	"github.com/chenningg/hermitboard-api/graph"
	"golang.org/x/crypto/bcrypt"
)

// LoginToAccount logs in an account.
// TODO: Make sure an already logged in account from the same IP and device can't login again.
func (authService AuthService) LoginToAccount(
	ctx context.Context, input LoginToAccountInput,
) (
	*LoginToAccountPayload, error,
) {
	// Check for missing values.
	if input.Username == "" || input.Password == "" {
		return nil, fmt.Errorf(
			"auth.LoginToAccount(): %w", graph.NewGraphQLError("missing username or password", graph.BadUserInput),
		)
	}

	if len(input.Password) < 8 {
		return nil, fmt.Errorf(
			"auth.LoginToAccount(): %w",
			graph.NewGraphQLError("password must be a minimum of 8 alphanumeric characters", graph.BadUserInput),
		)
	}

	session := GetSessionFromContext(ctx)

	if IsLoggedIn(session) {
		return nil, fmt.Errorf(
			"auth.LoginToAccount(): %w", graph.NewGraphQLError("already logged in", graph.Forbidden),
		)
	}

	// Username can be either email or nickname.
	dbClient := ent.FromContext(ctx)

	accounts, err := dbClient.Account.Query().Where(
		account.And(
			account.HasAuthTypeWith(authtype.ValueEQ(authtype.ValueLocal)), account.Or(
				account.NicknameEQ(input.Username), account.EmailEQ(input.Username),
			),
		),
	).All(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"auth.LoginToAccount(): %w: %v", graph.NewGraphQLError(
				"could not retrieve matching accounts from database for login", graph.InternalServerError,
			), err,
		)
	}
	// If no accounts found, return incorrect username or password.
	if len(accounts) == 0 {
		return nil, fmt.Errorf(
			"auth.LoginToAccount(): %w: could not find a matching account in the database",
			graph.NewGraphQLError("incorrect username or password", graph.Unauthenticated),
		)
	}

	for _, acc := range accounts {
		// Compare passwords.
		hashedPassword := acc.Password
		if hashedPassword == nil {
			continue
		}

		err = bcrypt.CompareHashAndPassword([]byte(*acc.Password), []byte(input.Password))
		if err == nil {
			// We have a match of username and password, create the user session.
			// Get authRoles of the user.
			authRoles, err := acc.QueryAuthRoles().All(ctx)
			if err != nil {
				return nil, fmt.Errorf(
					"auth.LoginToAccount(): could not retrieve auth roles of the logging in account: %w: %v",
					graph.NewGraphQLError("could not load account", graph.InternalServerError), err,
				)
			}

			authRoleValues := make([]authrole.Value, len(authRoles))
			for i, authRole := range authRoles {
				authRoleValues[i] = authRole.Value
			}

			// Create Session.
			newSession, err := authService.createSession(ctx, acc.ID, authRoleValues...)
			if err != nil {
				return nil, fmt.Errorf(
					"auth.LoginToAccount(): %w: %v",
					graph.NewGraphQLError("could not create a new session", graph.InternalServerError), err,
				)
			}

			return &LoginToAccountPayload{
				Account: acc,
				Session: newSession,
			}, nil
		}
	}

	// If we reach here, no accounts match the given username and password. Return an error.
	return nil, fmt.Errorf(
		"auth.LoginToAccount(): %w",
		graph.NewGraphQLError("invalid username or password", graph.Unauthenticated),
	)
}

func (authService AuthService) LoginToStaffAccount(
	ctx context.Context, input LoginToStaffAccountInput,
) (*LoginToStaffAccountPayload, error) {
	// Check for missing values.
	if input.Username == "" || input.Password == "" {
		return nil, fmt.Errorf(
			"auth.LoginToStaffAccount(): %w", graph.NewGraphQLError("missing username or password", graph.BadUserInput),
		)
	}

	if len(input.Password) < 8 {
		return nil, fmt.Errorf(
			"auth.LoginToStaffAccount(): %w",
			graph.NewGraphQLError("password must be a minimum of 8 alphanumeric characters", graph.BadUserInput),
		)
	}

	session := GetSessionFromContext(ctx)

	if IsLoggedIn(session) {
		return nil, fmt.Errorf(
			"auth.LoginToStaffAccount(): %w", graph.NewGraphQLError("already logged in", graph.Forbidden),
		)
	}

	// Username can be either email or nickname.
	dbClient := ent.FromContext(ctx)

	staffAccounts, err := dbClient.StaffAccount.Query().Where(
		staffaccount.And(
			staffaccount.HasAuthTypeWith(authtype.ValueEQ(authtype.ValueLocal)),
			staffaccount.Or(
				staffaccount.NicknameEQ(input.Username), staffaccount.EmailEQ(input.Password),
			),
		),
	).All(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"auth.LoginToStaffAccount(): %w: %v", graph.NewGraphQLError(
				"could not retrieve matching staff accounts from database for login", graph.InternalServerError,
			), err,
		)
	}
	// If no accounts found, return incorrect username or password.
	if len(staffAccounts) == 0 {
		return nil, fmt.Errorf(
			"auth.LoginToStaffAccount(): %w: could not find a matching staff account in the database",
			graph.NewGraphQLError(
				"incorrect username or password", graph.Unauthenticated,
			),
		)
	}

	for _, staffAcc := range staffAccounts {
		// Compare passwords.
		hashedPassword := staffAcc.Password
		if hashedPassword == nil {
			continue
		}

		err = bcrypt.CompareHashAndPassword([]byte(*staffAcc.Password), []byte(input.Password))
		if err == nil {
			// We have a match of username and password, create the user session.
			// Get authRoles of the user.
			authRoles, err := staffAcc.QueryAuthRoles().All(ctx)
			if err != nil {
				return nil, fmt.Errorf(
					"auth.LoginToStaffAccount(): %w: could not retrieve auth roles of the logging in staff account: %v",
					graph.NewGraphQLError("could not load staff account", graph.InternalServerError), err,
				)
			}

			authRoleValues := make([]authrole.Value, len(authRoles))
			for i, authRole := range authRoles {
				authRoleValues[i] = authRole.Value
			}

			// Create Session and store in Redis.
			newSession, err := authService.createSession(ctx, staffAcc.ID, authRoleValues...)
			if err != nil {
				return nil, fmt.Errorf(
					"auth.LoginToStaffAccount(): %w: %v",
					graph.NewGraphQLError("could not create a new session", graph.InternalServerError), err,
				)
			}

			return &LoginToStaffAccountPayload{
				StaffAccount: staffAcc,
				Session:      newSession,
			}, nil
		}
	}

	// If we reach here, no staff accounts match the given username and password. Return an error.
	return nil, fmt.Errorf(
		"auth.LoginToStaffAccount(): %w",
		graph.NewGraphQLError("invalid username or password", graph.Unauthenticated),
	)
}
