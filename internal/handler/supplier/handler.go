package supplier

import (
	"shopApi/internal/service"

	"github.com/go-playground/validator/v10"
)

type SupplierHandler struct {
	Service   *service.SupplierService
	Validator *validator.Validate
}

func NewSupplierHandler(service *service.SupplierService, validator *validator.Validate) *SupplierHandler {
	return &SupplierHandler{
		Service:   service,
		Validator: validator,
	}
}
