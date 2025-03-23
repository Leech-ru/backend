package ping

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Setup(group *ghttp.RouterGroup) {
	group.GET("/ping", h.Ping)
}
