package order

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateOrder create order and send it to email.
//
// @Summary      Create a new leech order
// @Description  Creates a new leech order and send it to email.
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateOrderRequest  true  "Leech order data"
// @Success      201      {object}  dto.CreateOrderResponse
// @Failure      400      {object}  dto.HTTPStatus "Invalid request body or validation error"
// @Failure      500      {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/order [post]
func (h *handler) CreateOrder(c echo.Context) error {
	var req dto.CreateOrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	resp, err := h.orderService.Create(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.TooMuchLeeches):
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, resp)
}
