package worker

import (
	"github.com/ethereum/go-ethereum/ethclient"
	db "github.com/myanhtruong304/parser/db/sqlc"
	"github.com/myanhtruong304/parser/package/config"
	explorerData "github.com/myanhtruong304/parser/package/explorer_data"
	"github.com/myanhtruong304/parser/package/rpc"
)

type Worker struct {
	config       *config.Config
	store        db.Store
	numOfWorkers int
	rpcHandler   *rpc.RpcHandler
	client       *ethclient.Client
	explorerData *explorerData.Explorer
}

func NewWorker(config *config.Config, store db.Store, numOfWorkers int, rpcHandler *rpc.RpcHandler, client *ethclient.Client, explorerData *explorerData.Explorer) Worker {
	return Worker{
		config:       config,
		store:        store,
		numOfWorkers: numOfWorkers,
		rpcHandler:   rpcHandler,
		client:       client,
		explorerData: explorerData,
	}
}
