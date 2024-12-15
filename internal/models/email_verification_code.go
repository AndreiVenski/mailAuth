package models

import (
	"github.com/google/uuid"
	"time"
)

type EmailVerificationCode struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Email     string    `json:"email" db:"email"`
	Code      string    `json:"code" db:"code"`
	ExpiresAt time.Time `json:"expires_at" db:"expire_at"`
}

type EmailCode struct {
	Email string `json:"email" db:"email" validate:"required"`
	Code  string `json:"code" db:"code" validate:"required"`
}
