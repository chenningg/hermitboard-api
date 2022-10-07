package auth

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/graph"
)

func (authService AuthService) Logout(ctx context.Context) error {
	session := GetSessionFromContext(ctx)

	// If not logged in, then naturally you cannot log out.
	if !IsLoggedIn(session) {
		return fmt.Errorf(
			"auth.Logout(): %w",
			graph.NewGraphQLError("you are not logged in, unable to logout", graph.Unauthenticated),
		)
	}

	// Otherwise delete all sessions from Redis.
	_, err := authService.redisService.Client().Del(
		ctx, fmt.Sprintf("%s:%s", SessionRedisKey, session.Token),
	).Result()
	if err != nil {
		return fmt.Errorf(
			"auth.Logout(): %w: %v",
			graph.NewGraphQLError("session could not be terminated", graph.InternalServerError), err,
		)
	}

	return nil
}
