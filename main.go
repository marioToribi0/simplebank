package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mariotoribi0/simplebank/api"
	db "github.com/mariotoribi0/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5400/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:9999"
)

func main() {
	mainDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(mainDB)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
