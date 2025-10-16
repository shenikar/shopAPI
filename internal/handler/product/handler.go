package product

import (
	service "github.com/shenikar/shopAPI/internal/service"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Service   *service.ProductService
	Validator *validator.Validate
}

func NewProductHandler(service *service.ProductService, validator *validator.Validate) *Handler {
	return &Handler{
		Service:   service,
		Validator: validator,
	}
}
