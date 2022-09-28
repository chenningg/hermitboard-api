package auth

import "github.com/jellydator/validation"

type AuthConfig struct {
	BcryptCost     int `env:"BCRYPT_COST" envDefault:"10"`
	SessionTimeout int `env:"SESSION_TIMEOUT" envDefault:"604800"`
}

func (authConfig AuthConfig) Validate() error {
	return validation.ValidateStruct(
		&authConfig,
		validation.Field(&authConfig.BcryptCost, validation.NotNil, validation.Min(8)),
		validation.Field(&authConfig.SessionTimeout, validation.NotNil),
	)
}
