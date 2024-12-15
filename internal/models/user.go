package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID   uuid.UUID `json:"user_id" db:"id" validate:"omitempty"`
	Name     string    `json:"name" db:"name" validate:"omitempty,max=255"`
	NickName string    `json:"nickname" db:"nickname" validate:"required,alphanum,max=255"`
	Email    string    `json:"email" db:"email" validate:"required,email,max=255"`
	Password string    `json:"password" db:"password_hash" validate:"required,gte=6"`
}

func (u *User) HashPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return nil
}

func (u *User) ComparePasswords(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) SanitizePassword() {
	u.Password = ""
}
