package config

import (
	"fmt"
	"os"

	envLoader "github.com/caarlos0/env/v6"
	"github.com/chenningg/hermitboard-api/url"
	"github.com/go-logr/logr"
	"github.com/joho/godotenv"
)

type ConfigService interface {
	Init(logger logr.Logger) error
	Config() *Config
}

type configService struct {
	config Config
	logger logr.Logger
}

func NewConfigService(logger logr.Logger) (*configService, error) {
	var configService = new(configService)

	// Load in configuration from .env or environment variables.
	err := configService.Init(logger)
	if err != nil {
		return nil, fmt.Errorf("NewConfigService(): %w", err)
	}

	return configService, nil
}

func (configService *configService) Init(logger logr.Logger) error {
	configService.logger = logger
	configService.logger.Info("initializing config service")

	// Load environment variable, defaults to development.
	env := os.Getenv("HB_APP_ENV")
	if env == "" {
		env = "development"
	}

	// Load environment files in order (files loaded first take precedence).
	// Load .local files.
	godotenv.Load(".env." + env + ".local")

	// Load .env files.
	godotenv.Load(".env." + env)

	// Load generic .env file.
	godotenv.Load()

	config := Config{}
	opts := envLoader.Options{Prefix: "HB_"}

	// Unmarshal into configuration struct.
	err := envLoader.Parse(&config, opts)
	if err != nil {
		return fmt.Errorf("Init(): failed to parse config provided: %w", err)
	}

	// Clean and encode all URLs
	config.Db.Password = url.Encode(config.Db.Password)
	config.Db.User = url.Encode(config.Db.User)

	// Create database url
	config.Db.Url = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Name)

	// Validate configuration and set it
	err = config.Validate()
	if err != nil {
		return fmt.Errorf("Init(): config provided is invalid: %w", err)
	}

	configService.config = config

	return nil
}

func (configService configService) Config() *Config {
	return &configService.config
}
