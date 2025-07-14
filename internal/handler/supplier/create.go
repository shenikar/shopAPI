package supplier

import (
	"net/http"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	supplier, _ := mapper.ToSupplierEntity(req)
	supplier.ID = uuid.New()

	err := h.Repo.CreateSupplier(c.Request.Context(), supplier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create supplier"})
		return
	}

	c.JSON(http.StatusCreated, mapper.ToSupplierResponseDTO(supplier))
}
