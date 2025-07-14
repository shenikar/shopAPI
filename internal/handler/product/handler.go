package product

import (
	"shopApi/internal/repository"

	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	Repo      *repository.ProductRepository
	Validator *validator.Validate
}

func NewProductHandler(repo *repository.ProductRepository, validator *validator.Validate) *ProductHandler {
	return &ProductHandler{
		Repo:      repo,
		Validator: validator,
	}
}
