//go:build ignore

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chenningg/hermitboard-api/config"
	"github.com/chenningg/hermitboard-api/ent/migrate"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
)

func main() {
	// Load in logger.
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"
	zl := zerolog.New(os.Stderr).With().Timestamp().Logger()
	var logger logr.Logger = zerologr.New(&zl).WithName("hermitboard-api").WithName("atlas-migrate")

	// Load in configuration.
	configService, err := config.NewConfigService(logger.WithName("config"))
	if err != nil {
		// Log configuration error and panic.
		logger.Error(err, "migrate: config could not be loaded")
		os.Exit(1)
	}
	config := configService.Config()

	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir(config.Db.MigrationsDir)
	if err != nil {
		// If there's an error creating the migrations' directory, panic.
		logger.Error(
			err, "migrate: failed creating atlas migration directory", "migrationsDir", config.Db.MigrationsDir,
		)
		os.Exit(1)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	if len(os.Args) != 2 {
		logger.Error(
			fmt.Errorf("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'"),
			"migrate: missing migration name",
		)
		os.Exit(1)
	}

	// Generate migrations using Atlas support for Postgres
	err = migrate.NamedDiff(ctx, config.Db.Url, os.Args[1], opts...)
	if err != nil {
		logger.Error(err, "migrate: failed generating migration file", "migrationFilename", os.Args[1])
		os.Exit(1)
	}
}
