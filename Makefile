# ===================================
# Utilities
# ===================================

.PHONY: go-tidy
go-tidy:
	go mod tidy

# ===================================
# Database development
# ===================================

# Installs SQLC.
.PHONY: install-sqlc
install-sqlc:
	docker pull kjconroy/sqlc

# Generates SQLC queries.
.PHONY: sqlc
sqlc: install-sqlc
	docker run --rm -v "$(CURDIR):/src" -w /src kjconroy/sqlc generate

# Installs EntGo.
.PHONY: install-ent
install-ent:
	go install github.com/crossworth/ent/cmd/ent@v0.10.1-0.20220913165355-0c5e260f5c84

# Creates a new Ent entity.
.PHONY: ent-new
ent-new:
	go run -mod=mod entgo.io/ent/cmd/ent init $(new)

# Runs Ent codegen.
.PHONY: ent
ent:
	go generate ./ent

# Get database schema
.PHONY: ent-schema
ent-schema:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

# Runs Ent Atlas migration.
.PHONY: migrate
migrate:
	go run -mod=mod ent/migrate/main.go $(name)

# ===================================
# GraphQL
# ===================================

.PHONY: install-gqlgen
install-gqlgen:
	go install github.com/99designs/gqlgen@latest

.PHONY: gqlgen
gqlgen: install-gqlgen
	go run github.com/99designs/gqlgen generate

# ===================================
# Server
# ===================================
.PHONY: dev
dev: go-tidy
	go run ./cmd/hermitboard_api/main.go