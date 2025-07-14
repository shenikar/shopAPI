package client

import (
	"shopApi/internal/repository"

	"github.com/go-playground/validator/v10"
)

type ClientHandler struct {
	Repo      *repository.ClientRepository
	Validator *validator.Validate
}

func NewClientHandler(repo *repository.ClientRepository, validator *validator.Validate) *ClientHandler {
	return &ClientHandler{
		Repo:      repo,
		Validator: validator,
	}
}
