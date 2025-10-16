package product

import (
	"errors"
	"net/http"

	"github.com/shenikar/shopAPI/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetProductByID
// @Summary      Get a product by its ID
// @Description  Retrieves a product from the store by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  dto.ProductResponseDTO "Successfully retrieved product"
// @Failure      400  {object}  map[string]string "Invalid product ID"
// @Failure      404  {object}  map[string]string "Product not found"
// @Failure      500  {object}  map[string]string "Failed to get product"
// @Router       /api/v1/products/{id} [get]
func (h *Handler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product, err := h.Service.GetProductByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get product"})
		return
	}

	c.JSON(http.StatusOK, product)
}
