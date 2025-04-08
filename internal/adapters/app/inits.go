package app

import (
	"LutiLeech/internal/adapters/app/service_provider"
	"LutiLeech/internal/adapters/config"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
)

func (a *App) initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	a.Config = cfg
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
