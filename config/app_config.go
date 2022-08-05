package config

import (
	"github.com/jellydator/validation"
)

type AppConfig struct {
	Env AppEnv `env:"ENV" envDefault:"Development"`
}

func (appConfig AppConfig) Validate() error {
	return validation.ValidateStruct(&appConfig,
		validation.Field(&appConfig.Env, validation.NotNil),
	)
}
