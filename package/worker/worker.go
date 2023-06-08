package worker

import (
	"github.com/ethereum/go-ethereum/ethclient"
	db "github.com/myanhtruong304/parser/db/sqlc"
	"github.com/myanhtruong304/parser/package/config"
	explorerData "github.com/myanhtruong304/parser/package/explorer_data"
	"github.com/myanhtruong304/parser/package/rpc"
	"github.com/myanhtruong304/parser/utils"
)

type Worker struct {
	config       *config.Config
	store        db.Store
	numOfWorkers int
	rpcHandler   *rpc.RpcHandler
	client       *ethclient.Client
	explorerData *explorerData.Explorer
	chainID      int32
	chainName    string
}

func NewWorker(config config.Config, store db.Store, numOfWorkers int, chain string, client *ethclient.Client, rpcHandler *rpc.RpcHandler, exploreHandler *explorerData.Explorer) Worker {
	return Worker{
		config:       &config,
		store:        store,
		numOfWorkers: numOfWorkers,
		rpcHandler:   rpcHandler,
		client:       client,
		explorerData: exploreHandler,
		chainName:    utils.ChainSelect(chain, config).ChainName,
		chainID:      int32(utils.ChainSelect(chain, config).ChainID),
	}
}
