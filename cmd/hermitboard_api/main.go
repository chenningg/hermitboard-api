package main

import (
	"os"

	"github.com/chenningg/hermitboard-api/config"
	"github.com/chenningg/hermitboard-api/db"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
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
	dbService, err := db.NewDbService(config.Db, logger.WithName("db"))
	if err != nil {
		// Log database error and panic.
		logger.Error(err, "database setup could not be completed")
		os.Exit(1)
	}

	// Defer the closing of the database pool.
	defer dbService.ClosePool()
}
