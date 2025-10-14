package service

import (
	"context"
	"shopApi/internal/domain/models"
	"shopApi/internal/repository"

	"github.com/google/uuid"
)

type AddressService struct {
	repo *repository.AddressRepository
}

func NewAddressService(repo *repository.AddressRepository) *AddressService {
	return &AddressService{
		repo: repo,
	}
}

func (s *AddressService) Create(ctx context.Context, address *models.Address) (uuid.UUID, error) {
	return s.repo.Create(ctx, address)
}

func (s *AddressService) GetByID(ctx context.Context, id uuid.UUID) (*models.Address, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *AddressService) Update(ctx context.Context, address *models.Address) error {
	return s.repo.Update(ctx, address)
}