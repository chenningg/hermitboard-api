package db

import (
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

type DbConfig struct {
	Url         string
	User        string `env:"USER" envDefault:"postgres"`
	Password    string `env:"PASSWORD"`
	Host        string `env:"HOST"`
	Port        string `env:"PORT" envDefault:"5432"`
	Name        string `env:"NAME"`
	PoolMaxConn int32  `env:"POOL_MAX_CONN" envDefault:"10"`
	PoolMinConn int32  `env:"POOL_MIN_CONN" envDefault:"0"`
	Migrations  string `env:"MIGRATIONS" envDefault:"postgres/migrations"`
	AutoMigrate bool   `env:"AUTO_MIGRATE" envDefault:"false"`
}

func (dbConfig DbConfig) Validate() error {
	return validation.ValidateStruct(&dbConfig,
		validation.Field(&dbConfig.Url, validation.Required),
		validation.Field(&dbConfig.User, validation.Required),
		validation.Field(&dbConfig.Password, validation.Required),
		validation.Field(&dbConfig.Host, validation.Required, is.Host),
		validation.Field(&dbConfig.Port, validation.Required, is.Port),
		validation.Field(&dbConfig.Name, validation.Required),
		validation.Field(&dbConfig.PoolMaxConn, validation.NotNil, validation.Min(dbConfig.PoolMaxConn), validation.Max(64)),
		validation.Field(&dbConfig.PoolMinConn, validation.NotNil, validation.Min(0), validation.Max(dbConfig.PoolMaxConn)),
		validation.Field(&dbConfig.Migrations, validation.Required),
		validation.Field(&dbConfig.AutoMigrate, validation.Required),
	)
}
