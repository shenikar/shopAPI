package product

import (
	"net/http"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) GetAllAvailableProducts(c *gin.Context) {
	products, err := h.Repo.GetAllAvailableProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	var response []dto.ProductResponseDTO
	for _, product := range products {
		response = append(response, mapper.ToProductResponseDTO(product))
	}
	c.JSON(http.StatusOK, response)

}
