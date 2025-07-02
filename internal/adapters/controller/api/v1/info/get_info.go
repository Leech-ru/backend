package info

import (
	"Leech-ru/internal/domain/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetInfo возвращает текущую информацию о корпорации.
//
// @Summary      Get corporation info
// @Description  Returns the current public corporation information: heading, description, fluid status, schedule, and links.
// @Tags         info
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.GetInfoResponse
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /info/corporation [get]
func (h *handler) GetInfo(c echo.Context) error {
	resp, err := h.infoService.Get(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, resp)
}
