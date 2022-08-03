package config

import (
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

type DbConfig struct {
	Url      string
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT" envDefault:"5432"`
	Name     string `env:"NAME"`
}

func (dbConfig DbConfig) Validate() error {
	return validation.ValidateStruct(&dbConfig,
		validation.Field(&dbConfig.Url, validation.Required),
		validation.Field(&dbConfig.User, validation.Required),
		validation.Field(&dbConfig.Password, validation.Required),
		validation.Field(&dbConfig.Host, validation.Required, is.Host),
		validation.Field(&dbConfig.Port, validation.Required, is.Port),
		validation.Field(&dbConfig.Name, validation.Required),
	)
}
