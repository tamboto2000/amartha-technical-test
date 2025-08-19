new-migration:
	@read -p "Enter migration name: " name; \
	goose -dir ./internal/database/migrations create $$name sql

new-seed:
	@read -p "Enter seed name: " name; \
	goose -dir ./internal/database/seeds create $$name sql

build:
	@go build -o amartha ./cmd/