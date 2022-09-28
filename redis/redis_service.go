package redis

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-redis/redis/v9"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type RedisServicer interface {
	Client() *redis.Client
	Close() error
}

type RedisService struct {
	logger logr.Logger
	config RedisConfig
	client *redis.Client
}

func NewRedisService(redisConfig RedisConfig, logger logr.Logger) (*RedisService, error) {
	var redisService = new(RedisService)

	// Initialize logger.
	redisService.logger = logger
	redisService.logger.V(2).Info("NewRedisService(): initializing Redis service")
	redisService.config = redisConfig

	// Connect to Redis database
	opt, err := redis.ParseURL(redisConfig.Url)
	if err != nil {
		return redisService, fmt.Errorf("NewRedisService(): could not connect to Redis database: %v", err)
	}

	rdb := redis.NewClient(opt)
	redisService.client = rdb

	return redisService, nil
}

func (redisService *RedisService) Client() *redis.Client {
	return redisService.client
}

func (redisService *RedisService) Close() error {
	if err := redisService.Client().Close(); err != nil {
		return fmt.Errorf("redis.Close(): failed to properly close redis connection: %v", err)
	}
	return nil
}
