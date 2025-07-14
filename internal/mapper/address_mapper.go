package mapper

import (
	"shopApi/internal/domain/models"
	"shopApi/internal/dto"

	"github.com/google/uuid"
)

func ToAddressEntity(dto dto.CreateAddressDTO) models.Address {
	return models.Address{
		ID:      uuid.New(),
		Country: dto.Country,
		City:    dto.City,
		Street:  dto.Street,
	}
}

func ToAddressResponseDTO(address models.Address) dto.AddressResponseDTO {
	return dto.AddressResponseDTO{
		ID:      address.ID.String(),
		Country: address.Country,
		City:    address.City,
		Street:  address.Street,
	}
}
