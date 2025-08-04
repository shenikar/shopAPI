package image

import (
	"net/http"
	"shopApi/internal/dto"

	"github.com/gin-gonic/gin"
)

func (h *ImageHandler) AddImage(c *gin.Context) {
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
