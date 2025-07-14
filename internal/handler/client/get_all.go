package client

import (
	"net/http"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	clientsWithAddress, err := h.Repo.GetAllClient(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch clients"})
		return
	}

	var response []dto.ClientResponseDTO
	for _, item := range clientsWithAddress {
		response = append(response, mapper.ToClientResponseDTO(item.Client, item.Address))
	}
	c.JSON(http.StatusOK, response)

}
