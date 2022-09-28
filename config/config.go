package config

import (
	"github.com/chenningg/hermitboard-api/auth"
	"github.com/chenningg/hermitboard-api/db"
	"github.com/chenningg/hermitboard-api/redis"
	"github.com/jellydator/validation"
)

type Config struct {
	App    AppConfig         `envPrefix:"APP_"`
	Server ServerConfig      `envPrefix:"SERVER_"`
	Db     db.DbConfig       `envPrefix:"DB_"`
	Redis  redis.RedisConfig `envPrefix:"REDIS_"`
	Auth   auth.AuthConfig   `envPrefix:"AUTH_"`
}

func (config Config) Validate() error {
	return validation.ValidateStruct(
		&config,
		validation.Field(&config.App, validation.Required),
		validation.Field(&config.Server, validation.Required),
		validation.Field(&config.Db, validation.Required),
		validation.Field(&config.Redis, validation.Required),
		validation.Field(&config.Auth, validation.Required),
	)
}

// Loader has a Load function that is called for each config struct.
type Loader interface {
	Load() error
}
