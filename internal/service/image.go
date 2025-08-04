package service

import (
	"context"
	"shopApi/internal/domain/models"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"
	"shopApi/internal/repository"

	"github.com/google/uuid"
)

type ImageService struct {
	repo *repository.ImageRepository
}

func NewImageService(repo *repository.ImageRepository) *ImageService {
	return &ImageService{
		repo: repo,
	}
}

func (s *ImageService) CreateImage(ctx context.Context, dto dto.ImageUploadDTO) (uuid.UUID, error) {
	productID, err := uuid.Parse(dto.ProductID)
	if err != nil {
		return uuid.Nil, err
	}

	imageEntity, err := mapper.ToImageEntity(dto)
	if err != nil {
		return uuid.Nil, err
	}

	// Сохраняем изображение в базе
	if err := s.repo.CreateImage(ctx, &imageEntity); err != nil {
		return uuid.Nil, err
	}

	// Обновляем связь product -> image (одиночное изображение на товар)
	if err := s.repo.UpdateProductImageID(ctx, productID, imageEntity.ID); err != nil {
		return uuid.Nil, err
	}

	return imageEntity.ID, nil
}

func (s *ImageService) UpdateImage(ctx context.Context, id uuid.UUID, newImage []byte) error {
	_, err := s.repo.GetImageByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.UpdateImage(ctx, id, newImage)
}

func (s *ImageService) GetImageByID(ctx context.Context, id uuid.UUID) (*models.Image, error) {
	return s.repo.GetImageByID(ctx, id)
}

func (s *ImageService) GetImageByProductID(ctx context.Context, productID uuid.UUID) (*models.Image, error) {
	return s.repo.GetByProductID(ctx, productID)
}

func (s *ImageService) DeleteImage(ctx context.Context, id uuid.UUID) error {
	err := s.repo.ClearProductImageIDByImageID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.DeleteImage(ctx, id)
}
