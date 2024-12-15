package auth

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Register(ctx *fiber.Ctx) error
	VerifyEmailCode(ctx *fiber.Ctx) error
}
