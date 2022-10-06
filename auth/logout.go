package auth

import (
	"context"
	"fmt"
)

func (authService AuthService) Logout(ctx context.Context) error {
	session := GetSessionFromContext(ctx)

	// If not logged in, then naturally you cannot log out.
	if !IsLoggedIn(session) {
		return fmt.Errorf("%w: you are not logged in, unable to logout", ErrBadInput)
	}

	// Otherwise delete all sessions from Redis.
	_, err := authService.redisService.Client().Del(
		ctx, fmt.Sprintf("%s:%s", SessionRedisKey, session.ID),
	).Result()
	if err != nil {
		return fmt.Errorf("%w: unable to delete session to logout: %v", ErrInternal, err)
	}

	return nil
}
