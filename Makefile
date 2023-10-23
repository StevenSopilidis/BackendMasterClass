DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name bank_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createDb:
	docker exec -it bank_postgres createdb --username=root --owner=root bank_postgres

dropDb:
	docker exec -it bank_postgres dropdb bank_postgres

migrateUp:
	migrate -path db/migrations -database "${DB_URL}" -verbose up

migrateDown:
	migrate -path db/migrations -database "${DB_URL}" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:	
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/StevenSopilidis/BackendMasterClass/db/sqlc Store

db_docks:
	dbdocs build docs/db.dml

db_schema:
	dbml2sql --postgres -o docs/schema.sql doc/db.dbml

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

.PHONY: postgres createDb dropDb migrateUp migrateDown sqlc test server mock db_docks db_schema proto