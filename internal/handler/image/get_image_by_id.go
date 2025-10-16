package image

import (
	"errors"
	"net/http"

	"github.com/shenikar/shopAPI/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetImageByID
// @Summary      Get an image by its ID
// @Description  Retrieves an image from the store by its ID
// @Tags         images
// @Accept       json
// @Produce      octet-stream
// @Param        id   path      string  true  "Image ID"
// @Success      200  {string}  byte[] "Successfully retrieved image"
// @Failure      400  {object}  map[string]string "Invalid image ID"
// @Failure      404  {object}  map[string]string "Image not found"
// @Failure      500  {object}  map[string]string "Failed to get image"
// @Router       /api/v1/images/{id} [get]
func (h *Handler) GetImageByID(c *gin.Context) {
	idStr := c.Param("id")
	imageID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid image ID"})
		return
	}

	image, err := h.Service.GetImageByID(c.Request.Context(), imageID)
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
