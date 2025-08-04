package image

import (
	"net/http"
	"shopApi/internal/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "image updated"})
}
