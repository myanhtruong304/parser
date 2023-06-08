package utils

import (
	"github.com/myanhtruong304/parser/package/config"
	"github.com/myanhtruong304/parser/package/model"
)

func ChainSelect(chain string, cfg config.Config) model.ChainSwitcher {
	switch c := chain; c {
	case "bsc":
		return model.ChainSwitcher{
			ExplorerUri:    cfg.BSC_SCAN_URI,
			ExplorerApiKey: cfg.BSC_SCAN_API_KEY,
			RpcUri:         cfg.BSC_RCP_URI,
			ChainID:        56,
			ChainName:      "Binance Smart Chain Mainnet",
		}
	case "eth":
		return model.ChainSwitcher{
			ExplorerUri:    cfg.ETH_SCAN_URI,
			ExplorerApiKey: cfg.ETH_SCAN_API_KEY,
			RpcUri:         cfg.ETH_RCP_URI,
			ChainID:        1,
			ChainName:      "Ethereum Mainnet",
		}
	default:
		return model.ChainSwitcher{
			ExplorerUri:    "",
			ExplorerApiKey: "",
			RpcUri:         "",
			ChainID:        0,
			ChainName:      chain,
		}
	}
}
