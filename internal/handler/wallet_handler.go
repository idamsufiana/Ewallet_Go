package handler

import (
	"net/http"

	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WalletHandler struct {
	Service *service.WalletService
}

func NewWalletHandler(svc *service.WalletService) *WalletHandler {
	return &WalletHandler{
		Service: svc,
	}
}

type WithdrawRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

func (h *WalletHandler) Withdraw(c *gin.Context) {

	var req WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	balance, err := h.Service.Withdraw(c.Request.Context(), uid, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"remaining_balance": balance,
	})
}

func (h *WalletHandler) GetBalance(c *gin.Context) {

	id := c.Param("user_id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	balance, err := h.Service.GetBalance(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": uid,
		"balance": balance,
	})
}
