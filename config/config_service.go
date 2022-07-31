package config

import (
	"fmt"
	"net/url"
	"os"

	envLoader "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type IConfigService interface {
	Load() *IConfigService
	Config() *Config
	SetConfig()
}

func NewConfigService() (*configService, error) {
	var cfgService = new(configService)

	// Load in configuration from .env or environment variables.
	err := cfgService.Load()
	if err != nil {
		return nil, fmt.Errorf("NewConfigService(): %w", err)
	}

	return cfgService, nil
}

type configService struct {
	config Config
}

func (configService *configService) Load() error {
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
		return fmt.Errorf("Load(): failed to load config: %w", err)
	}

	// Clean and encode all URLs
	config.Db.Password = urlEncode(config.Db.Password)
	config.Db.User = urlEncode(config.Db.User)

	// Create database url
	config.Db.Url = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Name)

	// Validate configuration
	if err := config.Validate(); err != nil {
		return fmt.Errorf("Load(): config validation did not pass: %w", err)
	}

	configService.SetConfig(config)

	return nil
}

func (configService configService) Config() *Config {
	return &configService.config
}

func (configService *configService) SetConfig(config Config) {
	configService.config = config
}

func urlEncode(str string) string {
	return url.PathEscape(str)
}
