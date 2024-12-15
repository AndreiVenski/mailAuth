package email

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
	"mailAuth/config"
	"mailAuth/internal/auth"
)

type authEmail struct {
	cfg    *config.Config
	dialer *gomail.Dialer
}

func NewAuthEmail(cfg *config.Config, dialer *gomail.Dialer) auth.Email {
	return &authEmail{
		cfg:    cfg,
		dialer: dialer,
	}
}

func (e *authEmail) SendMail(toEmail string, code string) error {
	message := e.generateMail(toEmail, code)

	if err := e.dialer.DialAndSend(message); err != nil {
		return errors.Wrap(err, "authEmail.SendMail.DialAndSend")
	}
	return nil
}

func (e *authEmail) generateMail(toEmail, code string) *gomail.Message {
	message := gomail.NewMessage()
	body := fmt.Sprintf(codeMessage, code)
	message.SetHeader("From", e.cfg.SMTPServer.Email)
	message.SetHeader("To", toEmail)
	message.SetHeader("Subject", codeSubject)
	message.SetBody("text/plain", body)

	return message
}
