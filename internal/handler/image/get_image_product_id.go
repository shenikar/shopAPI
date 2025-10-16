package image

import (
	"errors"
	"net/http"

	"github.com/shenikar/shopAPI/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetImageByProductID
// @Summary      Get an image by its product ID
// @Description  Retrieves an image from the store by its product ID
// @Tags         images
// @Accept       json
// @Produce      octet-stream
// @Param        product_id   path      string  true  "Product ID"
// @Success      200  {string}  byte[] "Successfully retrieved image"
// @Failure      400  {object}  map[string]string "Invalid product ID"
// @Failure      404  {object}  map[string]string "Image not found"
// @Failure      500  {object}  map[string]string "Failed to get image"
// @Router       /api/v1/images/product/{product_id} [get]
func (h *Handler) GetImageByProductID(c *gin.Context) {
	idStr := c.Param("product_id")
	productID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	image, err := h.Service.GetImageByProductID(c.Request.Context(), productID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get image"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+image.ID.String()+".bin")
	c.Data(http.StatusOK, "application/octet-stream", image.Image)
}
