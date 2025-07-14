package image

import (
	"shopApi/internal/repository"

	"github.com/go-playground/validator/v10"
)

type ImageHandler struct {
	Repo      *repository.ImageRepository
	Validator *validator.Validate
}

func NewImageHandler(repo *repository.ImageRepository, validator *validator.Validate) *ImageHandler {
	return &ImageHandler{
		Repo:      repo,
		Validator: validator,
	}
}
