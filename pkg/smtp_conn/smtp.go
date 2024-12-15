package smtp_conn

import (
	"github.com/gofiber/fiber/v2/log"
	"gopkg.in/gomail.v2"
	"mailAuth/config"
)

func NewSMTPDialer(cfg *config.Config) (*gomail.Dialer, error) {
	dialer := gomail.NewDialer(cfg.SMTPServer.Host, cfg.SMTPServer.Port, cfg.SMTPServer.Username, cfg.SMTPServer.Password)

	err := PingSMTP(dialer)
	if err != nil {
		return nil, err
	}
	return dialer, nil
}

func PingSMTP(dialer *gomail.Dialer) error {
	conn, err := dialer.Dial()
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			log.Errorf("Failed to close connection: %v\n", closeErr)
		}
	}()

	return nil
}
