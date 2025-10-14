package image

import (
	"errors"
	"net/http"
	"shopApi/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeleteImage
// @Summary      Delete an image by its ID
// @Description  Deletes an image from the store by its ID
// @Tags         images
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Image ID"
// @Success      200    {object}  map[string]string "Successfully deleted"
// @Failure      400    {object}  map[string]string "Invalid image ID"
// @Failure      404    {object}  map[string]string "Image not found"
// @Failure      500    {object}  map[string]string "Failed to delete image"
// @Router       /api/v1/images/{id} [delete]
func (h *ImageHandler) DeleteImage(c *gin.Context) {
	imageID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid image id"})
		return
	}

	err = h.Service.DeleteImage(c.Request.Context(), imageID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "image deleted"})
}
