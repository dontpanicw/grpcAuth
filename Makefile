.PHONY: migrations, migrations_test

migrations:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations
migrations_test:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_test
start_app:
	go run cmd/sso/main.go --config=./config/local.yaml