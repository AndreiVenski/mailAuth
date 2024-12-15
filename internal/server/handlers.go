package server

import (
	"github.com/gofiber/swagger"
	"mailAuth/internal/auth/delivery/http"
	"mailAuth/internal/auth/email"
	"mailAuth/internal/auth/repository"
	"mailAuth/internal/auth/usecase"
)

func (s *Server) MapHandlers() {
	authRepository := repository.NewAuthPostgresRepository(s.db)
	authEmail := email.NewAuthEmail(s.cfg, s.dialer)

	authUseCase := usecase.NewAuthUseCase(authRepository, authEmail, s.logger, s.cfg)

	authHandler := http.NewAuthHandler(s.cfg, s.logger, authUseCase)

	s.fiber.Get("/swagger/*", swagger.HandlerDefault)

	auth := s.fiber.Group("/v1/auth")
	http.MapAuthRouters(auth, authHandler)
}
