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

type ClientService struct {
	repo           *repository.ClientRepository
	addressService *AddressService
}

func NewClientService(repo *repository.ClientRepository, addressService *AddressService) *ClientService {
	return &ClientService{
		repo:           repo,
		addressService: addressService,
	}
}

func (s *ClientService) CreateClient(ctx context.Context, req dto.CreateClientDTO) (*dto.ClientResponseDTO, error) {
	client := mapper.ToClientEntity(req)
	address := mapper.ToAddressEntity(req.Address)

	addressID, err := s.addressService.Create(ctx, &address)
	if err != nil {
		return nil, err
	}
	client.AddressID = addressID

	createdClient, err := s.repo.CreateClient(ctx, client)
	if err != nil {
		return nil, err
	}

	response := mapper.ToClientResponseDTO(*createdClient, address)
	return &response, nil
}

func (s *ClientService) DeleteClient(ctx context.Context, clientID uuid.UUID) error {
	err := s.repo.DeleteClient(ctx, clientID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrNotFound
		}
		return err
	}
	return nil
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

func (s *ClientService) UpdateAddress(ctx context.Context, clientID uuid.UUID, req dto.CreateAddressDTO) (*dto.ClientResponseDTO, error) {
	clientWithAddress, err := s.repo.GetClientByID(ctx, clientID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	addressModel := mapper.ToAddressEntity(req)
	addressModel.ID = clientWithAddress.Address.ID

	err = s.addressService.Update(ctx, &addressModel)
	if err != nil {
		return nil, err
	}

	response := mapper.ToClientResponseDTO(clientWithAddress.Client, addressModel)
	return &response, nil
}
