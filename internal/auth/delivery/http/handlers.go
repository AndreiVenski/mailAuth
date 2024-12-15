package http

import (
	"github.com/gofiber/fiber/v2"
	"mailAuth/config"
	"mailAuth/internal/auth"
	"mailAuth/internal/models"
	"mailAuth/pkg/httpErrors"
	logger2 "mailAuth/pkg/logger"
	"mailAuth/pkg/utils"
)

type authHandler struct {
	authUC auth.UseCase
	logger logger2.Logger
	cfg    *config.Config
}

func NewAuthHandler(cfg *config.Config, logger logger2.Logger, authUC auth.UseCase) auth.Handler {
	return &authHandler{
		cfg:    cfg,
		logger: logger,
		authUC: authUC,
	}
}

func respondWithError(ctx *fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(fiber.Map{
		"error": message,
	})
}

// Register registers a new user.
// @Summary      Register a new user
// @Description  Creates a new user and sends a verification email.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.UserSwagger  true  "User Info"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{} "Invalid request data"
// @Failure      500   {object}  map[string]interface{} "Internal server error"
// @Router       /register [post]
func (h *authHandler) Register(ctx *fiber.Ctx) error {
	userInfo := &models.User{}
	err := utils.ReadFromRequest(ctx, userInfo)
	if err != nil {
		h.logger.Errorf("Register: Error with request data: %v", err)
		return respondWithError(ctx, 400, httpErrors.InvalidRequestDataError.Error())
	}

	createdUser, err := h.authUC.RegisterUser(ctx.UserContext(), userInfo)
	if err != nil {
		if httpErrors.IsUserError(err) {
			h.logger.Errorf("RegisterUser users error: %v", err.Error())
			return respondWithError(ctx, 400, err.Error())
		}
		h.logger.Errorf("RegisterUser server error: %v", err.Error())
		return respondWithError(ctx, 500, "Can't register user.")
	}

	return ctx.JSON(fiber.Map{"message": "user created successfully, please verify now your email to get tokens", "user": createdUser})
}

// VerifyEmailCode verifies the email code.
// @Summary      Verify email code
// @Description  Verifies the email code and generates tokens.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        emailCode  body      models.EmailCode  true  "Email and Code"
// @Success      200        {object}  map[string]interface{} "Tokens"
// @Failure      400        {object}  map[string]interface{} "Invalid request data"
// @Failure      500        {object}  map[string]interface{} "Internal server error"
// @Router       /verify [post]
func (h *authHandler) VerifyEmailCode(ctx *fiber.Ctx) error {
	emailCode := &models.EmailCode{}
	err := utils.ReadFromRequest(ctx, emailCode)
	if err != nil {
		h.logger.Errorf("VerifyEmailCode: Error with request data: %v", err)
		return respondWithError(ctx, 400, httpErrors.InvalidRequestDataError.Error())
	}

	tokens, err := h.authUC.VerifyCode(ctx.UserContext(), emailCode)
	if err != nil {
		if httpErrors.IsUserError(err) {
			h.logger.Errorf("VerifyEmailCode users error: %v", err.Error())
			return respondWithError(ctx, 400, err.Error())
		}
		h.logger.Errorf("VerifyEmailCode server error: %v", err.Error())
		return respondWithError(ctx, 500, "Can't verify email code.")
	}

	return ctx.JSON(fiber.Map{"message": "code is available", "tokens": tokens})
}
