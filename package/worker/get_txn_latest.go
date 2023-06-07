package worker

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	db "github.com/myanhtruong304/parser/db/sqlc"
)

func (w *Worker) GetLatestTxn(args ...string) error {
	for {
		blockNum, _ := w.store.GetNotProcessBlock(context.Background(), false)
		if blockNum.BlockNumber == 0 {
			continue
		}
		block, err := w.rpcHandler.GetBlockData(int64(blockNum.BlockNumber))
		if err != nil {
			if strings.Contains(err.Error(), "rate-limited") {
				fmt.Println("rate-limited")
				time.Sleep(30 * time.Second)
				continue
			}
			fmt.Println("[GetLatestTxn w.rpcHandler.GetBlockData]", err)
			time.Sleep(1 * time.Second)
			continue
		}

		walletList, err := w.store.GetListWallet(context.Background())
		if err != nil {
			fmt.Println("can not get list of wallets from db", err)
			continue
		}

		walletCh := make(chan string, 100)

		go func() {
			defer close(walletCh)

			for _, wallet := range walletList {
				walletCh <- wallet
			}
		}()
		w.GetTxn(walletList, block, walletCh)

		time.Sleep(1 * time.Second)
	}
}

func (w *Worker) GetTxn(walletList []string, block *types.Block, walletCh chan string) {
	for wallet := range walletCh {
		for _, tx := range block.Transactions() {
			from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
			if err != nil {
				fmt.Println("[block.Transactions()] ", err)
			}

			if &from == nil || tx.To() == nil {
				continue
			}

			_, err = w.store.GetOneTxn(context.Background(), tx.Hash().Hex())
			if err == nil {
				continue
			}
			if from.Hex() == wallet || tx.To().Hex() == wallet {
				txn := db.AddTxnParams{
					WalletAddress: wallet,
					Chain:         tx.ChainId().String(),
					ChainID:       271,
					TxnHash:       tx.Hash().Hex(),
					FromAdd:       from.Hex(),
					ToAdd:         tx.To().Hex(),
				}

				_, err := w.store.AddTxn(context.Background(), txn)
				if err != nil {
					log.Panic("[w.store.AddTxn] ", err)
				}
			}
		}
	}

}
