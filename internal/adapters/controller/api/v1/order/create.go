package order

import (
	"LutiLeech/internal/domain/dto"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (h *handler) CreateOrder(r *ghttp.Request) {
	var req dto.CreateOrderRequest

	if err := r.Parse(&req); err != nil {
		r.Response.WriteStatusExit(400, dto.HTTPStatus{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	if err := g.Validator().Data(req).Run(r.Context()); err != nil {
		r.Response.WriteStatusExit(400, dto.HTTPStatus{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.orderService.Create(r.Context(), &req)
	if err != nil {
		r.Response.WriteStatusExit(500, dto.HTTPStatus{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	r.Response.WriteJson(resp)
}
