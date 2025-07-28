package supplier

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *SupplierHandler) GetAllSuppliers(c *gin.Context) {
	suppliers, err := h.Service.GetAllSuppliers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch suppliers"})
		return
	}
	c.JSON(http.StatusOK, suppliers)
}
