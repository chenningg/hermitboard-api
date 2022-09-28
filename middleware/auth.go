package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/chenningg/hermitboard-api/auth"
	"github.com/chenningg/hermitboard-api/redis"
	goRedis "github.com/go-redis/redis/v9"
)

// Session ID key type for context
type sessionContextKey int

const SessionContextKey sessionContextKey = 0

// AuthorizationHeader is the name of the HTTP Header which contains the session id.
// Exported so that it can be changed by developers.
var AuthorizationHeader = "Authorization"

// Auth checks for a session ID in the Authorization header and hydrates the context with the session ID of the requester.
func Auth(redisService redis.RedisServicer, authService auth.AuthServicer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Check for session ID in Authorization header.
			sessionIDStr := r.Header.Get("Authorization")
			ctx := r.Context()

			// If sessionID exists, hydrate the request context with the authorization scopes.
			if sessionIDStr != "" {
				// Check if session ID is a valid format.
				sessionID, err := auth.ParseSessionID(sessionIDStr)
				if err != nil {
					http.Error(w, "invalid session ID format", http.StatusUnauthorized)
					panic(err)
				}

				// Retrieve authorization scopes from Redis.
				res := redisService.Client().HGetAll(ctx, fmt.Sprintf("%s:%s", auth.SessionRedisKey, sessionID))
				if err = res.Err(); err != nil {
					// Session key not found, probably expired. Redirect to log in.
					if errors.Is(err, goRedis.Nil) {
						http.Error(w, "session ID not found or expired", http.StatusUnauthorized)
					} else {
						// Throw panic, which will be recovered and internal server error returned and logged.
						panic(err)
					}
				}

				// Scan the results into a Session struct.
				var session auth.Session
				if err := res.Scan(&session); err != nil {
					// Throw panic, which will be recovered and internal server error returned and logged.
					panic(err)
				}

				// Store the Session struct into context.
				ctx = context.WithValue(ctx, SessionContextKey, session)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

// GetSessionFromContext returns the Session from a provided context, or an empty string if it is not found.
func GetSessionFromContext(ctx context.Context) *auth.Session {
	if ctx == nil {
		return nil
	}
	if session, ok := ctx.Value(SessionContextKey).(auth.Session); ok {
		return &session
	}
	return nil
}
