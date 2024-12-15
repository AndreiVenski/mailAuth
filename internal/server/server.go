package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"gopkg.in/gomail.v2"
	"mailAuth/config"
	"mailAuth/pkg/logger"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	dialer *gomail.Dialer
	db     *sqlx.DB
	cfg    *config.Config
	fiber  *fiber.App
	logger logger.Logger
}

func NewServer(db *sqlx.DB, cfg *config.Config, fiberApp *fiber.App, logger logger.Logger, dialer *gomail.Dialer) *Server {
	return &Server{
		db:     db,
		cfg:    cfg,
		fiber:  fiberApp,
		logger: logger,
		dialer: dialer,
	}
}

func (s *Server) Run() error {
	go func() {
		s.logger.Infof("Server starting on port %s", s.cfg.Server.RunningPort)
		if err := s.fiber.Listen(net.JoinHostPort("", s.cfg.Server.RunningPort)); err != nil {
			s.logger.Fatalf("Error starting server:", err)
		}
	}()

	s.MapHandlers()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), time.Second)
	defer shutdown()

	s.logger.Infof("Server exited properly")
	return s.fiber.ShutdownWithContext(ctx)
}
