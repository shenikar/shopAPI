package client

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllClients
// @Summary      Get all clients
// @Description  Retrieves a list of all clients with optional pagination
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        limit   query      int  false  "Limit"
// @Param        offset  query      int  false  "Offset"
// @Success      200  {array}   dto.ClientResponseDTO "Successfully retrieved clients"
// @Failure      400  {object}  map[string]string "Invalid limit or offset value"
// @Failure      500  {object}  map[string]string "Failed to fetch clients"
// @Router       /api/v1/clients [get]
func (h *ClientHandler) GetAllClients(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	var limit, offset *int

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil || l < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
			return
		}
		limit = &l
	}
	if offsetStr != "" {
		o, err := strconv.Atoi(offsetStr)
		if err != nil || o < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset value"})
			return
		}
		offset = &o
	}

	clients, err := h.Service.GetAllClient(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}

	c.JSON(http.StatusOK, clients)
}
