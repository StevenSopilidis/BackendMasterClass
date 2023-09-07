postgres:
	docker run --name bank_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createDb:
	docker exec -it bank_postgres createdb --username=root --owner=root bank_postgres

dropDb:
	docker exec -it bank_postgres dropdb bank_postgres

migrateUp:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/bank_postgres?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/bank_postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createDb dropDb migrateUp migrateDown sqlc test