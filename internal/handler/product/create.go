package product

import (
	"net/http"

	"github.com/shenikar/shopAPI/internal/dto"

	"github.com/gin-gonic/gin"
)

// CreateProduct
// @Summary      Create a new product
// @Description  Adds a new product to the store
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      dto.CreateProductDTO  true  "Product info"
// @Success      201      {object}  dto.ProductResponseDTO "Successfully created"
// @Failure      400      {object}  map[string]string "Invalid request or validation failed"
// @Failure      500      {object}  map[string]string "Failed to create product"
// @Router       /api/v1/products [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	product, err := h.Service.CreateProduct(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}
