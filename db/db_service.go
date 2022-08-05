package db

import (
	"database/sql"
	"fmt"

	"github.com/go-logr/logr"
	_ "github.com/jackc/pgx/v4/stdlib"
	migrate "github.com/rubenv/sql-migrate"
)

type DbService interface {
	Init(dbConfig DbConfig, logger logr.Logger) error
}

type dbService struct {
	config DbConfig
	logger logr.Logger
}

func NewDbService(dbConfig DbConfig, logger logr.Logger) (*dbService, error) {
	var dbService = new(dbService)

	// Load in configuration from .env or environment variables.
	err := dbService.Init(dbConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("NewDbService(): %w", err)
	}

	return dbService, nil
}

func (dbService *dbService) Init(dbConfig DbConfig, logger logr.Logger) error {
	dbService.logger = logger
	dbService.logger.Info("initializing db service")
	dbService.config = dbConfig

	// Carry out migrations if auto migrate is set to true.
	if dbService.config.AutoMigrate {
		dbService.logger.Info("carrying out database migrations")

		db, err := sql.Open("pgx", dbService.config.Url)
		if err != nil {
			return fmt.Errorf("Init(): unable to connect to database: %w", err)
		}

		migrations := migrate.FileMigrationSource{
			Dir: dbService.config.Migrations,
		}

		numMigrations, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
		if err != nil {
			return fmt.Errorf("Init(): database migration failed: %w", err)
		}

		dbService.logger.Info(fmt.Sprintf("database migration complete, successfully carried out %d migrations", numMigrations))

		// Close the database.
		db.Close()
	}

	return nil
}
