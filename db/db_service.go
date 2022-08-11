package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	migrate "github.com/rubenv/sql-migrate"
)

type DbService interface {
	Init(dbConfig DbConfig, logger logr.Logger) error
	Db() *pgxpool.Pool
	ClosePool()
}

type dbService struct {
	config DbConfig
	logger logr.Logger
	db     *pgxpool.Pool
}

func NewDbService(dbConfig DbConfig, logger logr.Logger) (*dbService, error) {
	var dbService = new(dbService)

	// Load in configuration from .env or environment variables.
	err := dbService.Init(dbConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("NewDbService(): %v", err)
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
			return fmt.Errorf("Init(): unable to connect to database: %v", err)
		}

		// Get directory for migration SQL files.
		migrations := migrate.FileMigrationSource{
			Dir: dbService.config.Migrations,
		}

		// Set migration table name.
		migrate.SetTable(dbService.config.MigrationsTableName)

		// Carry out the migrations.
		numMigrations, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
		if err != nil {
			return fmt.Errorf("Init(): database migration failed: %v", err)
		}

		dbService.logger.Info(fmt.Sprintf("database migration complete, successfully carried out %d migrations", numMigrations))

		// Close the database.
		db.Close()
	}

	// Set database pool configuration.
	poolConfig, err := pgxpool.ParseConfig(dbService.config.Url)
	if err != nil {
		return fmt.Errorf("Init(): unable to parse database pool configuration: %v", err)
	}

	poolConfig.MaxConns = dbService.config.PoolMaxConn
	poolConfig.MinConns = dbService.config.PoolMinConn

	// TODO: Handle custom database types.
	poolConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return nil
	}

	// Connect to database pool.
	dbPool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return fmt.Errorf("Init(): unable to connect to database pool: %v", err)
	}

	dbService.db = dbPool

	return nil
}

func (dbService *dbService) Pool() *pgxpool.Pool {
	return dbService.db
}

func (dbService *dbService) ClosePool() {
	dbService.db.Close()
}
