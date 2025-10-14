package client

import (
	"net/http"
	"shopApi/internal/dto"

	"github.com/gin-gonic/gin"
)

// CreateClient
// @Summary      Create a new client
// @Description  Adds a new client to the store
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        client  body      dto.CreateClientDTO  true  "Client info"
// @Success      201      {object}  dto.ClientResponseDTO "Successfully created"
// @Failure      400      {object}  map[string]string "Invalid request or validation failed"
// @Failure      500      {object}  map[string]string "Failed to create client"
// @Router       /api/v1/clients [post]
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
