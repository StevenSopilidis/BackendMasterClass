package main

import (
	"database/sql"
	"log"

	"github.com/StevenSopilidis/BackendMasterClass/api"
	db "github.com/StevenSopilidis/BackendMasterClass/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/bank_postgres?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Could not connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Could not start server")
	}
}
