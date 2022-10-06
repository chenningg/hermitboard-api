package auth

import (
	"context"

	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/redis"
	"github.com/go-logr/logr"
)

type AuthServicer interface {
	Config() *AuthConfig
	GetSessionFromStore(ctx context.Context, sessionID SessionID) (*Session, error)
	CreateAccount(
		ctx context.Context, input ent.CreateAccountInput,
	) (*CreateAccountPayload, error)
	CreateStaffAccount(
		ctx context.Context, input ent.CreateStaffAccountInput,
	) (*CreateStaffAccountPayload, error)
	LoginToAccount(ctx context.Context, input LoginToAccountInput) (
		*LoginToAccountPayload, error,
	)
	LoginToStaffAccount(
		ctx context.Context, input LoginToStaffAccountInput,
	) (*LoginToStaffAccountPayload, error)
	Logout(ctx context.Context) error
}

type AuthService struct {
	logger       logr.Logger
	config       AuthConfig
	redisService *redis.RedisService
}

func NewAuthService(
	authConfig AuthConfig, logger logr.Logger, redisService *redis.RedisService,
) *AuthService {
	var authService = new(AuthService)

	// Initialize logger.
	authService.logger = logger
	authService.logger.V(2).Info("NewAuthService(): initializing auth service")
	authService.config = authConfig

	authService.redisService = redisService

	return authService
}

func (authService AuthService) Config() *AuthConfig {
	return &authService.config
}
