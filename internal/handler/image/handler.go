package image

import (
	"shopApi/internal/service"

	"github.com/go-playground/validator/v10"
)

type ImageHandler struct {
	Service   *service.ImageService
	Validator *validator.Validate
}

func NewImageHandler(service *service.ImageService, validator *validator.Validate) *ImageHandler {
	return &ImageHandler{
		Service:   service,
		Validator: validator,
	}
}
