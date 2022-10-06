package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/chenningg/hermitboard-api/auth"
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
				if _, err := auth.ParseSessionID(sessionIDStr); err != nil {
					http.Error(w, "invalid session ID format", http.StatusUnauthorized)
					panic(err)
				}

				// Retrieve session from Redis and reset its expiry if found.
				session, err := authService.GetSessionFromStore(ctx, auth.SessionID(sessionIDStr))
				if err != nil {
					// Session key not found, probably expired. Redirect to log in.
					if errors.Is(err, auth.ErrNotFound) {
						http.Error(w, "session ID not found or expired", http.StatusNotFound)
					} else {
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					}
				}

				// Store the Session struct into context.
				ctx = context.WithValue(ctx, auth.SessionContextKey, session)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
