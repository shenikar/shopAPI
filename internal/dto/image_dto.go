package dto

type ImageUploadDTO struct {
	ImageDate []byte `json:"image" validate:"required"`
	ProductID string `json:"product_id" validate:"required,uuid"`
}

type ImageResponseDTO struct {
	ID       string `json:"id"`
	ImageRaw []byte `json:"image"`
}
