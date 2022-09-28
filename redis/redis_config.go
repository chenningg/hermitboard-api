package redis

import (
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

type RedisConfig struct {
	Url      string
	User     string `env:"USER" envDefault:"default"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT" envDefault:"14552"`
	Name     string `env:"NAME"`
}

func (redisConfig RedisConfig) Validate() error {
	return validation.ValidateStruct(
		&redisConfig,
		validation.Field(&redisConfig.Url, validation.Required),
		validation.Field(&redisConfig.User, validation.Required),
		validation.Field(&redisConfig.Password, validation.Required),
		validation.Field(&redisConfig.Host, validation.Required, is.Host),
		validation.Field(&redisConfig.Port, validation.Required, is.Port),
		validation.Field(&redisConfig.Name, validation.Required),
	)
}
