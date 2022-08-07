package config

import (
	"github.com/chenningg/hermitboard-api/db"
	"github.com/jellydator/validation"
)

type Config struct {
	App    AppConfig    `envPrefix:"APP_"`
	Server ServerConfig `envPrefix:"SERVER_"`
	Db     db.DbConfig  `envPrefix:"DB_"`
}

func (config Config) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.App, validation.Required),
		validation.Field(&config.Server, validation.Required),
		validation.Field(&config.Db, validation.Required),
	)
}
