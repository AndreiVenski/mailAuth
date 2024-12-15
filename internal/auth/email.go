package auth

type Email interface {
	SendMail(toEmail string, code string) error
}
