package order

import (
	"LutiLeech/internal/adapters/config"
	"LutiLeech/pkg/mail"
)

type mailService interface {
	SendEmail(to []string, topic, body string) error
}

type Service struct {
	mailService mailService
}

func NewService(config config.MailConfig) *Service {
	return &Service{
		mailService: mail.NewMailService(config.Mail(), config.Password(), config.Host(), config.Port()),
	}
}
