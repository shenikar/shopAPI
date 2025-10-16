package client

import (
	service "github.com/shenikar/shopAPI/internal/service"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Service   *service.ClientService
	Validator *validator.Validate
}

func NewClientHandler(service *service.ClientService, validator *validator.Validate) *Handler {
	return &Handler{
		Service:   service,
		Validator: validator,
	}
}
