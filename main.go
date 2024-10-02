package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mariotoribi0/simplebank/api"
	db "github.com/mariotoribi0/simplebank/db/sqlc"
	"github.com/mariotoribi0/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	mainDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(mainDB)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
