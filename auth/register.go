package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"github.com/chenningg/hermitboard-api/graph"
	"golang.org/x/crypto/bcrypt"
)

func (authService AuthService) CreateAccount(
	ctx context.Context, input ent.CreateAccountInput,
) (
	*CreateAccountPayload, error,
) {
	// Check for empty values.
	if input.Nickname == "" || input.Email == "" || input.AuthTypeID.String() == "" {
		return nil, fmt.Errorf(
			"auth.CreateAccount(): %w",
			graph.NewGraphQLError("missing or invalid fields in account creation", graph.BadUserInput),
		)
	}

	dbClient := ent.FromContext(ctx)
	session := GetSessionFromContext(ctx)

	// Don't allow logged in non-staff accounts to create an account.
	if IsLoggedIn(session) && HasAuthRoles(session, NonStaffAuthRoles...) {
		return nil, fmt.Errorf(
			"auth.CreateAccount(): %w: a non-staff account cannot create an account",
			graph.NewGraphQLError("already logged in", graph.Forbidden),
		)
	}

	// Get the authtype.
	authType, err := dbClient.AuthType.Query().Where(authtype.ID(input.AuthTypeID)).Only(ctx)
	if err != nil {
		if errors.Is(err, &ent.NotFoundError{}) {
			return nil, fmt.Errorf(
				"auth.CreateAccount(): %w: %v", graph.NewGraphQLError("invalid auth type supplied", graph.BadUserInput),
				err,
			)
		}
		return nil, fmt.Errorf(
			"auth.CreateAccount(): %w: %v",
			graph.NewGraphQLError("could not retrieve auth type", graph.InternalServerError), err,
		)
	}

	accCreator := dbClient.Account.Create().
		SetNickname(input.Nickname).
		SetEmail(input.Email).
		SetAuthType(authType)

	if authType.Value == authtype.ValueLocal {
		if input.Password == nil || *input.Password == "" {
			return nil, fmt.Errorf(
				"auth.CreateAccount(): %w: a local provider requires a password",
				graph.NewGraphQLError("no password was provided for account creation", graph.BadUserInput),
			)
		}
		// Check if password is minimum 8 characters.
		if len(*input.Password) < 8 {
			return nil, fmt.Errorf(
				"auth.CreateAccount(): %w", graph.NewGraphQLError(
					"password must have a minimum of 8 alphanumeric characters or symbols", graph.BadUserInput,
				),
			)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*input.Password), authService.config.BcryptCost)
		if err != nil {
			return nil, fmt.Errorf(
				"auth.CreateAccount(): %w: %v",
				graph.NewGraphQLError("unable to store provided password", graph.InternalServerError), err,
			)
		}
		hashedPasswordStr := string(hashedPassword)
		accCreator.SetNillablePassword(&hashedPasswordStr)

		// Set password updated at time.
		currTime := time.Now()
		accCreator.SetNillablePasswordUpdatedAt(&currTime)
	}

	if len(input.AuthRoleIDs) > 0 {
		// Only allow staff to add auth roles when creating an account.
		if HasAuthRoles(
			session, authrole.ValueSupport, authrole.ValueAdmin, authrole.ValueSuperAdmin,
		) {
			// Get auth roles from database.
			authRoles, err := dbClient.AuthRole.Query().Where(authrole.IDIn(input.AuthRoleIDs...)).All(ctx)
			if err != nil {
				return nil, fmt.Errorf(
					"auth.CreateAccount(): %w: %v",
					graph.NewGraphQLError("could not create account", graph.InternalServerError), err,
				)
			}
			if len(authRoles) == 0 {
				return nil, fmt.Errorf(
					"auth.CreateAccount(): %w: no matching auth roles found in database", graph.NewGraphQLError(
						"auth roles specified for account creation are not valid", graph.BadUserInput,
					),
				)
			}

			// Get auth role values.
			authRoleValues := make([]authrole.Value, len(authRoles))
			for i, ar := range authRoles {
				authRoleValues[i] = ar.Value
			}

			// Check that the auth roles to populate are within authorization. SuperAdmins can populate auth roles at will.
			if !HasAuthRoles(session, authrole.ValueSuperAdmin) && !IsHigherAuthority(
				session, authRoleValues...,
			) {
				return nil, fmt.Errorf(
					"auth.CreateAccount(): %w",
					graph.NewGraphQLError("insufficient authority to add the auth roles specified", graph.Forbidden),
				)
			}

			accCreator.AddAuthRoles(authRoles...)
		} else {
			// Otherwise this session user has no authority to add auth roles.
			return nil, fmt.Errorf(
				"auth.CreateAccount(): %w",
				graph.NewGraphQLError("unauthorized to add auth roles to an account", graph.Forbidden),
			)
		}
	} else {
		// Otherwise, by default, a newly created account is on the Free role.
		authRoleFree, err := dbClient.AuthRole.Query().Where(authrole.ValueEQ(authrole.ValueFree)).Only(ctx)
		if err != nil {
			return nil, fmt.Errorf(
				"auth.CreateAccount(): %w: could not retrieve free auth role from database: %v",
				graph.NewGraphQLError("could not create account", graph.InternalServerError), err,
			)
		}

		accCreator.AddAuthRoles(authRoleFree)
	}

	// Save the created account to the database.
	acc, err := accCreator.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"auth.CreateAccount(): %w: %v",
			graph.NewGraphQLError("could not create account", graph.InternalServerError), err,
		)
	}

	// Make a new session for login.
	newAuthRoles := make([]authrole.Value, len(acc.Edges.AuthRoles))
	for i, newAuthRole := range acc.Edges.AuthRoles {
		newAuthRoles[i] = newAuthRole.Value
	}

	newSession, err := authService.createSession(ctx, acc.ID, newAuthRoles...)
	if err != nil {
		return nil, fmt.Errorf(
			"auth.CreateAccount(): %w: %v",
			graph.NewGraphQLError("could not create new session", graph.InternalServerError), err,
		)
	}

	return &CreateAccountPayload{acc, newSession}, nil
}

func (authService AuthService) CreateStaffAccount(
	ctx context.Context, input ent.CreateStaffAccountInput,
) (*CreateStaffAccountPayload, error) {
	// Check for empty values.
	if input.Nickname == "" || input.Email == "" || input.AuthTypeID.String() == "" || len(input.AuthRoleIDs) == 0 {
		return nil, fmt.Errorf(
			"auth.CreateStaffAccount(): %w",
			graph.NewGraphQLError("missing or invalid fields in staff account creation", graph.BadUserInput),
		)
	}

	dbClient := ent.FromContext(ctx)
	session := GetSessionFromContext(ctx)

	// Only allow logged in admin/super admin staff accounts to create more staff accounts.
	if !IsLoggedIn(session) {
		return nil, fmt.Errorf(
			"auth.CreateStaffAccount(): %w",
			graph.NewGraphQLError("not logged in", graph.Unauthenticated),
		)
	}

	if !HasAuthRoles(session, authrole.ValueAdmin, authrole.ValueSuperAdmin) {
		return nil, fmt.Errorf(
			"auth.CreateStaffAccount(): %w",
			graph.NewGraphQLError("unauthorized to create staff accounts", graph.Forbidden),
		)
	}

	// Get the authtype.
	authType, err := dbClient.AuthType.Query().Where(authtype.ID(input.AuthTypeID)).Only(ctx)
	if err != nil {
		if errors.Is(err, &ent.NotFoundError{}) {
			return nil, fmt.Errorf(
				"auth.CreateStaffAccount(): %w: %v",
				graph.NewGraphQLError("invalid auth type supplied", graph.BadUserInput),
				err,
			)
		}
		return nil, fmt.Errorf(
			"auth.CreateStaffAccount(): %w: %v",
			graph.NewGraphQLError("could not retrieve auth type", graph.InternalServerError), err,
		)
	}

	staffAccCreator := dbClient.StaffAccount.Create().
		SetNickname(input.Nickname).
		SetEmail(input.Email).
		SetAuthType(authType)

	if authType.Value == authtype.ValueLocal {
		if input.Password == nil || *input.Password == "" {
			return nil, fmt.Errorf(
				"auth.CreateStaffAccount(): %w: a local provider requires a password",
				graph.NewGraphQLError("no password was provided for account creation", graph.BadUserInput),
			)
		}
		// Check if password is minimum 8 characters.
		if len(*input.Password) < 8 {
			return nil, fmt.Errorf(
				"auth.CreateStaffAccount(): %w", graph.NewGraphQLError(
					"password must have a minimum of 8 alphanumeric characters or symbols", graph.BadUserInput,
				),
			)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*input.Password), authService.config.BcryptCost)
		if err != nil {
			return nil, fmt.Errorf(
				"auth.CreateStaffAccount(): %w: %v",
				graph.NewGraphQLError("unable to store provided password", graph.InternalServerError), err,
			)
		}
		hashedPasswordStr := string(hashedPassword)
		staffAccCreator.SetNillablePassword(&hashedPasswordStr)

		// Set password updated at time.
		currTime := time.Now()
		staffAccCreator.SetNillablePasswordUpdatedAt(&currTime)
	}

	if len(input.AuthRoleIDs) > 0 {
		// Get auth roles from database.
		authRoles, err := dbClient.AuthRole.Query().Where(authrole.IDIn(input.AuthRoleIDs...)).All(ctx)
		if err != nil {
			return nil, fmt.Errorf(
				"auth.CreateAccount(): %w: %v",
				graph.NewGraphQLError("could not create staff account", graph.InternalServerError), err,
			)
		}
		if len(authRoles) == 0 {
			return nil, fmt.Errorf(
				"auth.CreateAccount(): %w: no matching auth roles found in database", graph.NewGraphQLError(
					"auth roles specified for staff account creation are not valid", graph.BadUserInput,
				),
			)
		}

		// Get auth role values.
		authRoleValues := make([]authrole.Value, len(authRoles))
		for i, ar := range authRoles {
			authRoleValues[i] = ar.Value
		}

		// Check that the auth roles to populate are within authorization. SuperAdmins can populate auth roles at will.
		if !HasAuthRoles(session, authrole.ValueSuperAdmin) && !IsHigherAuthority(
			session, authRoleValues...,
		) {
			return nil, fmt.Errorf(
				"auth.CreateStaffAccount(): %w",
				graph.NewGraphQLError("insufficient authority to add the auth roles specified", graph.Forbidden),
			)
		}

		staffAccCreator.AddAuthRoles(authRoles...)
	}

	// Save the created staff account to the database.
	staffAcc, err := staffAccCreator.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"auth.CreateStaffAccount(): %w: %v",
			graph.NewGraphQLError("could not create staff account", graph.InternalServerError), err,
		)
	}

	// Make a new session for login.
	newAuthRoles := make([]authrole.Value, len(staffAcc.Edges.AuthRoles))
	for i, newAuthRole := range staffAcc.Edges.AuthRoles {
		newAuthRoles[i] = newAuthRole.Value
	}

	newSession, err := authService.createSession(ctx, staffAcc.ID, newAuthRoles...)
	if err != nil {
		return nil, fmt.Errorf(
			"auth.CreateStaffAccount(): %w: %v",
			graph.NewGraphQLError("could not create new session", graph.InternalServerError), err,
		)
	}

	return &CreateStaffAccountPayload{staffAcc, newSession}, nil
}
