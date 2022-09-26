package main

import (
	"fmt"
	"github.com/chenningg/hermitboard-api/db"
	"net/http"
	"os"

	"github.com/chenningg/hermitboard-api/config"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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
		logger.Error(err, "main(): config could not be loaded")
		os.Exit(1)
	}
	config := configService.Config()

	// Open database connection
	dbService, err := db.NewDbService(config.Db, logger.WithName("db"))
	if err != nil {
		// Log database error and panic.
		logger.Error(err, "main(): database setup could not be completed")
		os.Exit(1)
	}

	// Defer closing of database connection till end of main function.
	defer dbService.Close()

	// Run the web server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/api"))
	http.Handle("/api", srv)

	logger.Info(
		fmt.Sprintf(
			"connect to http://%s:%s/playground for GraphQL playground", config.Server.Host, config.Server.Port,
		),
		"host", config.Server.Host,
		"port", config.Server.Port,
	)
	logger.Error(http.ListenAndServe(config.Server.Host+":"+config.Server.Port, nil), "main(): server error")
}
