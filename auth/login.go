package auth

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"github.com/chenningg/hermitboard-api/ent/staffaccount"
	"golang.org/x/crypto/bcrypt"
)

func (authService AuthService) LoginToAccount(
	ctx context.Context, username string, password string,
) (
	*ent.Account, SessionID, error,
) {
	// Check for missing values.
	if username == "" || password == "" {
		return nil, "", fmt.Errorf("%w: auth.LoginToAccount(): missing username or password", ErrBadInput)
	}

	// Username can be either email or nickname.
	dbClient := ent.FromContext(ctx)

	accounts, err := dbClient.Account.Query().Where(
		account.Or(
			account.NicknameEQ(username), account.EmailEQ(username),
			account.HasAuthTypeWith(authtype.ValueEQ(authtype.ValueLocal)),
		),
	).All(ctx)
	if err != nil {
		return nil, "", fmt.Errorf(
			"%w: auth.LoginToAccount(): could not retrieve matching accounts from database for login: %v", ErrInternal,
			err,
		)
	}

	// If no accounts found, return incorrect username or password.
	if len(accounts) == 0 {
		return nil, "", ErrIncorrectUsernameOrPassword
	}

	for _, acc := range accounts {
		// Compare passwords.
		hashedPassword := acc.Password
		if hashedPassword == nil {
			continue
		}

		err = bcrypt.CompareHashAndPassword([]byte(*acc.Password), []byte(password))
		if err == nil {
			// We have a match of username and password, create the user session.
			sessionID, err := NewSessionID()
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToAccount(): failed to create session ID: %v", ErrInternal, err,
				)
			}

			// Get expiration time for a session.
			expiration, err := SecondsToDuration(authService.Config().SessionTimeout)
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToAccount(): could not convert session expiry time: %v", ErrInternal, err,
				)
			}

			// Get authRoles of the user.
			authRoles, err := acc.QueryAuthRoles().All(ctx)
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToAccount(): could not retrieve auth roles of the account: %v", ErrInternal, err,
				)
			}

			var authRoleValues []authrole.Value
			for _, authRole := range authRoles {
				authRoleValues = append(authRoleValues, authRole.Value)
			}

			// Store user ID in Redis.
			_, err = authService.redisService.Client().Set(
				ctx,
				fmt.Sprintf("%s:%s:userId", SessionRedisKey, sessionID),
				acc.ID.String(), expiration,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToAccount(): could not store user ID in Redis: %v", ErrInternal, err,
				)
			}

			// Store authRoles in Redis.
			_, err = authService.redisService.Client().RPush(
				ctx,
				fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID),
				authRoleValues,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToAccount(): could not store auth roles in Redis: %v", ErrInternal, err,
				)
			}

			// Set expiry on the authRoles key as well.
			_, err = authService.redisService.Client().Expire(
				ctx, fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID), expiration,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToAccount(): could not set expiry for auth role key in Redis: %v", ErrInternal, err,
				)
			}

			return acc, sessionID, nil
		}
	}

	// If we reach here, no accounts match the given username and password. Return an error.
	return nil, "", ErrIncorrectUsernameOrPassword
}

func (authService AuthService) LoginToStaffAccount(
	ctx context.Context, username string, password string,
) (*ent.StaffAccount, SessionID, error) {
	// Check for missing values.
	if username == "" || password == "" {
		return nil, "", fmt.Errorf("%w: auth.LoginToStaffAccount(): missing username or password", ErrBadInput)
	}

	// Username can be either email or nickname.
	dbClient := ent.FromContext(ctx)

	staffAccounts, err := dbClient.StaffAccount.Query().Where(
		staffaccount.Or(
			staffaccount.NicknameEQ(username), staffaccount.EmailEQ(username),
			staffaccount.HasAuthTypeWith(authtype.ValueEQ(authtype.ValueLocal)),
		),
	).All(ctx)
	if err != nil {
		return nil, "", fmt.Errorf(
			"%w: auth.LoginToStaffAccount(): could not retrieve matching staff accounts from database for login: %v",
			ErrInternal,
			err,
		)
	}

	// If no accounts found, return incorrect username or password.
	if len(staffAccounts) == 0 {
		return nil, "", ErrIncorrectUsernameOrPassword
	}

	for _, staffAcc := range staffAccounts {
		// Compare passwords.
		hashedPassword := staffAcc.Password
		if hashedPassword == nil {
			continue
		}

		err = bcrypt.CompareHashAndPassword([]byte(*staffAcc.Password), []byte(password))
		if err == nil {
			// We have a match of username and password, create the user session.
			sessionID, err := NewSessionID()
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToStaffAccount(): failed to create session ID: %v", ErrInternal, err,
				)
			}

			// Get expiration time for a session.
			expiration, err := SecondsToDuration(authService.Config().SessionTimeout)
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToStaffAccount(): could not convert session expiry time: %v", ErrInternal, err,
				)
			}

			// Get authRoles of the user.
			authRoles, err := staffAcc.QueryAuthRoles().All(ctx)
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToStaffAccount(): could not retrieve auth roles of the staff account: %v",
					ErrInternal,
					err,
				)
			}

			var authRoleValues []authrole.Value
			for _, authRole := range authRoles {
				authRoleValues = append(authRoleValues, authRole.Value)
			}

			// Store user ID in Redis.
			_, err = authService.redisService.Client().Set(
				ctx,
				fmt.Sprintf("%s:%s:userId", SessionRedisKey, sessionID),
				staffAcc.ID.String(), expiration,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToStaffAccount(): could not store user ID in Redis: %v", ErrInternal, err,
				)
			}

			// Store authRoles in Redis.
			_, err = authService.redisService.Client().RPush(
				ctx,
				fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID),
				authRoleValues,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToStaffAccount(): could not store auth roles in Redis: %v", ErrInternal, err,
				)
			}

			// Set expiry on the authRoles key as well.
			_, err = authService.redisService.Client().Expire(
				ctx, fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID), expiration,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf(
					"%w: auth.LoginToStaffAccount(): could not set expiry for auth role key in Redis: %v", ErrInternal,
					err,
				)
			}

			return staffAcc, sessionID, nil
		}
	}

	// If we reach here, no accounts match the given username and password. Return an error.
	return nil, "", ErrIncorrectUsernameOrPassword
}
