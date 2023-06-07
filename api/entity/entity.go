package entity

import (
	db "github.com/myanhtruong304/parser/db/sqlc"
	"github.com/myanhtruong304/parser/package/config"
)

type Entity struct {
	cfg  config.Config
	repo db.Store
}

func NewEntity(cfg config.Config, store db.Store) Entity {
	return Entity{
		cfg:  cfg,
		repo: store,
	}
}
