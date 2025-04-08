package server

import (
	"LutiLeech/internal/adapters/app"
	"LutiLeech/internal/adapters/controller/api/v1/order"
	"LutiLeech/internal/adapters/controller/api/v1/ping"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(app *app.App) {
	app.Server.Use(middleware.Logger())
	//app.Server.Use(middleware.Recover())

	addRouters(app)
}

func addRouters(app *app.App) {
	apiV1 := app.Server.Group("/api/v1")

	pingHandler := ping.NewHandler()
	pingHandler.Setup(apiV1)

	orderHandler := order.NewHandler(app.ServiceProvider.OrderService(app.Config.Mail), app.ServiceProvider.Validator())
	orderHandler.Setup(apiV1)
}
