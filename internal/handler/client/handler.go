package client

import (
	service "shopApi/internal/service"

	"github.com/go-playground/validator/v10"
)

type ClientHandler struct {
	Service   *service.ClientService
	Validator *validator.Validate
}

func NewClientHandler(service *service.ClientService, validator *validator.Validate) *ClientHandler {
	return &ClientHandler{
		Service:   service,
		Validator: validator,
	}
}
