# Makefile для создания миграций

DB_DSN := postgres://postgres:postgres@localhost:5432/main?sslmode=disable
MIGRATE := migrate -path ./migrations -database "$(DB_DSN)"

migrate-new:
	$(MIGRATE) create -ext sql -dir ./migrations $(NAME)

migrate-down:
	$(MIGRATE) down

migrate:
	$(MIGRATE) up

run:
	go run cmd/app/main.go