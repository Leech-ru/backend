package order

import (
	"LutiLeech/internal/domain/dto"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

type orderService interface {
	Create(_ context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error)
}

type handler struct {
	orderService orderService
}

func NewHandler(
	orderService orderService,
) *handler {
	return &handler{
		orderService: orderService,
	}
}

func (h *handler) Setup(group *ghttp.RouterGroup) {
	group.POST("/order", h.CreateOrder)
}
