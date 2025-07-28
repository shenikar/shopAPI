package client

import (
	"database/sql"
	"errors"
	"net/http"
	"shopApi/internal/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateAddressRequest struct {
	AddressID string `json:"address_id" validate:"required"`
}

func (h *ClientHandler) UpdateAddress(c *gin.Context) {
	idStr := c.Param("id")
	clientID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
	}

	var req dto.CreateAddressDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	client, err := h.Service.UpdateAddress(c.Request.Context(), clientID, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update client address"})
		}
		return
	}

	c.JSON(http.StatusOK, client)

}
