package image

import (
	"errors"
	"net/http"
	"shopApi/internal/domain"
	"shopApi/internal/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UpdateImage
// @Summary      Update an image
// @Description  Updates an image by its ID
// @Tags         images
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Image ID"
// @Param        image  body      dto.ImageUploadDTO  true  "New image data"
// @Success      200    {object}  map[string]string "Successfully updated"
// @Failure      400    {object}  map[string]string "Invalid image ID or request body"
// @Failure      404    {object}  map[string]string "Image not found"
// @Failure      500    {object}  map[string]string "Failed to update image"
// @Router       /api/v1/images/{id} [patch]
func (h *ImageHandler) UpdateImage(c *gin.Context) {
	imageID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid image id"})
		return
	}

	var req dto.ImageUploadDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.Service.UpdateImage(c.Request.Context(), imageID, req.ImageData)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "image updated"})
}
