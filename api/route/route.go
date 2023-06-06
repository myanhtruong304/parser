package routes

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/myanhtruong304/parser/api/handler"
	"github.com/myanhtruong304/parser/package/config"
)

func setupCORS(r *gin.Engine, cfg *config.Config) {
	corsOrigins := strings.Split(cfg.ApiServer, ";")
	r.Use(func(c *gin.Context) {
		cors.New(
			cors.Config{
				AllowOrigins: corsOrigins,
				AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders: []string{
					"Origin", "Host", "Content-Type", "Content-Length", "Accept-Encoding", "Accept-Language", "Accept",
					"X-CSRF-Token", "Authorization", "X-Requested-With", "X-Access-Token",
				},
				AllowCredentials: true,
			},
		)(c)
	})
}

func NewRoute(r *gin.Engine, h *handler.Handler) gin.Engine {
	v1 := r.Group("/api/v1")
	groupTxn := v1.Group("/users")
	{
		groupTxn.POST("/post-wallet", h.CreateWallet)
	}

	return *r
}
