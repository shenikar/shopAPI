package image

import (
	"github.com/shenikar/shopAPI/internal/service"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Service   *service.ImageService
	Validator *validator.Validate
}

func NewImageHandler(service *service.ImageService, validator *validator.Validate) *Handler {
	return &Handler{
		Service:   service,
		Validator: validator,
	}
}
