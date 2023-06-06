package entity

import "github.com/myanhtruong304/parser/package/config"

type Entity struct {
	cfg config.Config
}

func NewEntity(cfg config.Config) Entity {
	return Entity{
		cfg: cfg,
	}
}
