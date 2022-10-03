package auth

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"github.com/chenningg/hermitboard-api/ent/staffaccount"
	"golang.org/x/crypto/bcrypt"
)

func (authService AuthService) LoginToAccount(
	ctx context.Context, input LoginToAccountInput,
) (
	SessionID, error,
) {
	// Check for missing values.
	if input.Username == "" || input.Password == "" {
		return "", fmt.Errorf("%w: auth.LoginToAccount(): missing username or password", ErrBadInput)
	}

	if len(input.Password) < 8 {
		return "", fmt.Errorf(
			"%w: auth.LoginToAccount(): password must be a minimum of 8 alphanumeric characters", ErrBadInput,
		)
	}

	session := GetSessionFromContext(ctx)

	if IsLoggedIn(session) {
		return "", fmt.Errorf(
			"%w: auth.LoginToAccount(): already logged in", ErrBadInput,
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
		return "", fmt.Errorf(
			"%w: auth.LoginToAccount(): could not retrieve matching accounts from database for login: %v", ErrInternal,
			err,
		)
	}
	// If no accounts found, return incorrect username or password.
	if len(accounts) == 0 {
		return "", ErrIncorrectUsernameOrPassword
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
			sessionID, err := NewSessionID()
			if err != nil {
				return "", fmt.Errorf(
					"%w: auth.LoginToAccount(): failed to create session ID: %v", ErrInternal, err,
				)
			}

			// Get authRoles of the user.
			authRoles, err := acc.QueryAuthRoles().All(ctx)
			if err != nil {
				return "", fmt.Errorf(
					"%w: auth.LoginToAccount(): could not retrieve auth roles of the account: %v", ErrInternal, err,
				)
			}

			authRoleValues := make([]string, len(authRoles))
			for i, authRole := range authRoles {
				authRoleValues[i] = authRole.Value.String()
			}

			// Store user ID in Redis.
			_, err = authService.redisService.Client().Set(
				ctx,
				fmt.Sprintf("%s:%s:userId", SessionRedisKey, sessionID),
				acc.ID.String(), authService.Config().SessionTimeout,
			).Result()
			if err != nil {
				return "", fmt.Errorf(
					"%w: auth.LoginToAccount(): could not store user ID in redis: %v", ErrInternal, err,
				)
			}

			// Store authRoles in Redis.
			_, err = authService.redisService.Client().SAdd(
				ctx,
				fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID),
				authRoleValues,
			).Result()
			if err != nil {
				return "", fmt.Errorf(
					"%w: auth.LoginToAccount(): could not store auth roles in redis: %v", ErrInternal, err,
				)
			}

			// Set expiry on the authRoles key as well.
			_, err = authService.redisService.Client().Expire(
				ctx, fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID), authService.Config().SessionTimeout,
			).Result()
			if err != nil {
				return "", fmt.Errorf(
					"%w: auth.LoginToAccount(): could not set expiry for auth role key in Redis: %v", ErrInternal, err,
				)
			}

			return sessionID, nil
		}
	}

	// If we reach here, no accounts match the given username and password. Return an error.
	return "", ErrIncorrectUsernameOrPassword
}

func (authService AuthService) LoginToStaffAccount(
	ctx context.Context, input LoginToStaffAccountInput,
) (SessionID, error) {
	// Check for missing values.
	if input.Username == "" || input.Password == "" {
		return "", fmt.Errorf("%w: auth.LoginToStaffAccount(): missing username or password", ErrBadInput)
	}

	if len(input.Password) < 8 {
		return "", fmt.Errorf(
			"%w: auth.LoginToStaffAccount(): password must be a minimum of 8 alphanumeric characters", ErrBadInput,
		)
	}

	session := GetSessionFromContext(ctx)

	if IsLoggedIn(session) {
		return "", fmt.Errorf(
			"%w: auth.LoginToStaffAccount(): already logged in", ErrBadInput,
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
		return "", fmt.Errorf(
			"%w: auth.LoginToStaffAccount(): could not retrieve matching staff accounts from database for login: %v",
			ErrInternal,
			err,
		)
	}
	// If no accounts found, return incorrect username or password.
	if len(staffAccounts) == 0 {
		return "", ErrIncorrectUsernameOrPassword
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
			sessionID, err := NewSessionID()
			if err != nil {
				return "", fmt.Errorf(
					"%w: auth.LoginToStaffAccount(): failed to create session ID: %v", ErrInternal, err,
				)
			}

			// Get authRoles of the user.
			authRoles, err := staffAcc.QueryAuthRoles().All(ctx)
			if err != nil {
				return "", fmt.Errorf(
					"%w: auth.LoginToStaffAccount(): could not retrieve auth roles of the staff account: %v",
					ErrInternal,
					err,
				)
			}

			authRoleValues := make([]string, len(authRoles))
			for i, authRole := range authRoles {
				authRoleValues[i] = authRole.Value.String()
			}

			// Store user ID in Redis.
			_, err = authService.redisService.Client().Set(
				ctx,
				fmt.Sprintf("%s:%s:userId", SessionRedisKey, sessionID),
				staffAcc.ID.String(), authService.Config().SessionTimeout,
			).Result()
			if err != nil {
				return "", fmt.Errorf(
					"%w: auth.LoginToStaffAccount(): could not store user ID in redis: %v", ErrInternal, err,
				)
			}

			// Store authRoles in Redis.
			_, err = authService.redisService.Client().SAdd(
				ctx,
				fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID), authRoleValues,
			).Result()
			if err != nil {
				return "", fmt.Errorf(
					"%w: auth.LoginToStaffAccount(): could not store auth roles in redis: %v", ErrInternal, err,
				)
			}

			// Set expiry on the authRoles key as well.
			_, err = authService.redisService.Client().Expire(
				ctx, fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID), authService.Config().SessionTimeout,
			).Result()
			if err != nil {
				return "", fmt.Errorf(
					"%w: auth.LoginToStaffAccount(): could not set expiry for auth role key in Redis: %v", ErrInternal,
					err,
				)
			}

			return sessionID, nil
		}
	}

	// If we reach here, no accounts match the given username and password. Return an error.
	return "", ErrIncorrectUsernameOrPassword
}
