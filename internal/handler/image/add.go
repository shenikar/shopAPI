package image

import (
	"net/http"

	"github.com/shenikar/shopAPI/internal/dto"

	"github.com/gin-gonic/gin"
)

// AddImage
// @Summary      Add a new image
// @Description  Adds a new image and associates it with a product
// @Tags         images
// @Accept       json
// @Produce      json
// @Param        image  body      dto.ImageUploadDTO  true  "Image data and product ID"
// @Success      201    {object}  map[string]string "Successfully created"
// @Failure      400    {object}  map[string]string "Invalid request or validation failed"
// @Failure      500    {object}  map[string]string "Failed to create image"
// @Router       /api/v1/images [post]
func (h *Handler) AddImage(c *gin.Context) {
	var req dto.ImageUploadDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageID, err := h.Service.CreateImage(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": imageID.String()})
}
