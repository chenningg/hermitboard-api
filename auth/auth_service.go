package auth

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/chenningg/hermitboard-api/db/seed"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/staffaccountauthrole"
	"github.com/go-logr/logr"
)

type AuthServicer interface {
	CreateAccount()
	CreateStaffAccount()
}

type DbService struct {
	logger logr.Logger
	config DbConfig
	client *ent.Client
}

func NewDbService(dbConfig DbConfig, logger logr.Logger) (*DbService, error) {
	var dbService = new(DbService)

	// Initialize logger.
	dbService.logger = logger
	dbService.logger.V(2).Info("NewDbService(): initializing db service")
	dbService.config = dbConfig

	// Connect to database
	err := dbService.open()
	if err != nil {
		return dbService, fmt.Errorf("NewDbService(): error opening connection to database %v", err)
		dbService.logger.Error(err, "NewDbService(): error opening connection to database")
	}

	ctx := context.Background()

	// Initialize migration if auto migrate is set to true.
	if dbService.config.AutoMigrate {
		dbService.logger.V(1).Info("NewDbService(): auto migration is set to true, conducting migration")
		err := dbService.migrateUp(ctx)
		if err != nil {
			return dbService, fmt.Errorf("NewDbService(): migration failed to create schema resources %v", err)
			dbService.logger.Error(err, "NewDbService(): migration failed to create schema resources")
		}
	}
	// Seed the database if database has not been seeded yet.
	// We check whether database is seeded or not by counting the number of staff accounts (whether owner exists).
	numStaffAccounts, err := dbService.Client().StaffAccountAuthRole.Query().Where(staffaccountauthrole.HasAuthRoleWith(authrole.AuthRoleEQ(authrole.AuthRoleSuperAdmin))).Count(ctx)
	if err != nil {
		return dbService, fmt.Errorf("NewDbService(): failed to check database seeding status %v", err)
		dbService.logger.Error(err, "NewDbService(): failed to check database seeding status")
	}
	// If not staff account with SuperAdmin role has been created yet, then we do database seeding.
	if numStaffAccounts == 0 {
		dbService.logger.V(1).Info("NewDbService(): database is empty, conducting database seeding")
		err = dbService.seedDb(ctx)
		if err != nil {
			return dbService, fmt.Errorf("NewDbService(): failed to seed database %v", err)
			dbService.logger.Error(err, "NewDbService(): failed to seed database")
		}
	}

	return dbService, nil
}

func (dbService *DbService) Client() *ent.Client {
	return dbService.client
}

func (dbService *DbService) Close() {
	dbService.Client().Close()
}

func (dbService *DbService) open() error {
	db, err := sql.Open("pgx", dbService.config.Url)
	if err != nil {
		return fmt.Errorf("open(): %v", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	dbService.client = ent.NewClient(ent.Driver(drv))
	return nil
}

func (dbService *DbService) migrateUp(ctx context.Context) error {
	err := dbService.Client().Schema.Create(
		ctx,
		migrate.WithDropIndex(dbService.config.AutoMigrateDropIndex),
		migrate.WithDropColumn(dbService.config.AutoMigrateDropColumn),
	)
	if err != nil {
		return fmt.Errorf("migrateUp(), %v", err)
	}

	return nil
}

func (dbService *DbService) seedDb(ctx context.Context) error {
	err := WithTx(
		ctx, dbService.Client(), func(tx *ent.Tx) error {
			return seed.Seed(ctx, tx.Client())
		},
	)
	if err != nil {
		return fmt.Errorf("seedDb(): %v", err)
	}

	return nil
}
