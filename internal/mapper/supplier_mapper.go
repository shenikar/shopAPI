package mapper

import (
	"github.com/shenikar/shopAPI/internal/domain/models"
	"github.com/shenikar/shopAPI/internal/dto"
)

func ToSupplierEntity(dto dto.CreateSupplierDTO) models.Supplier {
	return models.Supplier{
		Name:        dto.Name,
		PhoneNumber: dto.PhoneNumber,
	}
}

func ToSupplierResponseDTO(supplier models.Supplier, address models.Address) dto.SupplierResponseDTO {
	return dto.SupplierResponseDTO{
		ID:          supplier.ID.String(),
		Name:        supplier.Name,
		Address:     ToAddressResponseDTO(address),
		PhoneNumber: supplier.PhoneNumber,
	}
}
