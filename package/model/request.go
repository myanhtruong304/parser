package model

type CreateWalletRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	CreatedBlock  string `json:"created_block,omitempty"`
}

type GetWalletTransactionRequest struct {
	Chain   string `json:"id" binding:"required"`
	Address string `json:"address" binding:"required"`
}
