package repository

import (
	"context"

	"github.com/shenikar/shopAPI/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AddressRepository struct {
	db *sqlx.DB
}

func NewAddressRepository(db *sqlx.DB) *AddressRepository {
	return &AddressRepository{
		db: db,
	}
}

func (r *AddressRepository) Create(ctx context.Context, address *models.Address) (uuid.UUID, error) {
	address.ID = uuid.New()
	query := `INSERT INTO address (id, country, city, street) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, address.ID, address.Country, address.City, address.Street)
	if err != nil {
		return uuid.Nil, err
	}
	return address.ID, nil
}

func (r *AddressRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Address, error) {
	var address models.Address
	query := `SELECT id, country, city, street FROM address WHERE id = $1`
	err := r.db.GetContext(ctx, &address, query, id)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *AddressRepository) Update(ctx context.Context, address *models.Address) error {
	query := `UPDATE address SET country = $1, city = $2, street = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, address.Country, address.City, address.Street, address.ID)
	return err
}
