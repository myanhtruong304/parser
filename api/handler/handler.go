package handler

import (
	"context"

	"github.com/myanhtruong304/parser/api/entity"
	db "github.com/myanhtruong304/parser/db/sqlc"
	"github.com/myanhtruong304/parser/package/config"
)

type Handler struct {
	entity entity.Entity
	cfg    config.Config
}

func NewHandler(cfg config.Config, c context.Context, store db.Store) *Handler {
	return &Handler{
		entity: entity.NewEntity(cfg, store),
		cfg:    cfg,
	}
}
