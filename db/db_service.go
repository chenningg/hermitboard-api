package db

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/migrate"
	"github.com/go-logr/logr"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DbService interface {
	Client() *ent.Client
	Close()
}

type dbService struct {
	logger logr.Logger
	config DbConfig
	client *ent.Client
}

func NewDbService(dbConfig DbConfig, logger logr.Logger) (*dbService, error) {
	var dbService = new(dbService)

	// Initialize logger.
	dbService.logger = logger
	dbService.logger.Info("initializing db service")
	dbService.config = dbConfig

	// Connect to database
	err := dbService.open()
	if err != nil {
		dbService.logger.Error(err, "open(): error opening connection to database")
	}

	ctx := context.Background()

	// Initialize migration if auto migrate is set to true.
	if dbService.config.AutoMigrate {
		dbService.logger.V(1).Info("auto migration is set to true, conducting migration")
		err := dbService.migrateUp(ctx)
		if err != nil {
			dbService.logger.Error(err, "migrateUp(): migration failed to create schema resources")
		}
	}

	return dbService, nil
}

func (dbService *dbService) Client() *ent.Client {
	return dbService.client
}

func (dbService *dbService) Close() {
	dbService.Client().Close()
}

func (dbService *dbService) open() error {
	db, err := sql.Open("pgx", dbService.config.Url)
	if err != nil {
		return err
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	dbService.client = ent.NewClient(ent.Driver(drv))
	return nil
}

func (dbService *dbService) migrateUp(ctx context.Context) error {
	err := dbService.Client().Schema.Create(
		ctx,
		migrate.WithDropIndex(dbService.config.AutoMigrateDropIndex),
		migrate.WithDropColumn(dbService.config.AutoMigrateDropColumn),
	)
	if err != nil {
		return err
	}

	return nil
}
