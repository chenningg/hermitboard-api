# ===================================
# Utilities
# ===================================

# Generates SQLC functions to marshal database values into structs.
.PHONY: sqlc
sqlc:
	docker pull kjconroy/sqlc
	docker run --rm -v "${PWD}/:/src" -w /src kjconroy/sqlc generate

.PHONY: gqlgen
gqlgen:
	go run github.com/99designs/gqlgen generate

# ===================================
# Development
# ===================================
.PHONY: dev
dev:
	go mod tidy
	go run ./cmd/hermitboard_api/main.go