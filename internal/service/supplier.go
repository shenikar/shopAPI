package service

import (
	"context"
	"database/sql"
	"errors"
	"shopApi/internal/domain"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"
	"shopApi/internal/repository"

	"github.com/google/uuid"
)

type SupplierService struct {
	repo           *repository.SupplierRepository
	addressService *AddressService
}

func NewSupplierService(repo *repository.SupplierRepository, addressService *AddressService) *SupplierService {
	return &SupplierService{
		repo:           repo,
		addressService: addressService,
	}
}

func (s *SupplierService) CreateSupplier(ctx context.Context, req dto.CreateSupplierDTO) (*dto.SupplierResponseDTO, error) {
	supplier := mapper.ToSupplierEntity(req)
	address := mapper.ToAddressEntity(req.Address)

	addressID, err := s.addressService.Create(ctx, &address)
	if err != nil {
		return nil, err
	}

	supplier.ID = uuid.New()
	supplier.AddressID = addressID

	err = s.repo.CreateSupplier(ctx, supplier)
	if err != nil {
		return nil, err
	}

	response := mapper.ToSupplierResponseDTO(supplier, address)
	return &response, nil
}

func (s *SupplierService) GetSupplierByID(ctx context.Context, id uuid.UUID) (*dto.SupplierResponseDTO, error) {
	supplier, err := s.repo.GetSupplierByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	address, err := s.addressService.GetByID(ctx, supplier.AddressID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	response := mapper.ToSupplierResponseDTO(supplier, *address)
	return &response, nil
}

func (s *SupplierService) GetAllSuppliers(ctx context.Context) ([]dto.SupplierResponseDTO, error) {
	suppliers, err := s.repo.GetAllSupplier(ctx)
	if err != nil {
		return nil, err
	}
	var result []dto.SupplierResponseDTO
	for _, supplier := range suppliers {
		address, err := s.addressService.GetByID(ctx, supplier.AddressID)
		if err != nil {
			// In a real application, you might want to handle this more gracefully
			// For now, we'll just skip this supplier
			continue
		}
		result = append(result, mapper.ToSupplierResponseDTO(supplier, *address))
	}
	return result, nil
}

func (s *SupplierService) DeleteSupplier(ctx context.Context, id uuid.UUID) error {
	err := s.repo.DeleteSupplier(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrNotFound
		}
		return err
	}
	return nil
}

func (s *SupplierService) UpdateSupplierAddress(ctx context.Context, id uuid.UUID, req dto.CreateAddressDTO) (*dto.SupplierResponseDTO, error) {
	supplier, err := s.repo.GetSupplierByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	address := mapper.ToAddressEntity(req)
	address.ID = supplier.AddressID

	err = s.addressService.Update(ctx, &address)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	response := mapper.ToSupplierResponseDTO(supplier, address)
	return &response, nil
}
