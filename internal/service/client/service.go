package service

import (
	"context"
	"shopApi/internal/dto"
	"shopApi/internal/mapper"
	"shopApi/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ClientService struct {
	repo      repository.ClientRepository
	validator *validator.Validate
}

func NewClientService(repo repository.ClientRepository, validator *validator.Validate) *ClientService {
	return &ClientService{
		repo:      repo,
		validator: validator,
	}
}

func (s *ClientService) CreateClient(ctx context.Context, req dto.CreateClientDTO) (dto.ClientResponseDTO, error) {
	if err := s.validator.Struct(req); err != nil {
		return dto.ClientResponseDTO{}, err
	}

	client, address, err := s.repo.CreateClient(ctx, req)
	if err != nil {
		return dto.ClientResponseDTO{}, err
	}

	return mapper.ToClientResponseDTO(client, address), nil
}

func (s *ClientService) DeleteClient(ctx context.Context, clientID uuid.UUID) error {
	return s.repo.DeleteClient(ctx, clientID)
}

func (s *ClientService) GetClientByNameSurname(ctx context.Context, name, surname string) ([]dto.ClientResponseDTO, error) {
	clients, err := s.repo.FindByNameSurname(ctx, name, surname)
	if err != nil {
		return nil, err
	}

	var result []dto.ClientResponseDTO
	for _, item := range clients {
		result = append(result, mapper.ToClientResponseDTO(item.Client, item.Address))
	}

	return result, nil
}

func (s *ClientService) GetAllClient(ctx context.Context, limit, offset *int) ([]dto.ClientResponseDTO, error) {
	clients, err := s.repo.GetAllClient(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var result []dto.ClientResponseDTO
	for _, item := range clients {
		result = append(result, mapper.ToClientResponseDTO(item.Client, item.Address))
	}
	return result, nil
}

func (s *ClientService) UpdateAddress(ctx context.Context, clientID uuid.UUID, req dto.CreateAddressDTO) (dto.ClientResponseDTO, error) {
	clent, address, err := s.repo.UpdateAddress(ctx, clientID, req)
	if err != nil {
		return dto.ClientResponseDTO{}, err
	}

	return mapper.ToClientResponseDTO(clent, address), nil
}
