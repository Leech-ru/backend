package ping

import (
	"github.com/labstack/echo/v4"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Setup(router *echo.Group) {
	router.GET("/ping", h.Ping)
}
