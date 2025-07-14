package supplier

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid supplier ID"})
		return
	}

	err = h.Repo.DeleteSupplier(c.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "supplier not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete supplier"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "supplier deleted successfully"})
}
