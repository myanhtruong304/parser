package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	db "github.com/myanhtruong304/parser/db/sqlc"
	"github.com/myanhtruong304/parser/package/config"
	"github.com/myanhtruong304/parser/utils"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to load config", err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal("can not connect to database", err)
	}
	store := db.NewStore(conn)

	err = utils.FeedWalletsTable(config, store, context.Background())
}
