.PHONY : all
all: build

run:
	go run ./internal/cmd/main.go

db-dev:
	docker compose up -d

db-dev-stop:
	docker compose down	
