package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"mailAuth/internal/auth"
	"mailAuth/internal/models"
)

type authPostgresRepository struct {
	db *sqlx.DB
}

func NewAuthPostgresRepository(db *sqlx.DB) auth.Repository {
	return &authPostgresRepository{
		db: db,
	}
}

func (r *authPostgresRepository) IsUserExists(ctx context.Context, nickname, email string) (bool, error) {
	var exists bool
	err := r.db.GetContext(ctx, &exists, ifExistsUserQuery, nickname, email)
	if err != nil {
		return false, errors.Wrap(err, "authPostgresRepository.IsUserExists.GetContext")
	}
	return exists, nil
}

func (r *authPostgresRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	createdUser := &models.User{}
	err := r.db.QueryRowxContext(ctx, createUserQuery, user.UserID, user.Name, user.NickName, user.Email, user.Password).StructScan(createdUser)
	if err != nil {
		return nil, errors.Wrap(err, "authPostgresRepository.CreateUser.StructScan")
	}
	return createdUser, nil
}

func (r *authPostgresRepository) AddEmailCode(ctx context.Context, code *models.EmailVerificationCode) error {
	_, err := r.db.ExecContext(ctx, addEmailCode, code.UserID, code.Email, code.Code, code.ExpiresAt)
	if err != nil {
		return errors.Wrap(err, "authPostgresRepository.AddEmailCode.ExecContext")
	}
	return nil
}

func (r *authPostgresRepository) FindEmailCodeID(ctx context.Context, email, code string) (uuid.UUID, error) {
	var userID uuid.UUID
	err := r.db.QueryRowxContext(ctx, getIDAndUpdateUsedofEmailCode, email, code).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return uuid.Nil, nil
		}
		return uuid.Nil, errors.Wrap(err, "authPostgresRepository.FindEmailCodeID")
	}
	return userID, nil
}

func (r *authPostgresRepository) CreateRefreshToken(ctx context.Context, refreshTokenRecord *models.RefreshToken) error {
	_, err := r.db.ExecContext(ctx, createRefreshTokenRecord, refreshTokenRecord.RefreshTokenID, refreshTokenRecord.UserID, refreshTokenRecord.RefreshToken, refreshTokenRecord.ClientInfo, refreshTokenRecord.ExpiresAt)
	if err != nil {
		return errors.Wrap(err, "authPostgresRepository.CreateRefreshToken.ExecContext")
	}
	return nil
}
