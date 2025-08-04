package image

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *ImageHandler) GetImageByID(c *gin.Context) {
	idStr := c.Param("id")
	imageID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid image ID"})
		return
	}

	image, err := h.Service.GetImageByID(c.Request.Context(), imageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+image.ID.String()+".bin")
	c.Data(http.StatusOK, "application/octet-stream", image.Image)
}
