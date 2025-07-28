package supplier

import (
	"net/http"
	"shopApi/internal/dto"

	"github.com/gin-gonic/gin"
)

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
