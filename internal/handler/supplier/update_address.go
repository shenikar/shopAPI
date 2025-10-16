package supplier

import (
	"errors"
	"net/http"

	"github.com/shenikar/shopAPI/internal/domain"
	"github.com/shenikar/shopAPI/internal/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UpdateAddressSupplier
// @Summary      Update a supplier's address
// @Description  Updates the address of a supplier by their ID
// @Tags         suppliers
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Supplier ID"
// @Param        address  body      dto.CreateAddressDTO  true  "New address info"
// @Success      200  {object}  dto.SupplierResponseDTO "Successfully updated supplier address"
// @Failure      400  {object}  map[string]string "Invalid supplier ID or request body"
// @Failure      404  {object}  map[string]string "Supplier not found"
// @Failure      500  {object}  map[string]string "Failed to update supplier address"
// @Router       /api/v1/suppliers/{id}/address [patch]
func (h *Handler) UpdateAddressSupplier(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid supplier ID"})
		return
	}
	var req dto.CreateAddressDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}

	supplier, err := h.Service.UpdateSupplierAddress(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "supplier not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update address"})
		return
	}

	c.JSON(http.StatusOK, supplier)
}
