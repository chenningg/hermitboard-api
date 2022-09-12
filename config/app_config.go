package config

import (
	"github.com/chenningg/hermitboard-api/auth"
	"github.com/jellydator/validation"
)

type AppConfig struct {
	Env      AppEnv        `env:"ENV" envDefault:"Development"`
	AuthType auth.AuthType `env:"AUTH_TYPE" envDefault:"Local"`
}

func (appConfig AppConfig) Validate() error {
	return validation.ValidateStruct(&appConfig,
		validation.Field(&appConfig.Env, validation.NotNil),
		validation.Field(&appConfig.AuthType, validation.NotNil),
	)
}
