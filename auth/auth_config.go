package auth

import (
	"time"

	"github.com/jellydator/validation"
)

type AuthConfig struct {
	BcryptCost     int           `env:"BCRYPT_COST" envDefault:"10"`
	SessionTimeout time.Duration `env:"SESSION_TIMEOUT" envDefault:"120h"`
}

func (authConfig AuthConfig) Validate() error {
	return validation.ValidateStruct(
		&authConfig,
		validation.Field(&authConfig.BcryptCost, validation.NotNil, validation.Min(8)),
		validation.Field(&authConfig.SessionTimeout, validation.NotNil),
	)
}
