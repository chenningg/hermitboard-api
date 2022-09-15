//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	egql, err := entgql.NewExtension()
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	},
		entc.Extensions(egql))
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
