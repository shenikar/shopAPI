package models

import "github.com/google/uuid"

type Image struct {
	ID        uuid.UUID `db:"id"`
	Image     []byte    `db:"image"`
	ProductID uuid.UUID `db:"product_id"`
}
