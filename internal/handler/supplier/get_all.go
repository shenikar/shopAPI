package supplier

import (
	"net/http"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"

	"github.com/gin-gonic/gin"
)

func (h *SupplierHandler) GetAllSuppliers(c *gin.Context) {
	suppliers, err := h.Repo.GetAllSupplier(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch suppliers"})
		return
	}
	var response []dto.SupplierResponseDTO
	for _, supplier := range suppliers {
		response = append(response, mapper.ToSupplierResponseDTO(supplier))
	}
	c.JSON(http.StatusOK, response)
}
