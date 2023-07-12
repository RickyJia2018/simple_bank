package main

import (
	"database/sql"
	"log"

	"github.com/RickyJia2018/simple_bank/api"
	db "github.com/RickyJia2018/simple_bank/db/sqlc"
	"github.com/RickyJia2018/simple_bank/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LaodConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
