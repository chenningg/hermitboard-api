package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chenningg/hermitboard-api/db"
	"github.com/chenningg/hermitboard-api/graph/resolver"
	"github.com/chenningg/hermitboard-api/log"
	"github.com/chenningg/hermitboard-api/redis"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/chenningg/hermitboard-api/config"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// Load in logger.
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"
	zl := zerolog.New(os.Stderr).With().Timestamp().Logger()
	var logger logr.Logger = zerologr.New(&zl).WithName("hermitboard-api")

	// Load in configuration.
	configService, err := config.NewConfigService(logger.WithName("config"))
	if err != nil {
		// Log configuration error and panic.
		logger.Error(err, "main(): config could not be loaded")
		os.Exit(1)
	}
	cfg := configService.Config()

	// Open database connection.
	dbService, err := db.NewDbService(cfg.Db, logger.WithName("db"))
	if err != nil {
		// Log database error and panic.
		logger.Error(err, "main(): database setup could not be completed")
		os.Exit(1)
	}

	// Defer closing of database connection till end of main function.
	defer func(dbService *db.DbService) {
		err := dbService.Close()
		if err != nil {
			logger.Error(err, "main(): could not close database properly")
			os.Exit(1)
		}
	}(dbService)

	// Open Redis connection for session storage.
	redisService, err := redis.NewRedisService(cfg.Redis, logger.WithName("redis"))
	if err != nil {
		logger.Error(err, "main(): problem connecting to redis database")
		os.Exit(1)
	}

	// Defer closing of Redis connection till end of main function.
	defer func(redisService *redis.RedisService) {
		err := redisService.Close()
		if err != nil {
			logger.Error(err, "main(): could not close Redis database properly")
			os.Exit(1)
		}
	}(redisService)

	// Create router
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(log.Middleware(logger.WithName("router")))
	router.Use(middleware.Recoverer)

	// Initialize the web server and routes.
	srv := handler.NewDefaultServer(resolver.NewSchema(dbService.Client(), redisService.Client()))

	// Routes for healthcheck, GraphQL playground and API.
	router.Get(
		"/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Healthy! ^-^"))
		},
	)

	// Only show playground if in development.
	if cfg.App.Env == config.Development {
		router.Handle("/playground", playground.Handler("Hermitboard API", "/api"))
	}
	router.Handle("/api", srv)

	// Start the server.
	err = http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, router)
	if err != nil {
		logger.Error(
			err, "main(): server error",
			"host", cfg.Server.Host,
			"port", cfg.Server.Port,
		)
	}
	logger.Info(
		fmt.Sprintf(
			"connect to http://%s:%s/playground for GraphQL playground", cfg.Server.Host, cfg.Server.Port,
		),
		"host", cfg.Server.Host,
		"port", cfg.Server.Port,
	)
}
