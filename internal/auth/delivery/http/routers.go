package http

import (
	"github.com/gofiber/fiber/v2"
	"mailAuth/internal/auth"
)

func MapAuthRouters(api fiber.Router, handler auth.Handler) {
	api.Post("/register", handler.Register)
	api.Post("/verify", handler.VerifyEmailCode)
}
