package client

import (
	"net/http"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"

	"github.com/gin-gonic/gin"
)

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var req dto.CreateClientDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	client, address, err := h.Repo.CreateClient(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapper.ToClientResponseDTO(client, address))
}
