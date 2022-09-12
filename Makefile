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
	go install entgo.io/ent/cmd/ent@latest

# Creates a new Ent entity.
.PHONY: ent-new
ent-new: install-ent 
	go run -mod=mod entgo.io/ent/cmd/ent init $(new)

# Runs Ent codegen.
.PHONY: ent
ent: install-ent 
	go generate ./ent


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