package rpc

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	db "github.com/myanhtruong304/parser/db/sqlc"
	"github.com/myanhtruong304/parser/package/config"
	"github.com/myanhtruong304/parser/utils"
)

type RpcHandler struct {
	cfg     *config.Config
	client  *ethclient.Client
	store   db.Store
	chain   string
	chainId int
}

func NewRcpHandle(cfg config.Config, store db.Store, chain string) (*RpcHandler, error) {
	client, err := ethclient.Dial(utils.ChainSelect(chain, cfg).RpcUri)
	if err != nil {
		fmt.Println("[NewRcpHandle] Can not dial client ", chain)
	}

	return &RpcHandler{
		cfg:     &cfg,
		client:  client,
		store:   store,
		chain:   chain,
		chainId: utils.ChainSelect(chain, cfg).ChainID}, nil
}
