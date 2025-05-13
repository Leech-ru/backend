package order

import (
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/internal/domain/dto"
	"context"
	"github.com/labstack/echo/v4"
)

type orderService interface {
	Create(_ context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error)
}

type handler struct {
	orderService orderService
	validator    *validator.Validator
}

func NewHandler(
	orderService orderService,
	validator *validator.Validator,

) *handler {
	return &handler{
		orderService: orderService,
		validator:    validator,
	}
}

func (h *handler) Setup(router *echo.Group) {
	router.POST("/order", h.CreateOrder)
}
