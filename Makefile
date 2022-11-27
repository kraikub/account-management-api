.PHONY : all
all: build

run:
	go run ./api/v1/cmd/main.go

image:
	docker build -t kraikub/account-management-api -f ./build/docker/Dockerfile .

db-dev:
	docker compose up -d

db-dev-stop:
	docker compose down	
