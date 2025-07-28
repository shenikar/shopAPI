package client

import (
	"net/http"
	"shopApi/internal/dto"

	"github.com/gin-gonic/gin"
)

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var req dto.CreateClientDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	resp, err := h.Service.CreateClient(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
