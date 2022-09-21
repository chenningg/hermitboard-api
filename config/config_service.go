package config

import (
	"fmt"
	"os"
	"strings"

	envLoader "github.com/caarlos0/env/v6"
	"github.com/chenningg/hermitboard-api/url"
	"github.com/go-logr/logr"
	"github.com/joho/godotenv"
)

type ConfigServicer interface {
	Config() *Config
}

type ConfigService struct {
	config Config
	logger logr.Logger
}

func NewConfigService(logger logr.Logger) (*ConfigService, error) {
	var configService = new(ConfigService)

	// Load in configuration from .env or environment variables.
	configService.logger = logger
	configService.logger.V(2).Info("initializing config service")

	// Load in environment variables.
	if err := configService.loadEnv(); err != nil {
		return configService, fmt.Errorf("NewConfigService(): failed to load environment variables: %v", err)
		configService.logger.Error(err, "NewConfigService(): failed to load environment variables")
	}

	// Unmarshal into config struct.
	err := configService.loadConfig()
	if err != nil {
		return configService, fmt.Errorf("NewConfigService(): failed to load config: %v", err)
		configService.logger.Error(err, "NewConfigService(): failed to load config")
	}

	return configService, nil
}

func (configService *ConfigService) loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("loadEnv(): %v", err)
	}

	// Load environment variable, defaults to development.
	env := strings.ToLower(os.Getenv("HB_APP_ENV"))
	if env == "" {
		os.Setenv("HB_APP_ENV", "Development")
		configService.logger.V(1).Info("no application environment specified, defaulting to development")
	}

	return nil
}

func (configService *ConfigService) loadConfig() error {
	opts := envLoader.Options{Prefix: "HB_"}
	config := Config{}

	// Unmarshal into configuration struct.
	err := envLoader.Parse(&config, opts)
	if err != nil {
		return fmt.Errorf("loadConfig(): failed to parse config provided: %v", err)
	}

	// Clean and encode all URLs
	config.Server.Host = url.Encode(config.Server.Host)

	config.Db.Host = url.Encode(config.Db.Host)
	config.Db.Password = url.Encode(config.Db.Password)
	config.Db.User = url.Encode(config.Db.User)
	config.Db.Name = url.Encode(config.Db.Name)

	// Create database url
	config.Db.Url = fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s", config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port,
		config.Db.Name,
	)

	// Validate configuration and set it
	err = config.Validate()
	if err != nil {
		return fmt.Errorf("loadConfig(): config file provided is invalid: %v", err)
	}

	configService.config = config

	return nil
}

func (configService *ConfigService) Config() *Config {
	return &configService.config
}
