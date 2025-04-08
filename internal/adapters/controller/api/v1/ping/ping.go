package ping

import (
	"LutiLeech/internal/domain/dto"
	"github.com/labstack/echo/v4"
)

func (h *handler) Ping(c echo.Context) error {
	return c.JSON(200, dto.PingResponse{
		Message: "ok",
	})
}
