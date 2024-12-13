package repository

import (
	"github.com/jmoiron/sqlx"
	"mailAuth/internal/auth"
)

type authPostgresRepository struct {
	db *sqlx.DB
}

func NewAuthPostgresRepository(db *sqlx.DB) auth.Repository {
	return &authPostgresRepository{
		db: db,
	}
}

func (r *authPostgresRepository) IsUserExists() (bool, error) {
	return false, nil
}
