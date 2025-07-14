package dto

type CreateSupplierDTO struct {
	Name        string `json:"name" validate:"required"`
	AddressID   int    `json:"address_id" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type SupplierResponseDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AddressID   int    `json:"address_id"`
	PhoneNumber string `json:"phone_number"`
}
