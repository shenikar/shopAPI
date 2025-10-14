package repository

import (
	"context"
	"shopApi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SupplierRepository struct {
	db *sqlx.DB
}

func NewSupplierRepository(db *sqlx.DB) *SupplierRepository {
	return &SupplierRepository{
		db: db,
	}
}

// CreateSupplier создает нового поставщика
func (r *SupplierRepository) CreateSupplier(ctx context.Context, supplier models.Supplier) error {
	query := `INSERT INTO supplier (id, name, address_id, phone_number)
			 VALUES (:id, :name, :address_id, :phone_number)`
	_, err := r.db.NamedExecContext(ctx, query, supplier)
	return err
}

// DeleteSupplier удаление поставщика по ID
func (r *SupplierRepository) DeleteSupplier(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM supplier WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// GetAllSupplier получение всех поставщиков
func (r *SupplierRepository) GetAllSupplier(ctx context.Context) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	query := `SELECT * FROM supplier`
	err := r.db.SelectContext(ctx, &suppliers, query)
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

// GetSupplierByID получение поставщика по ID
func (r *SupplierRepository) GetSupplierByID(ctx context.Context, id uuid.UUID) (models.Supplier, error) {
	var supplier models.Supplier
	query := `SELECT * FROM supplier WHERE id = $1`
	err := r.db.GetContext(ctx, &supplier, query, id)
	if err != nil {
		return models.Supplier{}, err
	}
	return supplier, nil
}
