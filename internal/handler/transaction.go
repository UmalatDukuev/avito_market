package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) sendCoin(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// info, err := h.services.Information.GetInfo(userID)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	c.JSON(http.StatusOK, map[string]interface{}{
		"id_from": userID,
		"coins":   10,
	})
}
