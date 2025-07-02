package server

import (
	_ "Leech-ru/docs"
	"Leech-ru/internal/adapters/app"
	"Leech-ru/internal/adapters/controller/api/middleware/auth"
	"Leech-ru/internal/adapters/controller/api/middleware/role"
	"Leech-ru/internal/adapters/controller/api/v1/cosmetics"
	"Leech-ru/internal/adapters/controller/api/v1/info"
	"Leech-ru/internal/adapters/controller/api/v1/order"
	"Leech-ru/internal/adapters/controller/api/v1/ping"
	"Leech-ru/internal/adapters/controller/api/v1/token"
	"Leech-ru/internal/adapters/controller/api/v1/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"io"
	"net/http"
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
	server := app.Server
	serviceProvider := app.ServiceProvider

	authMiddleware := auth.NewAuthMiddleware(serviceProvider.TokenService())
	roleMiddleware := role.NewRoleMiddleware(serviceProvider.UserService())

	apiV1 := server.Group("/api/v1")

	apiV1.GET("/swagger/*", echoSwagger.WrapHandler)
	apiV1.GET("/swagger", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/api/v1/swagger/index.html")
	})

	pingHandler := ping.NewHandler()
	pingHandler.Setup(apiV1)

	refreshTokenHandler := token.NewHandler(serviceProvider.TokenService(), serviceProvider.JWTConfig(), serviceProvider.ServerConfig(), serviceProvider.Validator(), serviceProvider.Decoder())
	refreshTokenHandler.Setup(apiV1)

	orderHandler := order.NewHandler(serviceProvider.OrderService(serviceProvider.MailConfig()), serviceProvider.Validator())
	orderHandler.Setup(apiV1)

	userHandler := user.NewHandler(serviceProvider.UserService(), serviceProvider.JWTConfig(), serviceProvider.ServerConfig(), authMiddleware, roleMiddleware, serviceProvider.Validator(), serviceProvider.Decoder())
	userHandler.Setup(apiV1)

	cosmeticsHandler := cosmetics.NewHandler(serviceProvider.CosmeticsService(), authMiddleware, roleMiddleware, serviceProvider.Validator(), serviceProvider.Decoder())
	cosmeticsHandler.Setup(apiV1)

	infoHandler := info.NewHandler(serviceProvider.InfoService(), authMiddleware, roleMiddleware, serviceProvider.Validator())
	infoHandler.Setup(apiV1)
}
