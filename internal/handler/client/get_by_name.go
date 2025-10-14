package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetClientByNameSurname
// @Summary      Get clients by name and surname
// @Description  Retrieves a list of clients matching the given name and surname
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        name     query      string  true  "Client name"
// @Param        surname  query      string  true  "Client surname"
// @Success      200      {array}   dto.ClientResponseDTO "Successfully retrieved clients"
// @Failure      400      {object}  map[string]string "Missing name or surname parameters"
// @Failure      500      {object}  map[string]string "Failed to fetch clients"
// @Router       /api/v1/clients/search [get]
func (h *ClientHandler) GetClientByNameSurname(c *gin.Context) {
	name := c.Query("name")
	surname := c.Query("surname")

	if name == "" || surname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing name or surname parameters"})
		return
	}

	clients, err := h.Service.GetClientByNameSurname(c.Request.Context(), name, surname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}

	c.JSON(http.StatusOK, clients)
}
