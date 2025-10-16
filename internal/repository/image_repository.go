package repository

import (
	"context"
	"database/sql"

	"github.com/shenikar/shopAPI/internal/domain/models"

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
	query := `INSERT INTO images (id, image) VALUES (:id, :image)`
	_, err := r.db.NamedExecContext(ctx, query, image)
	return err
}

// UpdateImage обновление изображения по ID
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

// GetByProductID Получение изображения по ID товара
func (r *ImageRepository) GetByProductID(ctx context.Context, productID uuid.UUID) (*models.Image, error) {
	var image models.Image
	query := `SELECT i.* FROM images i
			 JOIN product p ON p.image_id = i.id WHERE p.id = $1
			 `
	err := r.db.GetContext(ctx, &image, query, productID)
	if err != nil {
		return nil, err
	}
	return &image, nil
}

// UpdateProductImageID обновляет поле image_id у товара
func (r *ImageRepository) UpdateProductImageID(ctx context.Context, productID, imageID uuid.UUID) error {
	query := `UPDATE product SET image_id = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, imageID, productID)
	return err
}

// ClearProductImageIDByImageID очищает ссылку image_id у товаров, которые ссылаются на данное изображение
func (r *ImageRepository) ClearProductImageIDByImageID(ctx context.Context, imageID uuid.UUID) error {
	query := `UPDATE product SET image_id = NULL WHERE image_id = $1`
	_, err := r.db.ExecContext(ctx, query, imageID)
	return err
}
