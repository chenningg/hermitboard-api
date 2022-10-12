package graph

import (
	"context"
	"errors"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/chenningg/hermitboard-api/resperror"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// AddGraphQLError unwraps an error to find a wrapped HTTP error and populate it in the GraphQL response.
func AddGraphQLError(ctx context.Context, err error) {
	var graphQLError *resperror.GraphQLError
	if errors.As(err, &graphQLError) {
		graphql.AddError(
			ctx, &gqlerror.Error{
				Path:    graphql.GetPath(ctx),
				Message: graphQLError.Msg,
				Extensions: map[string]interface{}{
					"code":           graphQLError.Code,
					"responseStatus": resperror.GraphQLErrorCodeToStatus(graphQLError.Code),
				},
			},
		)
	}
}

// ErrorPresenter captures errors returned from resolvers and formats them.
func ErrorPresenter(ctx context.Context, e error) *gqlerror.Error {
	log.Println(e)
	// If we manage to find our custom error, return our error instead.
	var graphQLError *resperror.GraphQLError
	if errors.As(e, &graphQLError) {
		return &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: graphQLError.Msg,
			Extensions: map[string]interface{}{
				"code":           graphQLError.Code,
				"responseStatus": resperror.GraphQLErrorCodeToStatus(graphQLError.Code),
			},
		}
	}

	// Otherwise, call the normal graphQL error presenter to add the request path to the error.
	err := graphql.DefaultErrorPresenter(ctx, e)
	return err
}

// Recover captures panics from resolver functions and turns them into errors.
func Recover(ctx context.Context, err interface{}) error {
	// TODO: Notify bug tracker.
	return gqlerror.Errorf("internal server error")
}
