package service_provider

import (
	"Leech-ru/internal/adapters/config"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/service/order"
	"context"
)

type orderService interface {
	Create(ctx context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error)
}

func (s *ServiceProvider) OrderService(config config.MailConfig) orderService {
	if s.orderService == nil {
		s.orderService = order.NewService(config)
	}
	return s.orderService
}
