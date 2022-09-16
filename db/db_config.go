package db

import (
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

type DbConfig struct {
	Url                   string
	User                  string `env:"USER" envDefault:"postgres"`
	Password              string `env:"PASSWORD"`
	Host                  string `env:"HOST"`
	Port                  string `env:"PORT" envDefault:"5432"`
	Name                  string `env:"NAME"`
	MigrationsDir         string `env:"MIGRATIONS_DIR" envDefault:"ent/migrate/migrations"`
	AutoMigrate           bool   `env:"AUTO_MIGRATE" envDefault:"false"`
	AutoMigrateDropIndex  bool   `env:"AUTO_MIGRATE_DROP_INDEX" envDefault:"false"`
	AutoMigrateDropColumn bool   `env:"AUTO_MIGRATE_DROP_COLUMN" envDefault:"false"`
}

func (dbConfig DbConfig) Validate() error {
	return validation.ValidateStruct(&dbConfig,
		validation.Field(&dbConfig.Url, validation.Required),
		validation.Field(&dbConfig.User, validation.Required),
		validation.Field(&dbConfig.Password, validation.Required),
		validation.Field(&dbConfig.Host, validation.Required, is.Host),
		validation.Field(&dbConfig.Port, validation.Required, is.Port),
		validation.Field(&dbConfig.Name, validation.Required),
		validation.Field(&dbConfig.MigrationsDir, validation.Required),
		validation.Field(&dbConfig.AutoMigrate, validation.Required),
		validation.Field(&dbConfig.AutoMigrateDropIndex, validation.Required),
		validation.Field(&dbConfig.AutoMigrateDropColumn, validation.Required),
	)
}
