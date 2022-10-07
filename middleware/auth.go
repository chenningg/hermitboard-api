package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/chenningg/hermitboard-api/auth"
	"github.com/chenningg/hermitboard-api/graph"
)

// Auth checks for a session ID in the Authorization header and hydrates the context with the session ID of the requester.
func Auth(authService auth.AuthServicer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Check for session ID in Authorization header.
			authorizationStr := r.Header.Get(auth.AuthorizationHeader)

			ctx := r.Context()

			// If sessionID exists, hydrate the request context with the authorization scopes.
			if authorizationStr != "" {
				authorizationStrArr := strings.Fields(authorizationStr)
				if len(authorizationStrArr) != 2 {
					http.Error(
						w, "invalid authorization token format, expecting `Bearer <token>`", http.StatusUnauthorized,
					)
					return
				}

				// Just ignore token type "Bearer"
				// tokenType := authorizationStr[0]
				sessionTokenStr := authorizationStrArr[1]

				// Check if session token is a valid format.
				sessionToken, err := auth.ParseSessionToken(sessionTokenStr)
				if err != nil {
					http.Error(w, "invalid authorization token format", http.StatusUnauthorized)
					return
				}

				// Retrieve session from Redis and reset its expiry if found.
				session, err := authService.GetSessionFromStore(ctx, sessionToken)
				if err != nil {
					// Session key not found, probably expired. Redirect to log in.
					var graphQLError *graph.GraphQLError
					if errors.As(err, graphQLError) {
						if graphQLError.Code == graph.InternalServerError {
							http.Error(w, graphQLError.Msg, http.StatusInternalServerError)
						} else {
							http.Error(w, graphQLError.Msg, http.StatusUnauthorized)
						}
					} else {
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					}
					return
				}

				// Store the Session struct into context.
				ctx = context.WithValue(ctx, auth.SessionContextKey, session)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
