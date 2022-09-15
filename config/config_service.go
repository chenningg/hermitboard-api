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

type ConfigService interface {
	Config() *Config
}

type configService struct {
	config Config
	logger logr.Logger
}

func NewConfigService(logger logr.Logger) (*configService, error) {
	var configService = new(configService)

	// Load in configuration from .env or environment variables.
	configService.logger = logger
	configService.logger.V(2).Info("initializing config service")

	// Load in environment variables.
	if err := configService.loadEnv(); err != nil {
		return nil, fmt.Errorf("NewConfigService(): failed to load environment variables: %v", err)
	}

	// Unmarshal into config struct.
	err := configService.loadConfig()
	if err != nil {
		return nil, fmt.Errorf("NewConfigService(): failed to load config: %v", err)
	}

	return configService, nil
}

func (configService *configService) loadEnv() error {
	configService.logger.V(2).Info("loading environment variables")

	// Load environment variable, defaults to development.
	env := strings.ToLower(os.Getenv("HB_APP_ENV"))
	if env == "" {
		env = "development"
		configService.logger.V(1).Info("no application environment specified, defaulting to development")
	}

	// Load any local env files
	err := godotenv.Load(".env." + env + ".local")
	if err == nil {
		return nil
	}

	// Also try to load any .env files with appenv suffix.
	err = godotenv.Load(".env." + env)
	if err == nil {
		return nil
	}

	// Finally, if all else fails load .env file.
	err = godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func (configService *configService) loadConfig() error {
	configService.logger.V(2).Info("loading config")
	opts := envLoader.Options{Prefix: "HB_"}
	config := Config{}

	// Unmarshal into configuration struct.
	err := envLoader.Parse(&config, opts)
	if err != nil {
		return fmt.Errorf("LoadConfig(): failed to parse config provided: %v", err)
	}

	// Clean and encode all URLs
	config.Db.Password = url.Encode(config.Db.Password)
	config.Db.User = url.Encode(config.Db.User)

	// Create database url
	config.Db.Url = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Name)

	// Validate configuration and set it
	err = config.Validate()
	if err != nil {
		return fmt.Errorf("LoadConfig(): config file provided is invalid: %v", err)
	}

	configService.config = config

	return nil
}

func (configService configService) Config() *Config {
	return &configService.config
}
