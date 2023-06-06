package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/myanhtruong304/parser/package/config"
	explorerData "github.com/myanhtruong304/parser/package/explorer_data"
	"github.com/myanhtruong304/parser/package/rpc"
)

type Server struct {
	router *gin.Engine
}

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	client, err := ethclient.Dial(cfg.BscRPCUri)
	if err != nil {
		panic(err)
	}

	rpcHandler, err := rpc.NewRcpHandle(&cfg, client)
	if err != nil {
		panic(err)
	}

	explorerHandler, err := explorerData.NewExplorerHandler(cfg)
	if err != nil {
		panic(err)
	}

	const (
		module    = "block"
		action    = "getblocknobytime"
		timestamp = "1677603600"
		closest   = "before"
	)

	blockAtTimestamp, err := explorerHandler.GetBlockAtTimestamp(timestamp)

	// Fetch the latest block number
	latestBlockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	blockCh := make(chan int64, 100)

	go func() {
		defer close(blockCh)

		block, err := strconv.Atoi(blockAtTimestamp.Result)
		if err != nil {
			panic(err)
		}

		for i := float64(block); i <= float64(latestBlockNumber); i++ {
			blockCh <- int64(i)
		}
	}()

	numOfWorkers := 10

	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for blockNum := range blockCh {
				for {
					block, err := rpcHandler.GetBlockData(blockNum)
					if err != nil {
						if strings.Contains(err.Error(), "rate-limited") {
							time.Sleep(30 * time.Second)
							continue
						}

						log.Panic(err)
					}
					for _, tx := range block.Transactions() {
						from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
						if err != nil {
							log.Panic(err)
						}

						if from.Hex() == "0x08d0b5378734b3ddDBa94dd3D53b141Ea762423F" || tx.To().Hex() == "0x08d0b5378734b3ddDBa94dd3D53b141Ea762423F" {
							pp.Println(tx.Hash().Hex())
							break
						} else {
							fmt.Println("searching")
						}
					}
				}
			}
		}()
	}
	wg.Wait()
}
