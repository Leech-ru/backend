package ping

import (
	"LutiLeech/internal/adapters/app"
	"LutiLeech/internal/domain/dto"
	"github.com/gogf/gf/v2/net/ghttp"
)

type PingHandler struct{}

func NewHandler(app *app.App) *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Ping(r *ghttp.Request) {
	r.Response.WriteJson(dto.Ping{Message: "ok"})
}

func (h *PingHandler) Setup(group *ghttp.RouterGroup) {
	group.GET("/ping", h.Ping)
}
