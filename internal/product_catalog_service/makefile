include .env
export

MIGRATION_DIR=internal/infra/db/migrations/sql

# Get the current timestamp
TIMESTAMP=$(shell date +%Y%m%d%H%M%S)

.PHONY: db_init db_migrate db_create_sql

db_init:
	DATABASE_URL=$(DATABASE_URL) go run main.go db init
db_seed:
	DATABASE_URL=$(DATABASE_URL) go run main.go db seed
db_migrate:
	DATABASE_URL=$(DATABASE_URL) go run main.go db migrate
db_create_sql:
	DATABASE_URL=$(DATABASE_URL) go run main.go db create_sql $(name)
app_start:
	DATABASE_URL=$(DATABASE_URL) go run main.go start
db_create_migration:
	@if [ -z "$(MIGRATION_NAME)" ]; then \
		read -p "Enter migration name: " MIGRATION_NAME; \
	fi; \
	UP_FILE=$(MIGRATION_DIR)/$(TIMESTAMP)_$$MIGRATION_NAME.up.sql; \
	DOWN_FILE=$(MIGRATION_DIR)/$(TIMESTAMP)_$$MIGRATION_NAME.down.sql; \
	echo "Creating migration files..."; \
	mkdir -p $(MIGRATION_DIR); \
	touch $$UP_FILE $$DOWN_FILE; \
	echo "-- Migration up script for $$MIGRATION_NAME" > $$UP_FILE; \
	echo "-- Migration down script for $$MIGRATION_NAME" > $$DOWN_FILE; \
	echo "Created $$UP_FILE and $$DOWN_FILE"