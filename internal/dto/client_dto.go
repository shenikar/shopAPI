package dto

import (
	"time"
)

type CreateClientDTO struct {
	ClientName    string           `json:"client_name" validate:"required"`
	ClientSurname string           `json:"client_surname" validate:"required"`
	Birthday      time.Time        `json:"birthday" validate:"required"`
	Gender        string           `json:"gender" validate:"oneof=male female other"`
	Address       CreateAddressDTO `json:"address" validate:"required"`
}

type ClientResponseDTO struct {
	ID               string             `json:"id"`
	ClientName       string             `json:"client_name"`
	ClientSurname    string             `json:"client_surname"`
	Birthday         time.Time          `json:"birthday"`
	Gender           string             `json:"gender"`
	RegistrationDate time.Time          `json:"registration_date"`
	Address          AddressResponseDTO `json:"address"`
}
