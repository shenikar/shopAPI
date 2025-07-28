package product

import (
	service "shopApi/internal/service"

	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	Service   *service.ProductService
	Validator *validator.Validate
}

func NewProductHandler(service *service.ProductService, validator *validator.Validate) *ProductHandler {
	return &ProductHandler{
		Service:   service,
		Validator: validator,
	}
}
