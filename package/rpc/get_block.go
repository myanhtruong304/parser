package rpc

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
)

func (r *RpcHandler) GetBlockData(blockNum int64) (*types.Block, error) {
	block, err := r.client.BlockByNumber(context.Background(), big.NewInt(blockNum))
	if err != nil {
		if strings.Contains(err.Error(), "429 Too Many Requests") {
			return nil, fmt.Errorf("rate-limited: %s", err.Error())
		}
		panic(err)
	}
	return block, nil
}
