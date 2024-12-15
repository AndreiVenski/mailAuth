package main

import (
	_ "mailAuth/api/doc"
	"mailAuth/config"
	server2 "mailAuth/internal/server"
	"mailAuth/pkg/db/postgres_conn"
	"mailAuth/pkg/db/test_data_script"
	logger2 "mailAuth/pkg/logger"
	"mailAuth/pkg/smtp_conn"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/pressly/goose"
)

// @title Auth Api
// @version 1.0
// @description This is API for register user and verify email code
// @contact.name Andrei Venski
// @contact.url https://github.com/andrew967
// @contact.email venskiandrei32@gmail.com
// @BasePath /v1/auth
func main() {

	logger := logger2.NewApiLogger()
	logger.InitLogger()

	cfg, err := config.InitConfig(".env")
	if err != nil {
		logger.Fatalf("Fatal init config: %v", err)
		os.Exit(1)
	}

	db, err := postgres_conn.NewPsqlDB(cfg)
	if err != nil {
		logger.Error("DB init failed", err)
		os.Exit(1)
	}
	defer db.Close()

	dialer, err := smtp_conn.NewSMTPDialer(cfg)
	if err != nil {
		logger.Error("Dialer init failed", err)
		os.Exit(1)
	}

	if err = goose.Up(db.DB, "migrations"); err != nil {
		logger.Error("Migrations up failed", err)
		os.Exit(1)
	}

	if cfg.Test.AddTestDataToDB {
		if err = test_data_script.ExecuteSQLFile(db.DB, cfg.Test.TestDataScriptSource); err != nil {
			logger.Error("Test data add failed", err)
			os.Exit(1)
		}
	}

	fiberApp := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: false,
	})

	server := server2.NewServer(db, cfg, fiberApp, logger, dialer)

	if err = server.Run(); err != nil {
		os.Exit(0)
	}

}
