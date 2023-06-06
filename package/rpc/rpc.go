package rpc

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/myanhtruong304/parser/package/config"
)

type RpcHandler struct {
	cfg    *config.Config
	client *ethclient.Client
}

func NewRcpHandle(cfg *config.Config, client *ethclient.Client) (*RpcHandler, error) {
	return &RpcHandler{
		cfg:    cfg,
		client: client}, nil
}
