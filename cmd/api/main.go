package main

import (
	"context"
	"database/sql"
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
	"github.com/myanhtruong304/parser/pb"
)

type Server struct {
	pb.UnimplementedPay68Server
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

	client, err := ethclient.Dial(config.BSC_RCP_URI)

	rpcHandler, err := rpc.NewRcpHandle(&config, client, store)
	if err != nil {
		log.Panic("main [rpc.NewRcpHandle]", err)
	}

	explorerHandler, err := explorerData.NewExplorerHandler(config)
	if err != nil {
		log.Panic("main [explorerData.NewExplorerHandler]", err)
	}

	worker := worker.NewWorker(&config, store, 10, rpcHandler, client, explorerHandler)
	go func() {
		err := server.Start(config.SERVER_ADDRESS)
		if err != nil {
			log.Fatal("can not start server", err)
		}
	}()

	go worker.GetLatestblock()
	go worker.GetLatestTxn()
	go worker.BackfillMissingBlocks()

	select {}

}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
