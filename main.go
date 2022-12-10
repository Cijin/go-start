package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/cijin/go-start/api"
	db "github.com/cijin/go-start/db/sqlc"
	"github.com/cijin/go-start/utils"
)

func main() {
	confg, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Error reading env config: ", err)
	}

	conn, err := sql.Open(confg.DBDriver, confg.DBSource)
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(confg.ServerAddress)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
