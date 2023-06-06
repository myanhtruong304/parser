package explorerData

import (
	"github.com/myanhtruong304/parser/package/config"
)

type Explorer struct {
	cfg config.Config
}

type Chain struct {
	ExplorerUri string
	APIKey      string
}

func NewExplorerHandler(cfg config.Config) (*Explorer, error) {
	return &Explorer{
		cfg: cfg,
	}, nil
}
func (e *Explorer) ChainSelect(chain string) Chain {
	switch c := chain; c {
	case "bsc":
		return Chain{
			ExplorerUri: e.cfg.BscScanUri,
			APIKey:      e.cfg.BscScanAPIKey,
		}
	case "eth":
		return Chain{
			ExplorerUri: e.cfg.EthScanUri,
			APIKey:      e.cfg.EthScanAPIKey,
		}
	default:
		return Chain{
			ExplorerUri: "",
			APIKey:      "",
		}
	}
}
