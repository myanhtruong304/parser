package utils

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	db "github.com/myanhtruong304/parser/db/sqlc"
	"github.com/myanhtruong304/parser/package/config"
	"github.com/myanhtruong304/parser/package/model"
)

func FeedWalletsTable(config config.Config, store db.Store, c context.Context) error {
	walletList := GenerateRandomWallets(100)

	for _, wallet := range walletList {
		q := db.CreateWalletParams{
			WalletAddress: wallet.WalletAddress,
			CreatedBlock:  sql.NullString{String: wallet.CreationBlock, Valid: true},
		}
		_, err := store.CreateWallet(c, q)
		if err != nil {
			return fmt.Errorf("addWallet: %v", err)
		}
		fmt.Printf("Seed wallet %s\n", wallet.WalletAddress)
	}
	return nil
}

func GenerateRandomWallets(numWallets int) []model.Wallet {
	var wallets []model.Wallet
	currentTime := time.Now().Unix()

	for i := 0; i < numWallets; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal("Failed to generate private key:", err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("Failed to convert public key to ECDSA")
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		// Generate a random timestamp within the last 24 hours
		randomTimestamp := currentTime - int64(rand.Intn(24*60*60))

		wallet := model.Wallet{
			WalletAddress: address,
			CreationBlock: strconv.Itoa(int(randomTimestamp)),
		}

		wallets = append(wallets, wallet)
	}

	return wallets
}
