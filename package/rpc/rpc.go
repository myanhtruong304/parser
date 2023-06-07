package rpc

import (
	"github.com/ethereum/go-ethereum/ethclient"
	db "github.com/myanhtruong304/parser/db/sqlc"
	"github.com/myanhtruong304/parser/package/config"
)

type RpcHandler struct {
	cfg    *config.Config
	client *ethclient.Client
	store  db.Store
}

func NewRcpHandle(cfg *config.Config, client *ethclient.Client, store db.Store) (*RpcHandler, error) {
	return &RpcHandler{
		cfg:    cfg,
		client: client,
		store:  store}, nil
}
