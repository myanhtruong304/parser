package entity

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	db "github.com/myanhtruong304/parser/db/sqlc"
	"github.com/myanhtruong304/parser/package/model"
	"github.com/myanhtruong304/parser/utils"
)

func (e *Entity) CreateWallet(c *gin.Context, req model.CreateWalletRequest) (*string, error) {
	a := db.CreateWalletParams{
		WalletAddress: req.WalletAddress,
		CreatedBlock:  sql.NullString{String: req.CreatedBlock, Valid: true},
	}

	createW, err := e.repo.CreateWallet(c, a)
	if err != nil {
		return nil, err
	}

	return &createW.WalletAddress, nil
}

func (e *Entity) GetWalletTransaction(c *gin.Context, req model.GetWalletTransactionRequest) ([]db.Transactions, error) {
	chain := utils.ChainSelect(req.Chain, e.cfg)
	q := db.GetAllTxnParams{
		WalletAddress: req.Address,
		ChainID:       int32(chain.ChainID),
	}

	txn, err := e.repo.GetAllTxn(c, q)
	if err != nil {
		return nil, err
	}

	return txn, nil
}
