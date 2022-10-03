package auth

import (
	"context"

	"github.com/chenningg/hermitboard-api/db"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/redis"
	"github.com/go-logr/logr"
)

type AuthServicer interface {
	Config() *AuthConfig
	RedisService() *redis.RedisService
	CreateAccount(
		ctx context.Context, input ent.CreateAccountInput,
	) (*ent.Account, error)
	CreateStaffAccount(
		ctx context.Context, input ent.CreateStaffAccountInput,
	) (*ent.StaffAccount, error)
	LoginToAccount(ctx context.Context, input LoginToAccountInput) (
		SessionID, error,
	)
	LoginToStaffAccount(
		ctx context.Context, input LoginToStaffAccountInput,
	) (SessionID, error)
	LogoutFromAccount(ctx context.Context) error
	LogoutFromStaffAccount(ctx context.Context) error
}

type AuthService struct {
	logger       logr.Logger
	config       AuthConfig
	dbService    *db.DbService
	redisService *redis.RedisService
}

func NewAuthService(
	authConfig AuthConfig, logger logr.Logger, dbService *db.DbService, redisService *redis.RedisService,
) *AuthService {
	var authService = new(AuthService)

	// Initialize logger.
	authService.logger = logger
	authService.logger.V(2).Info("NewAuthService(): initializing auth service")
	authService.config = authConfig

	authService.dbService = dbService
	authService.redisService = redisService

	return authService
}

func (authService AuthService) Config() *AuthConfig {
	return &authService.config
}

func (authService AuthService) RedisService() *redis.RedisService {
	return authService.redisService
}
