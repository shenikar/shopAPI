package mapper

import (
	"shopApi/internal/domain/models"
	"shopApi/internal/dto"
	"time"

	"github.com/google/uuid"
)

func ToClientEntity(dto dto.CreateClientDTO, addressID uuid.UUID) models.Client {
	return models.Client{
		ID:               uuid.New(),
		ClientName:       dto.ClientName,
		ClientSurname:    dto.ClientSurname,
		Birthday:         dto.Birthday,
		Gender:           dto.Gender,
		RegistrationDate: time.Now(),
		AddressID:        addressID,
	}
}

func ToClientResponseDTO(client models.Client, address models.Address) dto.ClientResponseDTO {
	return dto.ClientResponseDTO{
		ID:               client.ID.String(),
		ClientName:       client.ClientName,
		ClientSurname:    client.ClientSurname,
		Birthday:         client.Birthday,
		Gender:           client.Gender,
		RegistrationDate: client.RegistrationDate,
		Address:          ToAddressResponseDTO(address),
	}
}
