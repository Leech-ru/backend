package order

import (
	"LutiLeech/internal/domain/common/errorz"
	"LutiLeech/internal/domain/dto"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// CreateOrder godoc
// @Summary Create a new order
// @Description Processes an order request and returns order details if successful
// @Tags Order
// @Accept json
// @Produce json
// @Param request body dto.CreateOrderRequest true "Order request payload"
// @Success 200 {object} dto.CreateOrderResponse
// @Failure 400 {object} dto.HTTPStatus "Bad request"
// @Failure 500 {object} dto.HTTPStatus "Internal server error"
// @Router /order [post]
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
		if err == errorz.TooMuchLeeches {
			r.Response.WriteStatusExit(400, dto.HTTPStatus{
				Code:    400,
				Message: err.Error(),
			})
			return
		}
		r.Response.WriteStatusExit(500, dto.HTTPStatus{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	r.Response.WriteJson(resp)
}
