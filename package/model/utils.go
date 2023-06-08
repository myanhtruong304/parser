package model

type Wallet struct {
	WalletAddress string `json:"wallet_address"`
	CreationBlock string `json:"created_block"`
}

type ChainSwitcher struct {
	ExplorerUri    string
	ExplorerApiKey string
	RpcUri         string
	ChainID        int
	ChainName      string
}
