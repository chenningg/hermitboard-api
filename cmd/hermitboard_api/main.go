package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chenningg/hermitboard-api/config"
	"github.com/chenningg/hermitboard-api/db"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/chenningg/hermitboard-api/graph/generated"
	"github.com/chenningg/hermitboard-api/graph/resolver"
)

func main() {
	// Load in logger.
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"
	zl := zerolog.New(os.Stderr).With().Timestamp().Logger()
	var logger logr.Logger = zerologr.New(&zl).WithName("hermitboard-api")

	// Load in configuration.
	configService, err := config.NewConfigService(logger.WithName("config"))
	if err != nil {
		// Log configuration error and panic.
		logger.Error(err, "config could not be loaded")
		os.Exit(1)
	}
	config := configService.Config()

	// Open database pool connection
	dbService, err := db.NewDbService(config.Db, logger.WithName("db"))
	if err != nil {
		// Log database error and panic.
		logger.Error(err, "database setup could not be completed")
		os.Exit(1)
	}

	// Defer the closing of the database pool.
	defer dbService.ClosePool()

	// Run the web server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/api"))
	http.Handle("/api", srv)

	logger.Info(fmt.Sprintf("connect to http://%s:%s/playground for GraphQL playground", config.Server.Host, config.Server.Port))
	logger.Error(http.ListenAndServe(config.Server.Host+":"+config.Server.Port, nil), "server error")
}
