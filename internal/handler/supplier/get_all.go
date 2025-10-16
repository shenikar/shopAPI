package supplier

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllSuppliers
// @Summary      Get all suppliers
// @Description  Retrieves a list of all suppliers
// @Tags         suppliers
// @Accept       json
// @Produce      json
// @Success      200      {array}   dto.SupplierResponseDTO "Successfully retrieved suppliers"
// @Failure      500      {object}  map[string]string "Failed to fetch suppliers"
// @Router       /api/v1/suppliers [get]
func (h *Handler) GetAllSuppliers(c *gin.Context) {
	suppliers, err := h.Service.GetAllSuppliers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch suppliers"})
		return
	}
	c.JSON(http.StatusOK, suppliers)
}
