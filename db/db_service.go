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
	Queries() Queries
	MigrateUp() (int, error)
}

type dbService struct {
	logger  logr.Logger
	config  DbConfig
	queries Queries
}

func NewDbService(dbConfig DbConfig, logger logr.Logger) (*dbService, error) {
	var dbService = new(dbService)

	// Initialize logger.
	dbService.logger = logger
	dbService.logger.Info("initializing db service")
	dbService.config = dbConfig

	// Initialize migration if auto migrate is set to true.
	if dbService.config.AutoMigrate {
		dbService.logger.V(1).Info("auto migration is set to true, conducting migration")
		_, err := dbService.MigrateUp()
		if err != nil {
			return nil, fmt.Errorf("NewDbService(): failed to conduct auto migration: %v", err)
		}
	}

	// Set database pool configuration.
	poolConfig, err := pgxpool.ParseConfig(dbConfig.Url)
	if err != nil {
		return nil, fmt.Errorf("NewDbService(): unable to parse database pool configuration: %v", err)
	}

	poolConfig.MaxConns = dbConfig.PoolMaxConn
	poolConfig.MinConns = dbConfig.PoolMinConn

	// TODO: Handle custom database types.
	poolConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return nil
	}

	// Connect to database pool.
	dbPool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("NewDbService(): unable to connect to database pool: %v", err)
	}

	// Create new Queries (SQLC) object with database pool.
	dbService.queries = *New(dbPool)

	return dbService, nil
}

func (dbService *dbService) Queries() *Queries {
	return &dbService.queries
}

func (dbService *dbService) MigrateUp() (int, error) {
	// Connect to database. If this fails, panic.
	dbService.logger.V(2).Info("connecting to database for migrations")
	db, err := sql.Open("pgx", dbService.config.Url)
	if err != nil {
		return 0, fmt.Errorf("unable to connect to database")
	}

	defer db.Close()

	// Set migrations file source.
	migrations := &migrate.FileMigrationSource{
		Dir: dbService.config.Migrations,
	}

	// Set migration table name.
	migrate.SetTable(dbService.config.MigrationsTableName)

	// Carry out the migrations.
	numMigrations, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return 0, fmt.Errorf("database migration failed")
	}

	dbService.logger.Info("database migration complete", "migrationsRun", numMigrations)

	return numMigrations, nil
}
