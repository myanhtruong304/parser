package rpc

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	db "github.com/myanhtruong304/parser/db/sqlc"
)

func (r *RpcHandler) GetBlockData(blockNum int64) (*types.Block, error) {
	block, err := r.client.BlockByNumber(context.Background(), big.NewInt(blockNum))
	if err != nil {
		return nil, err
	}

	q := db.UpdateBlockProcessParams{
		BlockNumber: int32(blockNum),
		ChainID:     int32(r.chainId),
		Processed:   true,
	}
	_, err = r.store.UpdateBlockProcess(context.Background(), q)
	if err != nil {
		return nil, err
	}

	return block, nil
}
