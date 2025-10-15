package mapper

import (
	"shopApi/internal/domain/models"
	"shopApi/internal/dto"
	"time"

	"github.com/google/uuid"
)

func ToProductEntity(dto dto.CreateProductDTO) (models.Product, error) {
	supplierID, err := uuid.Parse(dto.SupplierID)
	if err != nil {
		return models.Product{}, err
	}

	var imageID *uuid.UUID
	if dto.ImageID != nil {
		parsedUUID, err := uuid.Parse(*dto.ImageID)
		if err != nil {
			return models.Product{}, err
		}
		imageID = &parsedUUID
	}

	return models.Product{
		Name:           dto.Name,
		Category:       dto.Category,
		Price:          dto.Price,
		AvailableStock: dto.AvailableStock,
		SupplierID:     supplierID,
		ImageID:        imageID,
		LastUpdateDate: time.Now(),
	}, nil
}

func ToProductResponseDTO(product models.Product) dto.ProductResponseDTO {
	var imageID string
	if product.ImageID != nil {
		imageID = product.ImageID.String()
	}

	return dto.ProductResponseDTO{
		ID:             product.ID.String(),
		Name:           product.Name,
		Category:       product.Category,
		Price:          product.Price,
		AvailableStock: product.AvailableStock,
		LastUpdateDate: product.LastUpdateDate.Format(time.RFC3339),
		SupplierID:     product.SupplierID.String(),
		ImageID:        imageID,
	}
}
