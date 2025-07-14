package supplier

import (
	"shopApi/internal/repository"

	"github.com/go-playground/validator/v10"
)

type SupplierHandler struct {
	Repo      *repository.SupplierRepository
	Validator *validator.Validate
}

func NewSupplierHandler(repo *repository.SupplierRepository, validator *validator.Validate) *SupplierHandler {
	return &SupplierHandler{
		Repo:      repo,
		Validator: validator,
	}
}
