package image

import (
	"net/http"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product_id"})
		return
	}

	image, _ := mapper.ToImageEntity(req, productID)

	err = h.Repo.CreateImage(c.Request.Context(), &image)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "image added"})
}
