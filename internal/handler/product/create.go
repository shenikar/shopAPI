package product

import (
	"net/http"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input dto.CreateProductDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data: " + err.Error()})
		return
	}

	if err := h.Validator.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_validation": err.Error()})
		return

	}

	product, _ := mapper.ToProductEntity(input)
	product.ID = uuid.New()

	err := h.Repo.CreateProduct(c.Request.Context(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, mapper.ToProductResponseDTO(product))

}
