package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Define GraphQL status codes.
type GraphQLErrorCode = string

const (
	GraphQLParsedFailed        GraphQLErrorCode = "GRAPHQL_PARSE_FAILED"
	GraphQlValidationFailed    GraphQLErrorCode = "GRAPHQL_VALIDATION_FAILED"
	BadUserInput               GraphQLErrorCode = "BAD_USER_INPUT"
	Unauthenticated            GraphQLErrorCode = "UNAUTHENTICATED"
	Forbidden                  GraphQLErrorCode = "FORBIDDEN"
	PersistedQueryNotFound     GraphQLErrorCode = "PERSISTED_QUERY_NOT_FOUND"
	PersistedQueryNotSupported GraphQLErrorCode = "PERSISTED_QUERY_NOT_SUPPORTED"
	InternalServerError        GraphQLErrorCode = "INTERNAL_SERVER_ERROR"
)

type GraphQLError struct {
	Msg  string
	Code GraphQLErrorCode
}

var CodeToStatus = map[GraphQLErrorCode]int{
	GraphQLParsedFailed:        400,
	GraphQlValidationFailed:    400,
	BadUserInput:               400,
	Unauthenticated:            401,
	Forbidden:                  403,
	PersistedQueryNotFound:     404,
	PersistedQueryNotSupported: 415,
	InternalServerError:        500,
}

// Error implements the error interface.
func (graphQLError GraphQLError) Error() string {
	return graphQLError.Msg
}

// NewGraphQLError creates a new GraphQLError. The message must be safe for external consumption.
func NewGraphQLError(msg string, code GraphQLErrorCode) GraphQLError {
	return GraphQLError{Msg: msg, Code: code}
}

// AddGraphQLError unwraps an error to find a wrapped HTTP error and populate it in the GraphQL response.
func AddGraphQLError(ctx context.Context, err error) {
	var graphQLError *GraphQLError
	if errors.As(err, graphQLError) {
		graphql.AddError(
			ctx, &gqlerror.Error{
				Path:    graphql.GetPath(ctx),
				Message: graphQLError.Msg,
				Extensions: map[string]interface{}{
					"code":           graphQLError.Code,
					"responseStatus": CodeToStatus[graphQLError.Code],
				},
			},
		)
	}
}

// ErrorPresenter captures errors returned from resolvers and formats them.
func ErrorPresenter(ctx context.Context, e error) *gqlerror.Error {
	// First, call the normal graphQL error presenter to add the request path to the error.
	err := graphql.DefaultErrorPresenter(ctx, e)

	// Turn the error into our custom error and add extra fields.
	var graphQLError *GraphQLError
	if errors.As(err, &graphQLError) {
		err.Message = graphQLError.Msg
		err.Extensions = map[string]interface{}{
			"code":           graphQLError.Code,
			"responseStatus": CodeToStatus[graphQLError.Code],
		}
	}

	return err
}

// Recover captures panics from resolver functions and turns them into errors.
func Recover(ctx context.Context, err interface{}) error {
	// TODO: Notify bug tracker.
	return gqlerror.Errorf("internal server error")
}
