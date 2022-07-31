package config

import (
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

type ServerConfig struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port uint16 `env:"PORT" envDefault:"4000"`
}

func (serverConfig ServerConfig) Validate() error {
	return validation.ValidateStruct(&serverConfig,
		validation.Field(&serverConfig.Host, validation.Required, is.Host),
		validation.Field(&serverConfig.Port, validation.Required, is.Port),
	)
}
