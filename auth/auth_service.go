package auth

import (
	"context"

	"github.com/chenningg/hermitboard-api/db"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/redis"
	"github.com/go-logr/logr"
)

type AuthServicer interface {
	CreateAccount(ctx context.Context, account ent.CreateAccountInput) (*ent.Account, SessionID, error)
	CreateStaffAccount(ctx context.Context, staffAccount ent.CreateStaffAccountInput) (
		*ent.StaffAccount, SessionID, error,
	)
	LoginToLocalAccount(ctx context.Context, username string, password string) (*ent.Account, error)
	LoginToLocalStaffAccount(ctx context.Context, username string, password string) (*ent.StaffAccount, error)
	LogoutFromAccount(ctx context.Context) error
	LogoutFromStaffAccount(ctx context.Context) error
}

type AuthService struct {
	logger       logr.Logger
	config       AuthConfig
	dbService    db.DbServicer
	redisService redis.RedisServicer
}

func NewAuthService(
	authConfig AuthConfig, logger logr.Logger, dbService db.DbServicer, redisService redis.RedisServicer,
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
