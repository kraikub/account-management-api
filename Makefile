.PHONY : all
all: build

run:
	go run ./api/v1/internal/cmd/main.go

run-production:
	export KRAIKUB_ENV=production && \
	export KRAIKUB_SERVER_NAME=account-management-api && \
	export KRAIKUB_SERVER_PORT=3061 && \
	make run

run-production-ps:
	sh ./run-production-ps.sh

image:
	docker build -t kraikub/account-management-api -f ./build/docker/Dockerfile .

db-dev:
	docker compose up -d

db-dev-stop:
	docker compose down	

