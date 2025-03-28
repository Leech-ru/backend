package server

import (
	"LutiLeech/internal/adapters/app"
	"LutiLeech/internal/adapters/controller/api/logger"
	"LutiLeech/internal/adapters/controller/api/v1/order"
	"LutiLeech/internal/adapters/controller/api/v1/ping"
	"LutiLeech/internal/adapters/controller/api/validation"
)

func Setup(app *app.App) {
	logger.SetupLogger(app.Server, app.Config.Logger)
	validation.RegisterCustomValidators()
	addRouters(app)
}

func addRouters(app *app.App) {
	apiV1 := app.Server.Group("/api/v1")

	pingHandler := ping.NewHandler()
	pingHandler.Setup(apiV1)

	orderHandler := order.NewHandler(app.ServiceProvider.OrderService(app.Config.Mail))
	orderHandler.Setup(apiV1)
}
