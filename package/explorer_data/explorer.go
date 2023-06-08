package explorerData

import (
	"github.com/myanhtruong304/parser/package/config"
)

type Explorer struct {
	cfg   config.Config
	chain string
}

func NewExplorerHandler(cfg config.Config, chain string) (*Explorer, error) {
	return &Explorer{
		cfg:   cfg,
		chain: chain,
	}, nil
}
