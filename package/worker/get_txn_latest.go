package worker

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	db "github.com/myanhtruong304/parser/db/sqlc"
)

func (w *Worker) GetLatestTxn(args ...string) error {
	// Define the number of workers in the pool
	numWorkers := 1

	// Create a channel to receive block numbers that need processing
	blockNumChan := make(chan db.Blocks)

	// Create a wait group to ensure all workers finish before exiting
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start the worker goroutines
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for blockNum := range blockNumChan {
				fmt.Println("Processing block: ", int64(blockNum.BlockNumber), " chain: ", w.chainName)

				block, err := w.rpcHandler.GetBlockData(int64(blockNum.BlockNumber))
				if err != nil {
					if strings.Contains(err.Error(), "429") {
						fmt.Println("failed [GetLatestTxn w.rpcHandler.GetBlockData]: rate limit", w.chainName)
						time.Sleep(20 * time.Second)
						continue
					}
					// fmt.Println("[GetLatestTxn w.rpcHandler.GetBlockData]: not found")
					time.Sleep(1 * time.Second)
					continue
				}

				walletList, err := w.store.GetListWallet(context.Background())
				if err != nil {
					// fmt.Println("can not get list of wallets from db", err)
					continue
				}

				w.GetTxn(walletList, block)

				time.Sleep(1 * time.Second)
			}
		}()
	}

	go func() {
		q := db.GetNotProcessBlockParams{
			Processed: false,
			ChainID:   w.chainID,
		}
		for {
			blockNum, _ := w.store.GetNotProcessBlock(context.Background(), q)

			if len(blockNum) == 0 {
				time.Sleep(1 * time.Second)
				continue
			}

			// Send the block numbers to the channel for processing
			for _, block := range blockNum {
				blockNumChan <- block
			}
		}
	}()

	// Wait for all workers to finish processing
	wg.Wait()

	return nil

}

func (w *Worker) GetTxn(walletList []string, block *types.Block) {
	for _, wallet := range walletList {
		for _, tx := range block.Transactions() {
			from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
			if err != nil {
				fmt.Println(tx.Hash().Hex())
				fmt.Println("[block.Transactions()] ", err)
			}

			if tx.To() == nil {
				continue
			}

			q := db.GetOneTxnParams{
				TxnHash: tx.Hash().Hex(),
				ChainID: w.chainID,
			}

			_, err = w.store.GetOneTxn(context.Background(), q)
			if err == nil {
				continue
			}

			txRcp, err := w.client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				if strings.Contains(err.Error(), "rate-limited") {
					fmt.Println("[w.client.TransactionReceipt]", err)
					time.Sleep(20 * time.Second)
					continue
				}
				fmt.Println("[w.client.TransactionReceipt]", err)
				time.Sleep(1 * time.Second)
				continue
			}

			if from.Hex() == wallet || tx.To().Hex() == wallet {
				txn := db.AddTxnParams{
					WalletAddress:  wallet,
					Chain:          w.chainName,
					ChainID:        w.chainID,
					TxnHash:        tx.Hash().Hex(),
					FromAddress:    from.Hex(),
					ToAddress:      tx.To().Hex(),
					BlockCreatedAt: time.Unix(int64(block.Time()), 0),
					Block:          int32(block.Number().Int64()),
					CreatedAt:      time.Unix(int64(block.Time()), 0),
					Sequence:       0,
					Fee:            tx.Cost().String(),
					Metadata:       tx.Value().String(),
				}

				if from.Hex() == wallet {
					txn.Type = "transfer"
				} else if tx.To().Hex() == wallet {
					txn.Type = "receive"
				} else {
					txn.Type = "undefine"
				}

				if &txRcp.Status == nil {
					break
				}
				if txRcp.Status == types.ReceiptStatusFailed {
					txn.Status = "failed"
				} else {
					txn.Status = "completed"
				}

				_, err := w.store.AddTxn(context.Background(), txn)
				if err != nil {
					fmt.Print("[w.store.AddTxn] ", tx.Hash().Hex(), err)
				}
			}
		}
	}

}
