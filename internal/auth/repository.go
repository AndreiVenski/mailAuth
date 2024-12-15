package auth

import (
	"context"
	"github.com/google/uuid"
	"mailAuth/internal/models"
)

type Repository interface {
	IsUserExists(ctx context.Context, nickname, email string) (bool, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	AddEmailCode(ctx context.Context, code *models.EmailVerificationCode) error
	FindEmailCodeID(ctx context.Context, email, code string) (uuid.UUID, error)
	CreateRefreshToken(ctx context.Context, refreshTokenRecord *models.RefreshToken) error
}
