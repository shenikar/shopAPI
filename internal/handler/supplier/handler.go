package supplier

import (
	"github.com/shenikar/shopAPI/internal/service"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Service   *service.SupplierService
	Validator *validator.Validate
}

func NewSupplierHandler(service *service.SupplierService, validator *validator.Validate) *Handler {
	return &Handler{
		Service:   service,
		Validator: validator,
	}
}
