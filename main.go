package main

import (
	"database/sql"
	"log"

	"github.com/edwins-leonardi/simplebank/api"
	db "github.com/edwins-leonardi/simplebank/db/sqlc"
	"github.com/edwins-leonardi/simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddr)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
