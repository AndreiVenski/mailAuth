package usecase

import (
	"mailAuth/config"
	"mailAuth/internal/auth"
	logger2 "mailAuth/pkg/logger"
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

func (uc *authUC) RegisterUser() {}

func (uc *authUC) VerifyCode() {}
