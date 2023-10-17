postgres:
	docker run --name c-simplebank -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:15.2-alpine

createdb:
	docker exec -it c-simplebank createdb --usernalsme=postgres --owner=postgres db_simple_bank

dropdb:
	docker exec -it c-simplebank dropdb db_simple_bank

migrateup:
	 migrate -path platform/migrations/ -database "postgresql://postgres:secret@localhost:5432/db_simple_bank?sslmode=disable" -verbose up

migratedown:
	 migrate -path platform/migrations/ -database "postgresql://postgres:secret@localhost:5432/db_simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test