package client

import (
	"errors"
	"net/http"
	"shopApi/internal/domain"
	"shopApi/internal/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UpdateAddress
// @Summary      Update a client's address
// @Description  Updates the address of a client by their ID
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Client ID"
// @Param        address  body      dto.CreateAddressDTO  true  "New address info"
// @Success      200  {object}  dto.ClientResponseDTO "Successfully updated client address"
// @Failure      400  {object}  map[string]string "Invalid client ID or request body"
// @Failure      404  {object}  map[string]string "Client not found"
// @Failure      500  {object}  map[string]string "Failed to update client address"
// @Router       /api/v1/clients/{id}/address [patch]
func (h *ClientHandler) UpdateAddress(c *gin.Context) {
	idStr := c.Param("id")
	clientID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
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
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update client address"})
		}
		return
	}

	c.JSON(http.StatusOK, client)
}
