include .env
export

DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)
postgres:
	docker run --name postgres12 \
	-p $(DB_PORT):5432 \
	-e POSTGRES_USER=$(DB_USER) \
	-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
	-e POSTGRES_DB=$(DB_NAME) \
	-v go_api_postgres_volume:/var/lib/postgresql/data \
	-d postgres:12-alpine

remove_postgres:
	docker stop postgres12
	docker rm postgres12
	docker volume rm go_api_postgres_volume


migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema

migratenew:
	@name=$(if $(name),$(name),$(shell date +%Y%m%d%H%M%S)); \
	migrate create -ext sql -dir db/migration -seq $$name

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go api/db/sqlc Store

server:
	go run main.go

build:
	go build

copyenv:
	cp .env.sample .env

docs:
	statik -src=./doc/swagger -dest=./doc


.PHONY: sqlc test server build postgres createdb migrateup migreatecreate mock migrateup1 migratedown migratedown1 remove_postgres docs