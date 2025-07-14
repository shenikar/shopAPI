package mapper

import (
	"shopApi/internal/domain/models"
	"shopApi/internal/dto"
	"time"
)

func ToProductEntity(dto dto.CreateProductDTO) (models.Product, error) {
	return models.Product{
		Name:           dto.Name,
		Category:       dto.Category,
		Price:          dto.Price,
		AvailableStock: dto.AvailableStock,
		SupplierID:     dto.SupplierID,
		ImageID:        dto.ImageID,
		LastUpdateDate: time.Now(),
	}, nil
}

func ToProductResponseDTO(product models.Product) dto.ProductResponseDTO {
	return dto.ProductResponseDTO{
		ID:             product.ID.String(),
		Name:           product.Name,
		Category:       product.Category,
		Price:          product.Price,
		AvailableStock: product.AvailableStock,
		LastUpdateDate: product.LastUpdateDate.Format("2006-01-02"),
		SupplierID:     product.SupplierID,
		ImageID:        product.ImageID.String(),
	}
}
