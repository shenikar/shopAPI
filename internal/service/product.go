package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/shenikar/shopAPI/internal/domain"
	"github.com/shenikar/shopAPI/internal/dto"
	"github.com/shenikar/shopAPI/internal/mapper"
	"github.com/shenikar/shopAPI/internal/repository"

	"github.com/google/uuid"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, dto dto.CreateProductDTO) (*dto.ProductResponseDTO, error) {
	product, err := mapper.ToProductEntity(dto)
	if err != nil {
		return nil, err
	}
	product.ID = uuid.New()
	err = s.repo.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	response := mapper.ToProductResponseDTO(product)
	return &response, nil
}

func (s *ProductService) GetProductByID(ctx context.Context, id uuid.UUID) (dto.ProductResponseDTO, error) {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.ProductResponseDTO{}, domain.ErrNotFound
		}
		return dto.ProductResponseDTO{}, err
	}
	return mapper.ToProductResponseDTO(product), nil
}

func (s *ProductService) GetAllAvailableProducts(ctx context.Context) ([]dto.ProductResponseDTO, error) {
	products, err := s.repo.GetAllAvailableProducts(ctx)
	if err != nil {
		return nil, err
	}
	var result []dto.ProductResponseDTO
	for _, product := range products {
		result = append(result, mapper.ToProductResponseDTO(product))
	}
	return result, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProductService) DecreaseStock(ctx context.Context, id uuid.UUID, quantity int) error {
	err := s.repo.DecreaseStock(ctx, id, quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrNotFound
		}
		return err
	}
	return nil
}
