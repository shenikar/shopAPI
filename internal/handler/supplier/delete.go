package supplier

import (
	"errors"
	"net/http"
	"shopApi/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeleteSupplier
// @Summary      Delete a supplier by its ID
// @Description  Deletes a supplier from the store by its ID
// @Tags         suppliers
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Supplier ID"
// @Success      200  {object}  map[string]string "Successfully deleted"
// @Failure      400  {object}  map[string]string "Invalid supplier ID"
// @Failure      404  {object}  map[string]string "Supplier not found"
// @Failure      500  {object}  map[string]string "Failed to delete supplier"
// @Router       /api/v1/suppliers/{id} [delete]
func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid supplier id"})
		return
	}

	err = h.Service.DeleteSupplier(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "supplier not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete supplier"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "supplier deleted"})
}
