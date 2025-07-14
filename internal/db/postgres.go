package db

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewDataBase(connString string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return db, nil
}
