package auth

import (
	"context"
	"fmt"
)

func (authService AuthService) LogoutFromAccount(ctx context.Context) error {
	session := GetSessionFromContext(ctx)

	// If not logged in, then naturally you cannot log out.
	if !IsLoggedIn(session) {
		return fmt.Errorf("%w: you are not logged in, unable to logout", ErrBadInput)
	}

	// Otherwise delete all sessions from Redis.
	_, err := authService.RedisService().Client().Del(
		ctx, fmt.Sprintf("%s:%s:userId", SessionRedisKey, session.SessionID),
	).Result()
	if err != nil {
		return fmt.Errorf("%w: unable to delete session to logout: %v", ErrInternal, err)
	}

	_, err = authService.RedisService().Client().Del(
		ctx, fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, session.SessionID),
	).Result()
	if err != nil {
		return fmt.Errorf("%w: unable to delete session to logout: %v", ErrInternal, err)
	}

	return nil
}

func (authService AuthService) LogoutFromStaffAccount(ctx context.Context) error {
	session := GetSessionFromContext(ctx)

	// If not logged in, then naturally you cannot log out.
	if !IsLoggedIn(session) {
		return fmt.Errorf("%w: you are not logged in, unable to logout", ErrBadInput)
	}

	// Otherwise delete all sessions from Redis.
	_, err := authService.RedisService().Client().Del(
		ctx, fmt.Sprintf("%s:%s:userId", SessionRedisKey, session.SessionID),
	).Result()
	if err != nil {
		return fmt.Errorf("%w: unable to delete session to logout: %v", ErrInternal, err)
	}

	_, err = authService.RedisService().Client().Del(
		ctx, fmt.Sprintf("%s:%s:authRoles", SessionRedisKey, session.SessionID),
	).Result()
	if err != nil {
		return fmt.Errorf("%w: unable to delete session to logout: %v", ErrInternal, err)
	}

	return nil
}
