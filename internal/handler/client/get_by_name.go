package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
