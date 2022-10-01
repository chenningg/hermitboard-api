package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/staffaccount"
	"golang.org/x/crypto/bcrypt"
)

func (authService AuthService) LoginToLocalAccount(ctx context.Context, username string, password string) (
	*ent.Account, SessionID, error,
) {
	// Username can be either email or nickname.
	dbClient := ent.FromContext(ctx)

	accounts, err := dbClient.Account.Query().Where(
		account.Or(
			account.NicknameEQ(username), account.EmailEQ(username),
		),
	).All(ctx)
	if err != nil {
		return nil, "", err
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
				return nil, "", fmt.Errorf("failed to create session ID: %v", err)
			}

			// Get expiration time for a session.
			expiration, err := time.ParseDuration(fmt.Sprintf("%ds", authService.config.SessionTimeout))
			if err != nil {
				return nil, "", fmt.Errorf("could not parse session expiry time: %v", err)
			}

			// Store user ID in Redis.
			_, err = authService.redisService.Client().Set(
				ctx,
				fmt.Sprintf("%s:%s:userId", SessionRedisKey, sessionID),
				acc.ID.String(), expiration,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf("could not store user ID in Redis: %v", err)
			}

			// Get authRoles of the user.
			authRoles, err := acc.QueryAuthRoles().All(ctx)
			if err != nil {
				return nil, "", fmt.Errorf("could not retrieve auth roles of the account: %v", err)
			}

			var authRoleValues []authrole.Value
			for _, authRole := range authRoles {
				authRoleValues = append(authRoleValues, authRole.Value)
			}

			// Store authRoles in Redis.
			_, err = authService.redisService.Client().RPush(
				ctx,
				fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID),
				authRoleValues,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf("could not store auth roles in Redis: %v", err)
			}

			// Set expiry on the authRoles key as well.
			_, err = authService.redisService.Client().Expire(
				ctx, fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID), expiration,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf("could not set expiry for auth role key in Redis: %v", err)
			}

			return acc, sessionID, nil
		}
	}

	// If we reach here, no accounts match the given username and password. Return an error.
	return nil, "", ErrIncorrectUsernameOrPassword
}

func (authService AuthService) LoginToLocalStaffAccount(
	ctx context.Context, username string, password string,
) (*ent.StaffAccount, SessionID, error) {
	// Username can be either email or nickname.
	dbClient := ent.FromContext(ctx)

	staffAccounts, err := dbClient.StaffAccount.Query().Where(
		staffaccount.Or(
			staffaccount.NicknameEQ(username), staffaccount.EmailEQ(username),
		),
	).All(ctx)
	if err != nil {
		return nil, "", err
	}

	for _, staffAccount := range staffAccounts {
		// Compare passwords.
		hashedPassword := staffAccount.Password
		if hashedPassword == nil {
			continue
		}
		err = bcrypt.CompareHashAndPassword([]byte(*staffAccount.Password), []byte(password))
		if err == nil {
			// We have a match of username and password, create the user session.
			sessionID, err := NewSessionID()
			if err != nil {
				return nil, "", fmt.Errorf("failed to create session ID: %v", err)
			}

			// Get expiration time for a session.
			expiration, err := time.ParseDuration(fmt.Sprintf("%ds", authService.config.SessionTimeout))
			if err != nil {
				return nil, "", fmt.Errorf("could not parse session expiry time: %v", err)
			}

			// Store user ID in Redis.
			_, err = authService.redisService.Client().Set(
				ctx,
				fmt.Sprintf("%s:%s:userId", SessionRedisKey, sessionID),
				staffAccount.ID.String(), expiration,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf("could not store user ID in Redis: %v", err)
			}

			// Get authRoles of the user.
			authRoles, err := staffAccount.QueryAuthRoles().All(ctx)
			if err != nil {
				return nil, "", fmt.Errorf("could not retrieve auth roles of the account: %v", err)
			}

			var authRoleValues []authrole.Value
			for _, authRole := range authRoles {
				authRoleValues = append(authRoleValues, authRole.Value)
			}

			// Store authRoles in Redis.
			_, err = authService.redisService.Client().RPush(
				ctx,
				fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID),
				authRoleValues,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf("could not store auth roles in Redis: %v", err)
			}

			// Set expiry on the authRoles key as well.
			_, err = authService.redisService.Client().Expire(
				ctx, fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, sessionID), expiration,
			).Result()
			if err != nil {
				return nil, "", fmt.Errorf("could not set expiry for auth role key in Redis: %v", err)
			}

			return staffAccount, sessionID, nil
		}
	}

	// If we reach here, no staff accounts match the given username and password. Return an error.
	return nil, "", ErrIncorrectUsernameOrPassword
}
