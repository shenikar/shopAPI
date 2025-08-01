package dto

import "github.com/google/uuid"

type CreateProductDTO struct {
	Name           string    `json:"name" validate:"required"`
	Category       string    `json:"category" validate:"required"`
	Price          float64   `json:"price" validate:"required,gt=0"`
	AvailableStock int       `json:"available_stock" validate:"required,gte=0"`
	SupplierID     int       `json:"supplier_id" validate:"required"`
	ImageID        uuid.UUID `json:"image_id" validate:"required"`
}

type ProductResponseDTO struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Price          float64 `json:"price"`
	AvailableStock int     `json:"available_stock"`
	LastUpdateDate string  `json:"last_update_date"`
	SupplierID     int     `json:"supplier_id"`
	ImageID        string  `json:"image_id"`
}
