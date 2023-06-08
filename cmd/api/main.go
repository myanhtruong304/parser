package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/myanhtruong304/parser/api/handler"
	routes "github.com/myanhtruong304/parser/api/route"
	db "github.com/myanhtruong304/parser/db/sqlc"
	"github.com/myanhtruong304/parser/package/config"
	explorerData "github.com/myanhtruong304/parser/package/explorer_data"
	"github.com/myanhtruong304/parser/package/rpc"
	"github.com/myanhtruong304/parser/package/worker"
	"github.com/myanhtruong304/parser/utils"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

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
	h := handler.NewHandler(config, context.Background(), store)
	r := gin.Default()
	routes := routes.NewRoute(&config, r, h)

	server := &Server{
		store:  store,
		router: &routes,
	}

	go func() {
		err := server.Start(config.SERVER_ADDRESS)
		if err != nil {
			log.Fatal("can not start server", err)
		}
	}()

	go chainProcess("bsc", &config, store, 2)
	go chainProcess("eth", &config, store, 2)
	select {}

}

func chainProcess(chain string, config *config.Config, store db.Store, numOfWorkers int) {
	client, err := ethclient.Dial(utils.ChainSelect(chain, *config).RpcUri)
	if err != nil {
		fmt.Println("[NewRcpHandle] Can not dial client ", chain)
	}

	rpcHandler, err := rpc.NewRcpHandle(*config, store, chain)
	if err != nil {
		log.Panic("main [rpc.NewRcpHandle]", err)
	}

	explorerHandler, err := explorerData.NewExplorerHandler(*config, chain)
	if err != nil {
		log.Panic("main [explorerData.NewExplorerHandler]", err)
	}

	worker := worker.NewWorker(*config, store, 10, chain, client, rpcHandler, explorerHandler)

	go worker.GetLatestblock()
	go worker.GetLatestTxn()
	go worker.BackfillMissingBlocks()
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
