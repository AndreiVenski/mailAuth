package models

import (
	"github.com/google/uuid"
	"time"
)

type Tokens struct {
	RefreshToken   string    `json:"refresh_token"`
	AccessToken    string    `json:"access_token"`
	UserID         uuid.UUID `json:"user_id"`
	RefreshTokenID uuid.UUID `json:"refresh_token_id"`
}

type RefreshToken struct {
	RefreshToken   string        `json:"refresh_token" db:"token"`
	UserID         uuid.UUID     `json:"user_id" db:"user_id"`
	RefreshTokenID uuid.UUID     `json:"refresh_token_id" db:"id"`
	ClientInfo     string        `json:"client_info" db:"client_info"`
	ExpiresAt      time.Duration `json:"expires_at" db:"expires_at"`
}
