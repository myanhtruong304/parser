package worker

import (
	"context"
	"fmt"
	"strconv"
	"time"

	db "github.com/myanhtruong304/parser/db/sqlc"
)

func (w *Worker) GetLatestblock(args ...string) error {
	for {
		timestamp := strconv.Itoa(int(time.Now().Unix()))
		latestBlockNumber, err := w.explorerData.GetBlockAtTimestamp(timestamp)
		if err != nil {
			fmt.Println("[GetLatestblock] can not get latest block")
			continue
		}
		blockNum, _ := strconv.Atoi(latestBlockNumber.Result)
		q := db.GetOneBlockParams{
			BlockNumber: int32(blockNum),
			ChainID:     w.chainID,
		}
		_, err = w.store.GetOneBlock(context.Background(), q)
		if err == nil {
			continue
		}

		q2 := db.AddBlockParams{
			BlockNumber: int32(blockNum),
			Processed:   false,
			ChainID:     w.chainID,
		}

		_, err = w.store.AddBlock(context.Background(), q2)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Millisecond)

		// fmt.Println("Latest block ", latestBlockNumber.Result)
	}
}
