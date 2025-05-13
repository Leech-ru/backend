package app

import (
	"Leech-ru/internal/adapters/app/service_provider"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func (a *App) initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider() error {
	a.ServiceProvider = service_provider.New()
	return nil
}

// initHTTPServer initializes the Echo server
func (a *App) initHTTPServer() error {
	e := echo.New()
	a.Server = e
	return nil
}
