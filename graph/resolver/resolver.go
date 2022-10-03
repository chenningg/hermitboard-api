package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/chenningg/hermitboard-api/auth"
	"github.com/chenningg/hermitboard-api/db"
	"github.com/chenningg/hermitboard-api/graph"
	"github.com/chenningg/hermitboard-api/redis"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	dbService    db.DbServicer
	redisService redis.RedisServicer
	authService  auth.AuthServicer
}

// NewSchema creates a graphql executable schema.
func NewSchema(
	dbService db.DbServicer, redisService redis.RedisServicer, authService auth.AuthServicer,
) graphql.ExecutableSchema {
	return graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &Resolver{dbService, redisService, authService},
		},
	)
}
