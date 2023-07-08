.PHONY: requirements
.PHONY: add-migration migrate

MIGRATIONS_PATH=./migrations

requirements:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

add-migration:
	MIGRATIONS_PATH=$(MIGRATIONS_PATH) bash ./add-migration.sh


