package order

import (
	"Leech-ru/internal/adapters/config"
	"Leech-ru/internal/domain/service/price"
	"Leech-ru/pkg/mail/builder"
	"Leech-ru/pkg/mail/gmail"
)

type mailService interface {
	SendEmail(to []string, topic, body string) error
}

type priceService interface {
	CalculateTotal(size1, size2, size3, packageType int) (float64, error)
}

type Service struct {
	priceService priceService
	mailService  mailService
	mailBuilder  *builder.EmailBuilder
}

func NewService(config config.MailConfig) *Service {
	return &Service{
		priceService: price.NewPriceService(),
		mailService:  gmail.NewMailService(config.Mail(), config.Password(), config.Host(), config.Port()),
		mailBuilder:  builder.NewEmailBuilder(),
	}
}
