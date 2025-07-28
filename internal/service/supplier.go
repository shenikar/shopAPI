package service

import (
	"context"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"
	"shopApi/internal/repository"

	"github.com/google/uuid"
)

type SupplierService struct {
	repo *repository.SupplierRepository
}

func NewSupplierService(repo *repository.SupplierRepository) *SupplierService {
	return &SupplierService{
		repo: repo,
	}
}

func (s *SupplierService) CreateSupplier(ctx context.Context, req dto.CreateSupplierDTO) (dto.SupplierResponseDTO, error) {
	supplier, _ := mapper.ToSupplierEntity(req)
	supplier.ID = uuid.New()

	err := s.repo.CreateSupplier(ctx, supplier)
	if err != nil {
		return dto.SupplierResponseDTO{}, err
	}
	return mapper.ToSupplierResponseDTO(supplier), nil
}

func (s *SupplierService) GetSupplierByID(ctx context.Context, id uuid.UUID) (dto.SupplierResponseDTO, error) {
	supplier, err := s.repo.GetSupplierByID(ctx, id)
	if err != nil {
		return dto.SupplierResponseDTO{}, err
	}
	return mapper.ToSupplierResponseDTO(supplier), nil
}

func (s *SupplierService) GetAllSuppliers(ctx context.Context) ([]dto.SupplierResponseDTO, error) {
	suppliers, err := s.repo.GetAllSupplier(ctx)
	if err != nil {
		return nil, err
	}
	var result []dto.SupplierResponseDTO
	for _, supplier := range suppliers {
		result = append(result, mapper.ToSupplierResponseDTO(supplier))
	}
	return result, nil
}

func (s *SupplierService) DeleteSupplier(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteSupplier(ctx, id)
}

func (s *SupplierService) UpdateSupplierADdress(ctx context.Context, id uuid.UUID, addressID int) error {
	return s.repo.UpdateAddress(ctx, id, addressID)
}
