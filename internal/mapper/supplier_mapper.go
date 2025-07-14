package mapper

import (
	"shopApi/internal/domain/models"
	"shopApi/internal/dto"
)

func ToSupplierEntity(dto dto.CreateSupplierDTO) (models.Supplier, error) {
	return models.Supplier{
		Name:        dto.Name,
		AddressID:   dto.AddressID,
		PhoneNumber: dto.PhoneNumber,
	}, nil
}

func ToSupplierResponseDTO(supplier models.Supplier) dto.SupplierResponseDTO {
	return dto.SupplierResponseDTO{
		ID:          supplier.ID.String(),
		Name:        supplier.Name,
		AddressID:   supplier.AddressID,
		PhoneNumber: supplier.PhoneNumber,
	}
}
