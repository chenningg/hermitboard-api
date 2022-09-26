//go:build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	// Load in logger.
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"
	zl := zerolog.New(os.Stderr).With().Timestamp().Logger()
	var logger logr.Logger = zerologr.New(&zl).
		WithName("hermitboard-api").
		WithName("entc")

	egql, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./graph/schema/ent.graphql"),
	)
	if err != nil {
		logger.Error(err, "entc: could not create entgql extension")
		os.Exit(1)
	}

	// Entc options.
	opts := []entc.Option{
		entc.Extensions(egql),
		// entc.TemplateDir("./ent/template"),
	}

	// Generate Ent schema.
	err = entc.Generate(
		"./ent/schema", &gen.Config{
			Features: []gen.Feature{gen.FeatureVersionedMigration},
		},
		opts...,
	)
	if err != nil {
		logger.Error(err, "entc: error running ent codegen")
		os.Exit(1)
	}
}
