package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) buyItem(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized user")
		return
	}

	itemName := c.Param("item")
	if itemName == "" {
		newErrorResponse(c, http.StatusBadRequest, "item name is required")
		return
	}

	err = h.services.Item.BuyItem(userID, itemName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "purchase successful"})
}
