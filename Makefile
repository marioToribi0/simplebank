postgres:
	sudo docker run --name simplebank_postgresql -p 5400:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

deletepostgres:
	sudo docker stop simplebank_postgresql
	sudo docker rm simplebank_postgresql

createdb:
	sudo docker exec -it simplebank_postgresql createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it simplebank_postgresql dropdb simple_bank

migrateup:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5400/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5400/simple_bank?sslmode=disable" --verbose down -all

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb deletepostgres  migrateup migratedown test