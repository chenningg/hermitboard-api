package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/chenningg/hermitboard-api/db"
	"github.com/chenningg/hermitboard-api/graph"
	"github.com/chenningg/hermitboard-api/redis"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db    db.DbServicer
	redis redis.RedisServicer
}

// NewSchema creates a graphql executable schema.
func NewSchema(db db.DbServicer, redis redis.RedisServicer) graphql.ExecutableSchema {
	return graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &Resolver{db, redis},
		},
	)
}
