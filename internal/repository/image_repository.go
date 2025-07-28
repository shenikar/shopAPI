package repository

import (
	"context"
	"database/sql"
	"shopApi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ImageRepository struct {
	db *sqlx.DB
}

func NewImageRepository(db *sqlx.DB) *ImageRepository {
	return &ImageRepository{
		db: db,
	}
}

// CreateImage создает новую запись изображения в базе данных
func (r *ImageRepository) CreateImage(ctx context.Context, image *models.Image) error {
	query := `INSERT INTO images (id, image, product_id) VALUES (:id, :image, :product_id)`
	_, err := r.db.NamedExecContext(ctx, query, image)
	return err
}

// UpdateImage обновляет изображение по ID
func (r *ImageRepository) UpdateImage(ctx context.Context, id uuid.UUID, newImage []byte) error {
	query := `UPDATE images SET image = $1 WHERE id = $2`
	res, err := r.db.ExecContext(ctx, query, newImage, id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// GetImageByID получает изображение по ID
func (r *ImageRepository) GetImageByID(ctx context.Context, id uuid.UUID) (*models.Image, error) {
	var image models.Image
	query := `SELECT * FROM images WHERE id = $1`
	err := r.db.GetContext(ctx, &image, query, id)
	if err != nil {
		return nil, err
	}
	return &image, nil
}

// DeleteImage удаляет изображение по ID
func (r *ImageRepository) DeleteImage(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM images WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// GetByProductID получает все изображения по ID товара
func (r *ImageRepository) GetByProductID(ctx context.Context, productID uuid.UUID) ([]models.Image, error) {
	var images []models.Image
	query := `SELECT * FROM images WHERE product_id = $1`
	err := r.db.SelectContext(ctx, &images, query, productID)
	if err != nil {
		return nil, err
	}
	return images, nil
}
