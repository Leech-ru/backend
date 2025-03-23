package ping

import (
	"LutiLeech/internal/domain/dto"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (h *handler) Ping(r *ghttp.Request) {
	r.Response.WriteStatusExit(200, dto.PingResponse{
		Message: "ok",
	})
}
