package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chenningg/hermitboard-api/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

func main() {
	// Load in configuration
	configService, err := config.NewConfigService()
	if err != nil {
		log.Fatal().Err(err).Msg("config error")
	}

	config := configService.Config()

	// Open database pool connection
	dbPool, err := pgxpool.Connect(context.Background(), config.Db.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	var greeting string
	err = dbPool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
