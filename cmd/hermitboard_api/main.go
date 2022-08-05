package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chenningg/hermitboard-api/config"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
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
	dbPool, err := pgxpool.Connect(context.Background(), config.Db.Url)
	if err != nil {
		// Log database connection error and panic.
		logger.Error(err, "unable to connect to database", "package", "config")
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	var greeting string
	err = dbPool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
