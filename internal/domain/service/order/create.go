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

	wg.Add(2)

	go func() {
		defer wg.Done()

		recipients := []string{CorporateEmail}
		topic := "!АЛЕРТ! НОВЫЙ ЗАКАЗ"
		body := "СРОЧНО СРОЧНО ЗАКАЗАЛИ ПИЯВОК\nААААААААААААА\nСРОЧНО НАДО ИХ ОТПРАВИТЬ\nВПЕРЕД РАБОТАТЬ СОЛНЦЕ ЕЩЕ ВЫСОКО"

		if sendErr := s.mailService.SendEmail(recipients, topic, body); sendErr != nil {
			mu.Lock()
			err = sendErr
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()

		recipients := []string{req.Customer}
		topic := "У вас новый заказ😇"
		body := "Спасибо, что выбрали нас\nМы уже начали собирать ваш заказ\nВ скором времени мы свяжемся с вами для обсуждения подробностей😊"

		if sendErr := s.mailService.SendEmail(recipients, topic, body); sendErr != nil {
			mu.Lock()
			err = sendErr
			mu.Unlock()
		}
	}()

	wg.Wait()
	if err != nil {
		return nil, err
	}
	return &dto.CreateOrderResponse{
		Status: "ok",
	}, nil
}
