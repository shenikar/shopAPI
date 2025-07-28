package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) GetAllAvailableProducts(c *gin.Context) {
	products, err := h.Service.GetAllAvailableProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)

}
