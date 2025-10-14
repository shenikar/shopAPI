package mapper

import (
	"shopApi/internal/domain/models"
	"shopApi/internal/dto"
)

func ToClientEntity(dto dto.CreateClientDTO) models.Client {
	return models.Client{
		ClientName:    dto.ClientName,
		ClientSurname: dto.ClientSurname,
		Birthday:      dto.Birthday,
		Gender:        dto.Gender,
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
