.PHONY dev_db_migrate_up:
	sql-migrate up -env="development"

.PHONY dev:
	go run ./cmd/hermitboard_api/main.go