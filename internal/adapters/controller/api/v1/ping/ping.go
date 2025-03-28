package ping

import (
	"LutiLeech/internal/domain/dto"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Ping godoc
// @Summary Server health check
// @Description Returns an "ok" message if the server is running
// @Tags Ping
// @Accept json
// @Produce json
// @Success 200 {object} dto.PingResponse
// @Router /ping [get]
func (h *handler) Ping(r *ghttp.Request) {
	r.Response.WriteStatusExit(200, dto.PingResponse{
		Message: "ok",
	})
}
