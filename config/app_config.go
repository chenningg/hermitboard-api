package config

import (
	"github.com/chenningg/hermitboard-api/hbtype"
	"github.com/jellydator/validation"
)

type AppConfig struct {
	Env      hbtype.AppEnv   `env:"ENV" envDefault:"Development"`
	AuthType hbtype.AuthType `env:"AUTH_TYPE" envDefault:"Local"`
}

func (appConfig AppConfig) Validate() error {
	return validation.ValidateStruct(&appConfig,
		validation.Field(&appConfig.Env, validation.NotNil),
		validation.Field(&appConfig.AuthType, validation.NotNil),
	)
}
