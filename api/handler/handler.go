package handler

import (
	"context"

	"github.com/myanhtruong304/parser/api/entity"
	"github.com/myanhtruong304/parser/package/config"
)

type Handler struct {
	entity entity.Entity
	cfg    config.Config
}

func NewHandler(cfg config.Config, c context.Context) *Handler {
	return &Handler{
		entity: entity.NewEntity(cfg),
		cfg:    cfg,
	}
}
