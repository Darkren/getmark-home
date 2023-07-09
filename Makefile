.PHONY: requirements generate
.PHONY: add-migration migrate

MIGRATIONS_PATH=./migrations

requirements:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/vektra/mockery/v2@v2.30.1

add-migration:
	MIGRATIONS_PATH=$(MIGRATIONS_PATH) bash ./add-migration.sh

generate:
	go generate ./...
