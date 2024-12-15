package auth

import (
	"mailAuth/internal/models"

	"context"
)

type UseCase interface {
	RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
	VerifyCode(ctx context.Context, emailCode *models.EmailCode) (*models.Tokens, error)
}
