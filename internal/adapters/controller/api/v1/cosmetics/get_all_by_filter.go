package cosmetics

import (
	"Leech-ru/internal/domain/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) GetAllByFilter(c echo.Context) error {
	var req dto.GetAllByFilterCosmeticsRequest

	if err := h.formDecoder.Decode(&req, c.QueryParams()); err != nil {
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
	resp, err := h.cosmeticsService.GetAllByFilter(c.Request().Context(), &req)
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, resp)
}
