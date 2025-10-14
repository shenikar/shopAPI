package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllAvailableProducts
// @Summary      Get all available products
// @Description  Retrieves a list of all products with available stock
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {array}  dto.ProductResponseDTO "Successfully retrieved products"
// @Failure      500  {object}  map[string]string "Failed to get products"
// @Router       /api/v1/products/available [get]
func (h *ProductHandler) GetAllAvailableProducts(c *gin.Context) {
	products, err := h.Service.GetAllAvailableProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
