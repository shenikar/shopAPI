package repository

import (
	"context"
	"database/sql"
	"shopApi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// CreateProduct создает новый продукт в бд
func (r *ProductRepository) CreateProduct(ctx context.Context, product models.Product) error {
	query := `INSERT INTO product (id, name, category, price, available_stock, last_update_date, supplier_id, image_id
	) VALUES (:id, :name, :category, :price, :available_stock, :last_update_date, :supplier_id, :image_id)
	`

	_, err := r.db.NamedExecContext(ctx, query, product)
	return err
}

// GetProductByID получает продукт по ID
func (r *ProductRepository) GetProductByID(ctx context.Context, id uuid.UUID) (models.Product, error) {
	var product models.Product
	query := `SELECT * FROM product WHERE id = $1`
	err := r.db.GetContext(ctx, &product, query, id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// GetAllAvailableProducts получение всех доступных товаров
func (r *ProductRepository) GetAllAvailableProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	query := `SELECT * FROM product WHERE available_stock > 0`
	err := r.db.SelectContext(ctx, &products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Delete удаление товара по ID
func (r *ProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM product WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// DecreaseStock Уменьшение количества на складе
func (r *ProductRepository) DecreaseStock(ctx context.Context, id uuid.UUID, quantity int) error {
	query := `UPDATE product 
			 SET available_stock = available_stock - $1 
			 WHERE id = $2 AND available_stock >= $1
			 `
	res, err := r.db.ExecContext(ctx, query, quantity, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
