package repository

import (
	"context"
	"fmt"
	"shopApi/internal/domain/models"
	"shopApi/internal/dto"

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
func (r *ClientRepository) CreateClient(ctx context.Context, req dto.CreateClientDTO) (models.Client, models.Address, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Client{}, models.Address{}, err
	}
	defer tx.Rollback()

	var addressID uuid.UUID
	err = tx.QueryRowContext(ctx, `
		INSERT INTO address (country, city, street)
		VALUES ($1, $2, $3) RETURNING id
	`, req.Address.Country, req.Address.City, req.Address.Street).Scan(&addressID)
	if err != nil {
		return models.Client{}, models.Address{}, err
	}

	var address models.Address
	err = tx.QueryRowContext(ctx, `
	SELECT id, country, city, street
	FROM address
	WHERE id = $1
`, addressID).Scan(&address.ID, &address.Country, &address.City, &address.Street)
	if err != nil {
		return models.Client{}, models.Address{}, err
	}

	var client models.Client
	clientID := uuid.New()
	err = tx.QueryRowContext(ctx, `
	INSERT INTO client (id, client_name, client_surname, birthday, gender, registration_date, address_id)
	VALUES ($1, $2, $3, $4, $5, NOW(), $6)
	RETURNING id, client_name, client_surname, birthday, gender, registration_date, address_id
`, clientID, req.ClientName, req.ClientSurname, req.Birthday, req.Gender, addressID).
		Scan(&client.ID, &client.ClientName, &client.ClientSurname, &client.Birthday, &client.Gender, &client.RegistrationDate, &client.AddressID)

	if err != nil {
		return models.Client{}, models.Address{}, err
	}

	if err := tx.Commit(); err != nil {
		return models.Client{}, models.Address{}, err
	}

	return client, address, nil
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

// UpdateAddress обновление адреса пользователя
func (r *ClientRepository) UpdateAddress(ctx context.Context, ID uuid.UUID, req dto.CreateAddressDTO) (models.Client, models.Address, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Client{}, models.Address{}, err
	}
	defer tx.Rollback()

	var addressID uuid.UUID
	err = tx.QueryRowContext(ctx, `
		SELECT address_id FROM client WHERE id = $1
	`, ID).Scan(&addressID)
	if err != nil {
		return models.Client{}, models.Address{}, err
	}

	_, err = tx.ExecContext(ctx, `
		UPDATE address SET country = $1, city = $2, street = $3 WHERE id = $4
	`, req.Country, req.City, req.Street, addressID)
	if err != nil {
		return models.Client{}, models.Address{}, err
	}

	var client models.Client
	err = tx.QueryRowContext(ctx, `
		SELECT id, client_name, client_surname, birthday, gender, registration_date, address_id
		FROM client WHERE id = $1
	`, ID).Scan(&client.ID, &client.ClientName, &client.ClientSurname, &client.Birthday, &client.Gender, &client.RegistrationDate, &client.AddressID)
	if err != nil {
		return models.Client{}, models.Address{}, err
	}

	var address models.Address
	err = tx.QueryRowContext(ctx, `
		SELECT id, country, city, street FROM address WHERE id = $1
	`, addressID).Scan(&address.ID, &address.Country, &address.City, &address.Street)
	if err != nil {
		return models.Client{}, models.Address{}, err
	}

	if err := tx.Commit(); err != nil {
		return models.Client{}, models.Address{}, err
	}

	return client, address, nil
}

func (r *ClientRepository) CreateAddress(ctx context.Context, address models.Address) (uuid.UUID, error) {
	address.ID = uuid.New()

	query := `INSERT INTO address (id, country, city, street) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, address.ID, address.Country, address.City, address.Street)
	if err != nil {
		return uuid.Nil, err
	}

	return address.ID, nil
}

func (r *ClientRepository) GetAddressByID(ctx context.Context, id uuid.UUID) (models.Address, error) {
	var address models.Address
	query := `SELECT id, country, city, street FROM address WHERE id = $1`
	err := r.db.GetContext(ctx, &address, query, id)
	return address, err
}
