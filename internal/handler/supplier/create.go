package supplier

import (
	"net/http"
	"shopApi/internal/dto"

	"github.com/gin-gonic/gin"
)

// CreateSupplier
// @Summary      Create a new supplier
// @Description  Adds a new supplier to the store
// @Tags         suppliers
// @Accept       json
// @Produce      json
// @Param        supplier  body      dto.CreateSupplierDTO  true  "Supplier info"
// @Success      201      {object}  dto.SupplierResponseDTO "Successfully created"
// @Failure      400      {object}  map[string]string "Invalid request or validation failed"
// @Failure      500      {object}  map[string]string "Failed to create supplier"
// @Router       /api/v1/suppliers [post]
func (h *SupplierHandler) CreateSupplier(c *gin.Context) {
	var req dto.CreateSupplierDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}

	supplier, err := h.Service.CreateSupplier(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create supplier"})
		return
	}

	c.JSON(http.StatusCreated, supplier)
}
