package client

import (
	"net/http"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"

	"github.com/gin-gonic/gin"
)

func (h *ClientHandler) GetClientByNameSurname(c *gin.Context) {
	name := c.Query("name")
	surname := c.Query("surname")

	if name == "" || surname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and surname are required"})
		return
	}

	clientWithAddresses, err := h.Repo.FindByNameSurname(c.Request.Context(), name, surname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}
	if len(clientWithAddresses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	var response []dto.ClientResponseDTO
	for _, cwa := range clientWithAddresses {
		response = append(response, mapper.ToClientResponseDTO(cwa.Client, cwa.Address))
	}
	c.JSON(http.StatusOK, response)
}
