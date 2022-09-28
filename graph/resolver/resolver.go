package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/graph"
	"github.com/go-redis/redis/v9"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	dbClient    *ent.Client
	redisClient *redis.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(dbClient *ent.Client, redisClient *redis.Client) graphql.ExecutableSchema {
	return graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &Resolver{dbClient, redisClient},
		},
	)
}
