package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

// mailService представляет сервис для отправки писем
type mailService struct {
	username string
	password string
	host     string
	port     int
}

// NewMailService создает новый экземпляр mailService
func NewMailService(username, password, host string, port int) mailService {
	return mailService{
		username: username,
		password: password,
		host:     host,
		port:     port,
	}
}

// SendEmail отправляет письмо на указанный адрес
func (s mailService) SendEmail(to []string, topic, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.username)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", topic)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(s.host, s.port, s.username, s.password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
