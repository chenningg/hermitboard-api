package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"golang.org/x/crypto/bcrypt"
)

func (authService AuthService) CreateAccount(
	ctx context.Context, nickname string, email string, password string, provider authtype.Value,
	authRoles []authrole.Value,
) (
	*ent.Account, error,
) {
	// Check for empty values.
	if nickname == "" || email == "" || provider == "" || authtype.ValueValidator(provider) != nil {
		return nil, fmt.Errorf(
			"%w: auth.CreateAccount(): missing or invalid fields in account creation", ErrBadInput,
		)
	}

	dbClient := ent.FromContext(ctx)
	session := GetSessionFromContext(ctx)

	// Don't allow logged in non-staff accounts to create an account.
	if IsLoggedIn(session) && HasAuthRoles(session, NonStaffAuthRoles) {
		return nil, fmt.Errorf(
			"%w: auth.CreateAccount(): non-staff account already logged in, cannot create account", ErrAlreadyLoggedIn,
		)
	}

	// If provider is local, a password must be provided to create an account.
	if provider == authtype.ValueLocal && password == "" {
		return nil, fmt.Errorf(
			"%w: auth.CreateAccount(): no password was provided for account creation with a local provider",
			ErrBadInput,
		)
	}

	// Get the auth type to populate.
	authType, err := dbClient.AuthType.Query().Where(authtype.ValueEQ(provider)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: auth.CreateAccount(): invalid auth type provided: %v", ErrBadInput, err)
	}

	accCreator := dbClient.Account.Create().
		SetNickname(nickname).
		SetEmail(email).
		SetAuthType(authType)

	if provider == authtype.ValueLocal {
		// Check if password is minimum 8 characters.
		if len(password) < 8 {
			return nil, fmt.Errorf("%w: auth.CreateAccount(): password provided is too weak", ErrBadInput)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), authService.config.BcryptCost)
		if err != nil {
			return nil, fmt.Errorf(
				"%w: auth.CreateAccount(): unable to hash provided password: %v", ErrInternal, err,
			)
		}
		hashedPasswordStr := string(hashedPassword)
		accCreator.SetNillablePassword(&hashedPasswordStr)

		// Set password updated at time.
		currTime := time.Now()
		accCreator.SetNillablePasswordUpdatedAt(&currTime)
	}

	if len(authRoles) > 0 {
		// Only allow staff to add auth roles when creating an account.
		if HasAuthRoles(
			session, []authrole.Value{authrole.ValueSupport, authrole.ValueAdmin, authrole.ValueSuperAdmin},
		) {
			// Check that the auth roles to populate are within authorization. SuperAdmins can populate auth roles at will.
			if !HasAuthRoles(session, []authrole.Value{authrole.ValueSuperAdmin}) && !IsHigherAuthority(
				session, authRoles,
			) {
				return nil, fmt.Errorf(
					"%w: auth.CreateAccount(): insufficient authority to add the auth roles specified", ErrUnauthorized,
				)
			}

			// Get the auth roles to populate.
			authRoleValues, err := dbClient.AuthRole.Query().Where(authrole.ValueIn(authRoles...)).All(ctx)
			if err != nil {
				return nil, fmt.Errorf(
					"%w: auth.CreateAccount(): invalid auth roles provided: %v", ErrBadInput, err,
				)
			}

			accCreator.AddAuthRoles(authRoleValues...)
		} else {
			// Otherwise this session user has no authority to add auth roles.
			return nil, fmt.Errorf(
				"%w: auth.CreateAccount(): unauthorized to add auth roles to account", ErrUnauthorized,
			)
		}
	} else {
		// Otherwise, by default, a newly created account is on the Free role.
		authRoleFree, err := dbClient.AuthRole.Query().Where(authrole.ValueEQ(authrole.ValueFree)).Only(ctx)
		if err != nil {
			return nil, fmt.Errorf(
				"%w: auth.CreateAccount(): could not add free auth role to account: %v", ErrInternal, err,
			)
		}

		accCreator.AddAuthRoles(authRoleFree)
	}

	// Save the created account to the database.
	acc, err := accCreator.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: auth.CreateAccount(): could not create account: %v", ErrInternal, err)
	}

	return acc, nil
}

func (authService AuthService) CreateStaffAccount(
	ctx context.Context, nickname string, email string, password string, provider authtype.Value,
	authRoles []authrole.Value,
) (*ent.StaffAccount, error) {
	// Check for empty values.
	if nickname == "" || email == "" || provider == "" || authtype.ValueValidator(provider) != nil {
		return nil, fmt.Errorf(
			"%w: auth.CreateStaffAccount(): missing or invalid fields in staff account creation", ErrBadInput,
		)
	}

	dbClient := ent.FromContext(ctx)
	session := GetSessionFromContext(ctx)

	// Only allow logged in admins and super admins to create new staff accounts.
	if !IsLoggedIn(session) || !HasAuthRoles(session, []authrole.Value{authrole.ValueAdmin, authrole.ValueSuperAdmin}) {
		return nil, fmt.Errorf(
			"%w: auth.CreateStaffAccount(): not authorized to create a staff account",
			ErrUnauthorized,
		)
	}

	// If provider is local, a password must be provided to create an account.
	if provider == authtype.ValueLocal && password == "" {
		return nil, fmt.Errorf(
			"%w: auth.CreateStaffAccount(): no password was provided for staff account creation with a local provider",
			ErrBadInput,
		)
	}

	// Get the auth type to populate.
	authType, err := dbClient.AuthType.Query().Where(authtype.ValueEQ(provider)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: auth.CreateStaffAccount(): invalid auth type provided: %v", ErrBadInput, err)
	}

	staffAccCreator := dbClient.StaffAccount.Create().
		SetNickname(nickname).
		SetEmail(email).
		SetAuthType(authType)

	if provider == authtype.ValueLocal {
		// Check if password is minimum 8 characters.
		if len(password) < 8 {
			return nil, fmt.Errorf("%w: auth.CreateStaffAccount(): password provided is too weak", ErrBadInput)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), authService.config.BcryptCost)
		if err != nil {
			return nil, fmt.Errorf(
				"%w: auth.CreateStaffAccount(): unable to hash provided password: %v", ErrInternal, err,
			)
		}
		hashedPasswordStr := string(hashedPassword)
		staffAccCreator.SetNillablePassword(&hashedPasswordStr)

		// Set password updated at time.
		currTime := time.Now()
		staffAccCreator.SetNillablePasswordUpdatedAt(&currTime)
	}

	if len(authRoles) > 0 {
		// Check that the auth roles to populate are within authorization. SuperAdmins can populate auth roles at will.
		if !HasAuthRoles(session, []authrole.Value{authrole.ValueSuperAdmin}) && !IsHigherAuthority(
			session, authRoles,
		) {
			return nil, fmt.Errorf(
				"%w: auth.CreateStaffAccount(): insufficient authority to add the auth roles specified",
				ErrUnauthorized,
			)
		}

		// Get the auth roles to populate.
		authRoleValues, err := dbClient.AuthRole.Query().Where(authrole.ValueIn(authRoles...)).All(ctx)
		if err != nil {
			return nil, fmt.Errorf(
				"%w: auth.CreateStaffAccount(): invalid auth roles provided: %v", ErrBadInput, err,
			)
		}

		staffAccCreator.AddAuthRoles(authRoleValues...)
	}

	// Save the created staff account to the database.
	staffAcc, err := staffAccCreator.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: auth.CreateStaffAccount(): could not create staff account: %v", ErrInternal, err)
	}

	return staffAcc, nil
}
