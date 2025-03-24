package order

import (
	"LutiLeech/internal/adapters/config"
	"LutiLeech/pkg/mail/builder"
	"LutiLeech/pkg/mail/gmail"
)

type mailService interface {
	SendEmail(to []string, topic, body string) error
}

type Service struct {
	mailService mailService
	mailBuilder *builder.EmailBuilder
}

func NewService(config config.MailConfig) *Service {
	return &Service{
		mailService: gmail.NewMailService(config.Mail(), config.Password(), config.Host(), config.Port()),
		mailBuilder: builder.NewEmailBuilder(),
	}
}
