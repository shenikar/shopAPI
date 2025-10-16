package dto

type CreateSupplierDTO struct {
	Name        string           `json:"name" validate:"required"`
	Address     CreateAddressDTO `json:"address" validate:"required"`
	PhoneNumber string           `json:"phone_number" validate:"required"`
}

type SupplierResponseDTO struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Address     AddressResponseDTO `json:"address"`
	PhoneNumber string             `json:"phone_number"`
}
