package resperror

// Define GraphQL error codes.
type GraphQLErrorCode = string

const (
	GQLParsedFailed               GraphQLErrorCode = "GRAPHQL_PARSE_FAILED"
	GQLValidationFailed           GraphQLErrorCode = "GRAPHQL_VALIDATION_FAILED"
	GQLBadUserInput               GraphQLErrorCode = "BAD_USER_INPUT"
	GQLUnauthenticated            GraphQLErrorCode = "UNAUTHENTICATED"
	GQLForbidden                  GraphQLErrorCode = "FORBIDDEN"
	GQLPersistedQueryNotFound     GraphQLErrorCode = "PERSISTED_QUERY_NOT_FOUND"
	GQLPersistedQueryNotSupported GraphQLErrorCode = "PERSISTED_QUERY_NOT_SUPPORTED"
	GQLInternalServerError        GraphQLErrorCode = "INTERNAL_SERVER_ERROR"
)

type GraphQLError struct {
	Msg  string
	Code GraphQLErrorCode
}

var graphQLErrorCodeToStatus = map[GraphQLErrorCode]int{
	GQLParsedFailed:               400,
	GQLValidationFailed:           400,
	GQLBadUserInput:               400,
	GQLUnauthenticated:            401,
	GQLForbidden:                  403,
	GQLPersistedQueryNotFound:     404,
	GQLPersistedQueryNotSupported: 415,
	GQLInternalServerError:        500,
}

// Error implements the error interface.
func (graphQLError *GraphQLError) Error() string {
	return graphQLError.Msg
}

// NewGraphQLError creates a new GraphQLError. The message must be safe for external consumption.
func NewGraphQLError(msg string, code GraphQLErrorCode) error {
	return &GraphQLError{Msg: msg, Code: code}
}

func GraphQLErrorCodeToStatus(code GraphQLErrorCode) int {
	if statusCode, ok := graphQLErrorCodeToStatus[code]; ok {
		return statusCode
	} else {
		return 500
	}
}
