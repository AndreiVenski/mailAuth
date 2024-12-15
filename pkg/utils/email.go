package utils

import (
	"crypto/rand"
	"fmt"
	"mailAuth/config"
	"mailAuth/internal/models"
	"math/big"
	"time"
)

func GenerateEmailVerificationCode(cfg *config.Config, user *models.User) (*models.EmailVerificationCode, error) {
	code := ""
	for i := 0; i < 6; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		code += fmt.Sprintf("%d", n)
	}

	return &models.EmailVerificationCode{
		UserID:    user.UserID,
		Email:     user.Email,
		Code:      code,
		ExpiresAt: time.Now().Add(cfg.Server.EmailCodeExpiresAt),
	}, nil
}
