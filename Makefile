.PHONY: migrate-up, migrate-up-test

migrate-up:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations
migrate-up-test:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_test
run:
	go run cmd/sso/main.go --config=./config/local.yaml