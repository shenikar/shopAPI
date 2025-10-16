package mapper

import (
	"github.com/shenikar/shopAPI/internal/domain/models"
	"github.com/shenikar/shopAPI/internal/dto"

	"github.com/google/uuid"
)

func ToImageEntity(dto dto.ImageUploadDTO) (models.Image, error) {
	return models.Image{
		ID:    uuid.New(),
		Image: dto.ImageData,
	}, nil
}

func ToImageResponseDTO(image models.Image) dto.ImageResponseDTO {
	return dto.ImageResponseDTO{
		ID:       image.ID.String(),
		ImageRaw: image.Image,
	}
}
