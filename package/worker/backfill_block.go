package worker

import (
	"context"
	"fmt"
	"time"

	db "github.com/myanhtruong304/parser/db/sqlc"
)

func (w *Worker) findMissingBlock(chain string) ([]int, error) {
	blocks, err := w.store.GetAllBlock(context.Background(), w.chainID)

	if err != nil {
		return nil, err
	}

	if len(blocks) < 1 {
		return nil, nil
	}

	start := int(blocks[0])
	n := len(blocks) + int(blocks[0]) - 1
	missingNumbers := []int{}

	// Create a map to mark the presence of each number
	presence := make(map[int]bool)
	for _, num := range blocks {
		presence[int(num)] = true
	}

	// Iterate from 1 to n and check for missing numbers
	for i := start; i <= n; i++ {
		if !presence[i] {
			missingNumbers = append(missingNumbers, i)
		}
	}

	return missingNumbers, nil
}

func (w *Worker) BackfillMissingBlocks() {
	for {
		missingblock, err := w.findMissingBlock(w.chainName)
		// fmt.Println("missing block", missingblock)
		if err != nil {
			fmt.Println("can not find missing block", err)
		}
		q := db.AddBlockParams{}
		for _, block := range missingblock {
			q.BlockNumber = int32(block)
			q.Processed = false
			q.ChainID = w.chainID
			_, err := w.store.AddBlock(context.Background(), q)
			if err != nil {
				fmt.Println("can not backfill block", err)
			}
		}
		time.Sleep(2 * time.Second)
	}
}
