package utils

import (
	"crypto/rand"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"mailAuth/config"
	"mailAuth/internal/models"
	"time"
)

type TokenClaims struct {
	UserID         uuid.UUID
	refreshTokenID uuid.UUID
	jwt.RegisteredClaims
}

func GenerateTokens(cfg *config.Config, userID uuid.UUID) (*models.Tokens, error) {
	refreshTokenID := uuid.New()
	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	accessToken, err := GenerateAccessToken(cfg, userID, refreshTokenID)
	if err != nil {
		return nil, err
	}

	return &models.Tokens{
		RefreshToken:   refreshToken,
		AccessToken:    accessToken,
		RefreshTokenID: refreshTokenID,
		UserID:         userID,
	}, nil
}

func GenerateRefreshToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	return string(tokenBytes), nil
}

func GenerateAccessToken(cfg *config.Config, userID, refreshTokenID uuid.UUID) (string, error) {
	claims := TokenClaims{
		UserID:         userID,
		refreshTokenID: refreshTokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.Server.AccessTokenExpires)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(cfg.Server.JWTSecret))

	return tokenString, err
}
