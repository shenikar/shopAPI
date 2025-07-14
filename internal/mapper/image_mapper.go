package mapper

import (
	"shopApi/internal/domain/models"
	"shopApi/internal/dto"

	"github.com/google/uuid"
)

func ToImageEntity(dto dto.ImageUploadDTO, productID uuid.UUID) (models.Image, error) {
	return models.Image{
		ID:    uuid.New(),
		Image: dto.ImageDate,
	}, nil
}

func ToImageResponseDTO(image models.Image) dto.ImageResponseDTO {
	return dto.ImageResponseDTO{
		ID:       image.ID.String(),
		ImageRaw: image.Image,
	}
}
