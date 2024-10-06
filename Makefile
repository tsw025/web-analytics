DB_URL="$(DATABASE_URL)"

.PHONY: migrate-up migrate-down migrate-new

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	@read -p "Enter number of migrations to rollback: " n; \
	migrate -path ./migrations -database "$(DB_URL)" down $$n

migrate-down-all:
	migrate -path ./migrations -database "$(DB_URL)" down -all

migrate-new:
	@read -p "Enter migration name: " name; \
	migrate create -seq -ext sql -dir migrations "$$name"
