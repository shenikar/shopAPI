package mapper

import (
	"shopApi/internal/domain/models"
	"shopApi/internal/dto"

	"github.com/google/uuid"
)

func ToSupplierEntity(dto dto.CreateSupplierDTO) (models.Supplier, error) {
	addressUUID, err := uuid.Parse(dto.AddressID)
	if err != nil {
		return models.Supplier{}, err
	}
	return models.Supplier{
		Name:        dto.Name,
		AddressID:   addressUUID,
		PhoneNumber: dto.PhoneNumber,
	}, nil
}

func ToSupplierResponseDTO(supplier models.Supplier) dto.SupplierResponseDTO {
	return dto.SupplierResponseDTO{
		ID:          supplier.ID.String(),
		Name:        supplier.Name,
		AddressID:   supplier.AddressID.String(),
		PhoneNumber: supplier.PhoneNumber,
	}
}
