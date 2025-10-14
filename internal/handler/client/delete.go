package client

import (
	"errors"
	"net/http"
	"shopApi/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeleteClient
// @Summary      Delete a client by its ID
// @Description  Deletes a client from the store by its ID
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Client ID"
// @Success      200  {object}  map[string]string "Successfully deleted"
// @Failure      400  {object}  map[string]string "Invalid client ID"
// @Failure      404  {object}  map[string]string "Client not found"
// @Failure      500  {object}  map[string]string "Failed to delete client"
// @Router       /api/v1/clients/{id} [delete]
func (h *ClientHandler) DeleteClient(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	err = h.Service.DeleteClient(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}
