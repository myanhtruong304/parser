package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myanhtruong304/parser/package/model"
)

func (h *Handler) CreateWallet(c *gin.Context) {
	var req model.CreateWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	wallet, err := h.entity.CreateWallet(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"wallet": wallet})
}

func (h *Handler) GetWalletTransaction(c *gin.Context) {
	var req model.GetWalletTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	wallet, err := h.entity.GetWalletTransaction(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}
