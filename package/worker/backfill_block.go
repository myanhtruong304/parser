package worker

import (
	"context"
	"fmt"

	db "github.com/myanhtruong304/parser/db/sqlc"
)

func (w *Worker) findMissingBlock() ([]int, error) {
	blocks, err := w.store.GetAllBlock(context.Background())
	if err != nil {
		return nil, nil
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

	missingblock, err := w.findMissingBlock()
	if err != nil {
		fmt.Println("can not find missing block", err)
	}
	q := db.AddBlockParams{}
	for _, block := range missingblock {
		q.BlockNumber = int32(block)
		q.Processed = false
		_, err := w.store.AddBlock(context.Background(), q)
		if err != nil {
			fmt.Println("can not backfill block", err)
		}
	}
}
