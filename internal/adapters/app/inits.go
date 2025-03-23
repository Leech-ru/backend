package app

import (
	"LutiLeech/internal/adapters/config"
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
	//a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer() error {
	return nil
}
