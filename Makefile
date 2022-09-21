# ===================================
# Constants
# ===================================
MIGRATIONS_DIR = file://ent/migrate/migrations

# ===================================
# Database development
# ===================================

# Installs EntGo.
.PHONY: install-ent
install-ent:
	go install entgo.io/ent/cmd/ent@latest

# Creates a new Ent entity.
.PHONY: ent-new
ent-new:
	go run -mod=mod entgo.io/ent/cmd/ent init $(name)

# Runs Ent codegen.
.PHONY: ent-gen
ent-gen:
	go generate ./ent

# Get database schema
.PHONY: ent-schema
ent-schema:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

# Installs Atlas migration
.PHONY: install-atlas
install-atlas:
	go install ariga.io/atlas/cmd/atlas@latest

# Runs Ent Atlas migration.
.PHONY: atlas-gen
atlas-gen:
	go run -mod=mod ent/migrate/main.go $(name)

# Lints Atlas migration.
.PHONY: atlas-lint
atlas-lint: install-atlas
	atlas migrate lint \
	--dev-url=docker://postgres?search_path=public \
	--dir=$(MIGRATIONS_DIR) \
	--dir-format=atlas \
	--latest=1

# Creates a new manual Atlas migration.
.PHONY: atlas-new
atlas-new: install-atlas
	atlas migrate new $(name) \
	--dir=$(MIGRATIONS_DIR)

# Rehashes the migration directory.
.PHONY: atlas-hash
atlas-hash: install-atlas
	atlas migrate hash --force \
	--dir=$(MIGRATIONS_DIR)

# Get the status of migrations.
.PHONY: atlas-status
atlas-status: install-atlas
	atlas migrate status \
	--dir=$(MIGRATIONS_DIR) \
	--url=$(url)

# Apply an Atlas migration.
.PHONY: atlas-apply
atlas-apply:
	atlas migrate apply \
	--dir=$(MIGRATIONS_DIR) \
	--url=$(url) \
	--allow-dirty

# Applies an Atlas dry run migration.
.PHONY: atlas-apply-dry
atlas-apply-dry:
	atlas migrate apply \
	--dir=$(MIGRATIONS_DIR) \
	--url=$(url) \
	--allow-dirty \
	--dry-run

# ===================================
# GraphQL
# ===================================

.PHONY: install-gqlgen
install-gqlgen:
	go install github.com/99designs/gqlgen@latest

.PHONY: gqlgen-gen
gqlgen: install-gqlgen
	go run github.com/99designs/gqlgen generate

# ===================================
# Server
# ===================================
.PHONY: dev
dev:
	go run ./cmd/hermitboard_api/main.go