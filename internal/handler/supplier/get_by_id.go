package supplier

import (
	"errors"
	"net/http"
	"shopApi/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetSupplierByID
// @Summary      Get a supplier by its ID
// @Description  Retrieves a supplier from the store by its ID
// @Tags         suppliers
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Supplier ID"
// @Success      200  {object}  dto.SupplierResponseDTO "Successfully retrieved supplier"
// @Failure      400  {object}  map[string]string "Invalid supplier ID"
// @Failure      404  {object}  map[string]string "Supplier not found"
// @Failure      500  {object}  map[string]string "Failed to fetch supplier"
// @Router       /api/v1/suppliers/{id} [get]
func (h *SupplierHandler) GetSupplierByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid supplier ID"})
		return
	}

	supplier, err := h.Service.GetSupplierByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "supplier not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch supplier"})
		return
	}
	c.JSON(http.StatusOK, supplier)
}
