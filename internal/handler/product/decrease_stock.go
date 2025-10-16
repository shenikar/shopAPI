package product

import (
	"errors"
	"net/http"

	"github.com/shenikar/shopAPI/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DecreaseStockRequest struct {
	Quantity int `json:"quantity" validate:"required,min=1"`
}

// DecreaseStock
// @Summary      Decrease product stock
// @Description  Decreases the available stock of a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Param        quantity body      DecreaseStockRequest  true  "Quantity to decrease"
// @Success      200  {object}  dto.ProductResponseDTO "Successfully updated product"
// @Failure      400  {object}  map[string]string "Invalid request or validation failed"
// @Failure      404  {object}  map[string]string "Product not found"
// @Failure      500  {object}  map[string]string "Failed to decrease stock or fetch updated product"
// @Router       /api/v1/products/{id}/decrease [patch]
func (h *Handler) DecreaseStock(c *gin.Context) {
	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	var req DecreaseStockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}

	err = h.Service.DecreaseStock(c.Request.Context(), productID, req.Quantity)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not enough stock or product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decrease stock"})
		return
	}

	// Получаем обновлённый товар
	product, err := h.Service.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch updated product"})
		return
	}

	c.JSON(http.StatusOK, product)
}
