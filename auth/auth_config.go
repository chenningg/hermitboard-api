package auth

import "github.com/jellydator/validation"

type AuthConfig struct {
	BcryptCost int `env:"BCRYPT_COST" envDefault:"10"`
}

func (authConfig AuthConfig) Validate() error {
	return validation.ValidateStruct(
		&authConfig,
		validation.Field(&authConfig.BcryptCost, validation.NotNil, validation.Min(8)),
	)
}
