package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/chenningg/hermitboard-api/auth"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/pulid"
	goRedis "github.com/go-redis/redis/v9"
)

// Auth checks for a session ID in the Authorization header and hydrates the context with the session ID of the requester.
func Auth(authService auth.AuthServicer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Check for session ID in Authorization header.
			sessionIDStr := r.Header.Get(auth.AuthorizationHeader)
			ctx := r.Context()

			// If sessionID exists, hydrate the request context with the authorization scopes.
			if sessionIDStr != "" {
				// Check if session ID is a valid format.
				sessionID, err := auth.ParseSessionID(sessionIDStr)
				if err != nil {
					http.Error(w, "invalid session ID format", http.StatusUnauthorized)
					panic(err)
				}

				// Retrieve userID from Redis
				userIDStr, err := authService.RedisService().Client().Get(
					ctx, fmt.Sprintf("%s:%s:userID", auth.SessionRedisKey, sessionID),
				).Result()
				if err != nil {
					// Session key not found, probably expired. Redirect to log in.
					if errors.Is(err, goRedis.Nil) {
						http.Error(w, "session ID not found or expired", http.StatusUnauthorized)
					} else {
						// Throw panic, which will be recovered and internal server error returned and logged.
						panic(err)
					}
				}
				userID := pulid.PULID(userIDStr)
				if err = pulid.IsAPULID(userID); err != nil {
					http.Error(w, "session user ID is invalid", http.StatusUnauthorized)
				}

				// Retrieve authRoles from Redis
				authRolesSet, err := authService.RedisService().Client().SMembers(
					ctx, fmt.Sprintf("%s:%s:authRoles", auth.SessionRedisKey, sessionID),
				).Result()
				if err != nil {
					// Session key not found, probably expired. Redirect to log in.
					if errors.Is(err, goRedis.Nil) {
						http.Error(w, "session ID not found or expired", http.StatusUnauthorized)
					} else {
						// Throw panic, which will be recovered and internal server error returned and logged.
						panic(err)
					}
				}

				// Scan the results into a Session struct.
				var authRoles = map[authrole.Value]struct{}{}
				for _, authRole := range authRolesSet {
					authRoleValue := authrole.Value(authRole)
					if err = authrole.ValueValidator(authRoleValue); err != nil {
						http.Error(w, "invalid auth role value in stored session", http.StatusUnauthorized)
					}
					authRoles[authRoleValue] = struct{}{}
				}

				session := auth.Session{SessionID: sessionID, UserID: userID, AuthRoles: authRoles}

				// Store the Session struct into context.
				ctx = context.WithValue(ctx, auth.SessionContextKey, session)

				// Reset the expiry for the session.
				expiration := authService.Config().SessionTimeout
				if err != nil {
					http.Error(
						w, "could not convert seconds to duration for session timeout", http.StatusInternalServerError,
					)
				}

				_, err = authService.RedisService().Client().Expire(
					ctx, fmt.Sprintf("%s:%s:userID", auth.SessionRedisKey, sessionID),
					expiration,
				).Result()
				if err != nil {
					http.Error(w, "could not refresh session token in cache", http.StatusInternalServerError)
				}

				_, err = authService.RedisService().Client().Expire(
					ctx, fmt.Sprintf("%s:%s:authRoles", auth.SessionRedisKey, sessionID),
					expiration,
				).Result()
				if err != nil {
					http.Error(w, "could not refresh session token in cache", http.StatusInternalServerError)
				}
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
