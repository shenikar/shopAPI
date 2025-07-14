package dto

type CreateAddressDTO struct {
	Country string `json:"country" validate:"required"`
	City    string `json:"city" validate:"required"`
	Street  string `json:"street" validate:"required"`
}

type AddressResponseDTO struct {
	ID      string `json:"id"`
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
}
