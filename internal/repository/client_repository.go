package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/shenikar/shopAPI/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ClientWithAddress struct {
	Client  models.Client
	Address models.Address
}

type ClientRepository struct {
	db *sqlx.DB
}

func NewClientRepository(db *sqlx.DB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

// CreateClient создает нового пользователя
func (r *ClientRepository) CreateClient(ctx context.Context, client models.Client) (*models.Client, error) {
	client.ID = uuid.New()
	client.RegistrationDate = time.Now()

	query := `
		INSERT INTO client (id, client_name, client_surname, birthday, gender, registration_date, address_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.ExecContext(ctx, query, client.ID, client.ClientName, client.ClientSurname, client.Birthday, client.Gender, client.RegistrationDate, client.AddressID)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

// DeleteClient удаление пользователя
func (r *ClientRepository) DeleteClient(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM client WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// GetAllClient получение всех пользователей с пагинацией
func (r *ClientRepository) GetAllClient(ctx context.Context, limit, offset *int) ([]ClientWithAddress, error) {
	query := `
	SELECT 
		c.id, c.client_name, c.client_surname, c.birthday, c.gender, c.registration_date, c.address_id,
		a.id, a.country, a.city, a.street
	FROM client c
	JOIN address a ON c.address_id = a.id
	`
	var args []interface{}
	argIdx := 1

	if limit != nil {
		query += fmt.Sprintf(" LIMIT $%d", argIdx)
		args = append(args, *limit)
		argIdx++
	}
	if offset != nil {
		query += fmt.Sprintf(" OFFSET $%d", argIdx)
		args = append(args, *offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []ClientWithAddress
	for rows.Next() {
		var client models.Client
		var address models.Address

		err := rows.Scan(
			&client.ID,
			&client.ClientName,
			&client.ClientSurname,
			&client.Birthday,
			&client.Gender,
			&client.RegistrationDate,
			&client.AddressID,
			&address.ID,
			&address.Country,
			&address.City,
			&address.Street,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, ClientWithAddress{
			Client:  client,
			Address: address,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// FindByNameSurname поиск пользователей по имени и фамилии
func (r *ClientRepository) FindByNameSurname(ctx context.Context, name, surname string) ([]ClientWithAddress, error) {
	query := `
	SELECT 
		c.id, c.client_name, c.client_surname, c.birthday, c.gender, c.registration_date, c.address_id,
		a.id, a.country, a.city, a.street
	FROM client c
	JOIN address a ON c.address_id = a.id
	WHERE c.client_name = $1 AND c.client_surname = $2
	`

	rows, err := r.db.QueryContext(ctx, query, name, surname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []ClientWithAddress
	for rows.Next() {
		var client models.Client
		var address models.Address

		err := rows.Scan(
			&client.ID,
			&client.ClientName,
			&client.ClientSurname,
			&client.Birthday,
			&client.Gender,
			&client.RegistrationDate,
			&client.AddressID,
			&address.ID,
			&address.Country,
			&address.City,
			&address.Street,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, ClientWithAddress{
			Client:  client,
			Address: address,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// GetClientByID получение пользователя по ID
func (r *ClientRepository) GetClientByID(ctx context.Context, id uuid.UUID) (*ClientWithAddress, error) {
	query := `
	SELECT 
		c.id, c.client_name, c.client_surname, c.birthday, c.gender, c.registration_date, c.address_id,
		a.id, a.country, a.city, a.street
	FROM client c
	JOIN address a ON c.address_id = a.id
	WHERE c.id = $1
	`
	var result ClientWithAddress
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&result.Client.ID,
		&result.Client.ClientName,
		&result.Client.ClientSurname,
		&result.Client.Birthday,
		&result.Client.Gender,
		&result.Client.RegistrationDate,
		&result.Client.AddressID,
		&result.Address.ID,
		&result.Address.Country,
		&result.Address.City,
		&result.Address.Street,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
