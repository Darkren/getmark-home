.PHONY: requirements generate
.PHONY: add-migration migrate
.PHONY: build docker-build run

MIGRATIONS_PATH=./migrations

requirements:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/vektra/mockery/v2@v2.30.1

add-migration:
	MIGRATIONS_PATH=$(MIGRATIONS_PATH) bash ./add-migration.sh

generate:
	go generate ./...

build:
	rm -rf bin
	mkdir bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=netgo -o ./bin/getmark-home ./cmd/api

docker-build:
	docker build -t getmark-home -f ./cmd/api/Dockerfile .
	docker build -t migrations -f ./migrations/Dockerfile .

run: build docker-build
	docker compose up

