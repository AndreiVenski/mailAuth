package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"mailAuth/config"
	"mailAuth/internal/auth"
	"mailAuth/internal/models"
	"mailAuth/pkg/httpErrors"
	logger2 "mailAuth/pkg/logger"
	"mailAuth/pkg/utils"
	"time"
)

type authUC struct {
	authRepo  auth.Repository
	authEmail auth.Email
	logger    logger2.Logger
	cfg       *config.Config
}

func NewAuthUseCase(repository auth.Repository, email auth.Email, logger logger2.Logger, cfg *config.Config) auth.UseCase {
	return &authUC{
		authRepo:  repository,
		authEmail: email,
		logger:    logger,
		cfg:       cfg,
	}
}

func (uc *authUC) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	exists, err := uc.authRepo.IsUserExists(ctx, user.NickName, user.NickName)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, httpErrors.ExistedUserError
	}

	err = user.HashPassword(user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "authUC.RegisterUser.HashPassword")
	}
	user.UserID = uuid.New()
	createdUser, err := uc.authRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser.SanitizePassword()
	emailToken, err := utils.GenerateEmailVerificationCode(uc.cfg, createdUser)
	if err != nil {
		return nil, errors.Wrap(err, "authUC.RegisterUser.GenerateEmailVerificationCode")
	}

	if err = uc.authRepo.AddEmailCode(ctx, emailToken); err != nil {
		return nil, err
	}

	if err = uc.authEmail.SendMail(emailToken.Email, emailToken.Code); err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (uc *authUC) VerifyCode(ctx context.Context, emailCode *models.EmailCode) (*models.Tokens, error) {
	userID, err := uc.authRepo.FindEmailCodeID(ctx, emailCode.Email, emailCode.Code)
	if err != nil {
		return nil, err
	}

	if userID == uuid.Nil {
		return nil, httpErrors.NotExistedCodeError
	}
	tokens, err := utils.GenerateTokens(uc.cfg, userID)
	if err != nil {
		return nil, errors.Wrap(err, "authUC.RegisterUser.GenerateTokens")
	}

	refreshTokenRecord := &models.RefreshToken{RefreshTokenID: tokens.RefreshTokenID, UserID: userID, RefreshToken: tokens.RefreshToken, ClientInfo: " ", ExpiresAt: time.Now().Add(uc.cfg.Server.RefreshTokenExpires)}
	if err = uc.authRepo.CreateRefreshToken(ctx, refreshTokenRecord); err != nil {
		return nil, err
	}

	return tokens, nil
}
