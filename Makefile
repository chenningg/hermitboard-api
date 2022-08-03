.PHONY: dev_db_migrate_up
dev_db_migrate_up:
	sql-migrate up -env="development"

.PHONY: dev
dev:
	go run ./cmd/hermitboard_api/main.go