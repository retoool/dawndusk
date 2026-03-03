.PHONY: build run test migrate-up migrate-down docker-up docker-down

build:
	go build -o bin/server cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	go test -v ./...

migrate-up:
	migrate -path internal/infrastructure/database/migrations -database "postgresql://postgres:password@localhost:5432/dawndusk?sslmode=disable" up

migrate-down:
	migrate -path internal/infrastructure/database/migrations -database "postgresql://postgres:password@localhost:5432/dawndusk?sslmode=disable" down

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

install-deps:
	go mod download
	go mod tidy
