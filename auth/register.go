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
	ctx context.Context, input ent.CreateAccountInput,
) (
	*ent.Account, error,
) {
	// Check for empty values.
	if input.Nickname == "" || input.Email == "" || input.AuthTypeID.String() == "" {
		return nil, fmt.Errorf(
			"%w: auth.CreateAccount(): missing or invalid fields in account creation", ErrBadInput,
		)
	}

	dbClient := ent.FromContext(ctx)
	session := GetSessionFromContext(ctx)

	// Don't allow logged in non-staff accounts to create an account.
	if IsLoggedIn(session) && HasAuthRoles(session, NonStaffAuthRoles...) {
		return nil, fmt.Errorf(
			"%w: auth.CreateAccount(): non-staff account already logged in, cannot create account", ErrUnauthorized,
		)
	}

	// Get the authtype.
	authType, err := dbClient.AuthType.Query().Where(authtype.ID(input.AuthTypeID)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: auth.CreateAccount(): invalid auth type supplied", ErrBadInput,
		)
	}

	accCreator := dbClient.Account.Create().
		SetNickname(input.Nickname).
		SetEmail(input.Email).
		SetAuthType(authType)

	if authType.Value == authtype.ValueLocal {
		if input.Password == nil || *input.Password == "" {
			return nil, fmt.Errorf(
				"%w: auth.CreateAccount(): no password was provided for account creation with a local provider",
				ErrBadInput,
			)
		}
		// Check if password is minimum 8 characters.
		if len(*input.Password) < 8 {
			return nil, fmt.Errorf("%w: auth.CreateAccount(): password provided is too weak", ErrBadInput)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*input.Password), authService.config.BcryptCost)
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

	if len(input.AuthRoleIDs) > 0 {
		// Only allow staff to add auth roles when creating an account.
		if HasAuthRoles(
			session, authrole.ValueSupport, authrole.ValueAdmin, authrole.ValueSuperAdmin,
		) {
			// Get auth roles from database.
			authRoles, err := dbClient.AuthRole.Query().Where(authrole.IDIn(input.AuthRoleIDs...)).All(ctx)
			if err != nil {
				return nil, fmt.Errorf(
					"%w: auth.CreateAccount(): could not retrieve auth roles from database", ErrInternal,
				)
			}
			if len(authRoles) == 0 {
				return nil, fmt.Errorf(
					"%w: auth.CreateAccount(): auth roles specified for account creation are not valid", ErrBadInput,
				)
			}

			// Get auth role values.
			authRoleValues := make([]authrole.Value, len(authRoles))
			for i, ar := range authRoles {
				authRoleValues[i] = ar.Value
			}

			// Check that the auth roles to populate are within authorization. SuperAdmins can populate auth roles at will.
			if !HasAuthRoles(session, authrole.ValueSuperAdmin) && !IsHigherAuthority(
				session, authRoleValues,
			) {
				return nil, fmt.Errorf(
					"%w: auth.CreateAccount(): insufficient authority to add the auth roles specified", ErrUnauthorized,
				)
			}

			accCreator.AddAuthRoles(authRoles...)
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
	ctx context.Context, input ent.CreateStaffAccountInput,
) (*ent.StaffAccount, error) {
	// Check for empty values.
	if input.Nickname == "" || input.Email == "" || input.AuthTypeID.String() == "" {
		return nil, fmt.Errorf(
			"%w: auth.CreateStaffAccount(): missing or invalid fields in staff account creation", ErrBadInput,
		)
	}

	dbClient := ent.FromContext(ctx)
	session := GetSessionFromContext(ctx)

	// Only allow logged in admin/super admin staff accounts to create more staff accounts.
	if !IsLoggedIn(session) || !HasAuthRoles(session, authrole.ValueAdmin, authrole.ValueSuperAdmin) {
		return nil, fmt.Errorf(
			"%w: auth.CreateStaffAccount(): not logged-in or unauthorized to create new staff accounts",
			ErrUnauthorized,
		)
	}

	// Get the authtype.
	authType, err := dbClient.AuthType.Query().Where(authtype.ID(input.AuthTypeID)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: auth.CreateStaffAccount(): invalid auth type supplied", ErrBadInput,
		)
	}

	staffAccCreator := dbClient.StaffAccount.Create().
		SetNickname(input.Nickname).
		SetEmail(input.Email).
		SetAuthType(authType)

	if authType.Value == authtype.ValueLocal {
		if input.Password == nil || *input.Password == "" {
			return nil, fmt.Errorf(
				"%w: auth.CreateStaffAccount(): no password was provided for staff account creation with a local provider",
				ErrBadInput,
			)
		}
		// Check if password is minimum 8 characters.
		if len(*input.Password) < 8 {
			return nil, fmt.Errorf("%w: auth.CreateStaffAccount(): password provided is too weak", ErrBadInput)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*input.Password), authService.config.BcryptCost)
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

	if len(input.AuthRoleIDs) > 0 {
		// Get auth roles from database.
		authRoles, err := dbClient.AuthRole.Query().Where(authrole.IDIn(input.AuthRoleIDs...)).All(ctx)
		if err != nil {
			return nil, fmt.Errorf(
				"%w: auth.CreateStaffAccount(): could not retrieve auth roles from database", ErrInternal,
			)
		}
		if len(authRoles) == 0 {
			return nil, fmt.Errorf(
				"%w: auth.CreateStaffAccount(): auth roles specified for staff account creation are not valid",
				ErrBadInput,
			)
		}

		// Get auth role values.
		authRoleValues := make([]authrole.Value, len(authRoles))
		for i, ar := range authRoles {
			authRoleValues[i] = ar.Value
		}

		// Check that the auth roles to populate are within authorization. SuperAdmins can populate auth roles at will.
		if !HasAuthRoles(session, authrole.ValueSuperAdmin) && !IsHigherAuthority(
			session, authRoleValues,
		) {
			return nil, fmt.Errorf(
				"%w: auth.CreateStaffAccount(): insufficient authority to add the auth roles specified",
				ErrUnauthorized,
			)
		}

		staffAccCreator.AddAuthRoles(authRoles...)
	} else {
		// Otherwise, by default, a newly created staff account is on the Support role.
		authRoleSupport, err := dbClient.AuthRole.Query().Where(authrole.ValueEQ(authrole.ValueSupport)).Only(ctx)
		if err != nil {
			return nil, fmt.Errorf(
				"%w: auth.CreateStaffAccount(): could not add support auth role to staff account: %v", ErrInternal, err,
			)
		}

		staffAccCreator.AddAuthRoles(authRoleSupport)
	}

	// Save the created staff account to the database.
	staffAcc, err := staffAccCreator.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: auth.CreateStaffAccount(): could not create staff account: %v", ErrInternal, err)
	}

	return staffAcc, nil
}
