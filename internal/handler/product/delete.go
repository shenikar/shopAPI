package product

import (
	"errors"
	"net/http"

	"github.com/shenikar/shopAPI/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeleteProduct
// @Summary      Delete a product by its ID
// @Description  Deletes a product from the store by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  map[string]string "Successfully deleted"
// @Failure      400  {object}  map[string]string "Invalid product ID"
// @Failure      404  {object}  map[string]string "Product not found"
// @Failure      500  {object}  map[string]string "Failed to delete product"
// @Router       /api/v1/products/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	err = h.Service.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
