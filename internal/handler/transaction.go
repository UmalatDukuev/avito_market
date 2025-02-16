package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) sendCoin(c *gin.Context) {
	var req SendCoinRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "auth error")
		return
	}

	err = h.services.Transaction.SendCoins(userID, req.ToUser, req.Amount)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"from_user": userID,
		"to_user":   req.ToUser,
		"amount":    req.Amount,
	})
}
