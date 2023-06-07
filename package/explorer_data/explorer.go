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
			ExplorerUri: e.cfg.BSC_SCAN_URI,
			APIKey:      e.cfg.BSC_SCAN_API_KEY,
		}
	case "eth":
		return Chain{
			ExplorerUri: e.cfg.ETH_SCAN_URI,
			APIKey:      e.cfg.ETH_SCAN_API_KEY,
		}
	default:
		return Chain{
			ExplorerUri: "",
			APIKey:      "",
		}
	}
}
