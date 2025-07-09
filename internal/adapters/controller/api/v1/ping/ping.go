package ping

import (
	"Leech-ru/internal/domain/dto"
	"github.com/labstack/echo/v4"
)

// Ping check server for working.
//
// @Summary      Checking the server performance
// @Description  Checking the server performance.
// @Tags         ping
// @Accept       json
// @Produce      json
// @Success      200      {object}  dto.PingResponse "Successful check"
// @Failure      500      {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/ping [get]
func (h *handler) Ping(c echo.Context) error {
	return c.JSON(200, dto.PingResponse{
		Message: "ok",
	})
}
