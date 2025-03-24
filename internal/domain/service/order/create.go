package order

import (
	"LutiLeech/internal/domain/dto"
	"context"
	"sync"
)

const (
	CorporateEmail = "corporate0mail@gmail.com"
)

func (s *Service) Create(_ context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error) {
	var wg sync.WaitGroup
	var err error
	var mu sync.Mutex

	commonBody, buildErr := s.mailBuilder.
		SetRecipient("МишМиш").
		SetSubject("Купи пиявка дешева").
		SetMessage("На сайте leech.ru лучшие пиявки купи паже").
		Build()

	if buildErr != nil {
		return nil, buildErr
	}

	commonSubject := "У вас новый заказ 😇"
	recipients := []string{CorporateEmail, req.Customer}

	wg.Add(len(recipients))

	for _, email := range recipients {
		go func(email string) {
			defer wg.Done()

			if sendErr := s.mailService.SendEmail(
				[]string{email},
				commonSubject,
				commonBody,
			); sendErr != nil {
				mu.Lock()
				err = sendErr
				mu.Unlock()
			}
		}(email)
	}

	wg.Wait()
	if err != nil {
		return nil, err
	}

	return &dto.CreateOrderResponse{
		Status: "ok",
	}, nil
}
