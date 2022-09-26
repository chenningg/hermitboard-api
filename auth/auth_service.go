package auth

import (
	"github.com/go-logr/logr"
)

type AuthServicer interface {
	CreateAccount()
	CreateStaffAccount()
}

type AuthService struct {
	logger logr.Logger
	config AuthConfig
}

func NewAuthService(authConfig AuthConfig, logger logr.Logger) (*AuthService, error) {
	var authService = new(AuthService)

	// Initialize logger.
	authService.logger = logger
	authService.logger.V(2).Info("NewAuthService(): initializing auth service")
	authService.config = authConfig

	return authService, nil
}
