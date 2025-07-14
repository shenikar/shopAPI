package product

import (
	"database/sql"
	"net/http"
	"shopApi/internal/mapper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DecreaseStockRequest struct {
	Quantity int `json:"quantity" validate:"required,min=1"`
}

func (h *ProductHandler) DecreaseStock(c *gin.Context) {
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

	err = h.Repo.DecreaseStock(c.Request.Context(), productID, req.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "not enough stock or product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decrease stock"})
		return
	}

	// Получаем обновлённый товар
	product, err := h.Repo.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch updated product"})
		return
	}

	c.JSON(http.StatusOK, mapper.ToProductResponseDTO(product))
}
