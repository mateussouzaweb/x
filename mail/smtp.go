package mail

import (
	"net/smtp"
	"strings"
)

// SMTP struct
type SMTP struct {
	Host     string
	Port     string
	Username string
	Password string
}

// Address method
func (s *SMTP) Address() string {
	return strings.Join([]string{s.Host, s.Port}, ":")
}

// Auth method
func (s *SMTP) Auth() smtp.Auth {
	return smtp.PlainAuth("", s.Username, s.Password, s.Host)
}

// Delivery method
func (s *SMTP) Delivery(mail *Mail, content []byte) error {

	addr := s.Address()
	auth := s.Auth()
	to := []string{mail.To}

	err := smtp.SendMail(addr, auth, mail.From.Email, to, content)

	if err != nil {
		return err
	}

	return nil
}
