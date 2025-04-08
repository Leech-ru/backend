package server

import (
	"LutiLeech/internal/adapters/app"
	"LutiLeech/internal/adapters/controller/api/v1/order"
	"LutiLeech/internal/adapters/controller/api/v1/ping"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
)

func Setup(app *app.App) {
	app.Server.Logger.SetOutput(io.Discard)
	app.Server.HideBanner = true
	app.Server.Debug = false

	//app.Server.Use(middleware.Recover())
	app.Server.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		HandleError: true,
		LogError:    true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				app.ServiceProvider.Logger().Infow("request completed",
					"ip", v.RemoteIP,
					"method", v.Method,
					"uri", v.URI,
					"status", v.Status,
				)
			} else {
				app.ServiceProvider.Logger().Errorw("request failed",
					"ip", v.RemoteIP,
					"method", v.Method,
					"uri", v.URI,
					"status", v.Status,
					"error", v.Error.Error(),
				)
			}
			return nil
		},
	}))

	addRouters(app)
}

func addRouters(app *app.App) {
	apiV1 := app.Server.Group("/api/v1")

	pingHandler := ping.NewHandler()
	pingHandler.Setup(apiV1)

	orderHandler := order.NewHandler(app.ServiceProvider.OrderService(app.ServiceProvider.MailConfig()), app.ServiceProvider.Validator())
	orderHandler.Setup(apiV1)
}
