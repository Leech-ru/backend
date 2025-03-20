package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config представляет собой общую структуру конфигурации,
// которая содержит все загруженные микроконфиги.
type Config struct {
	Database   PGConfig
	Logger     LoggerConfig
	HttpServer HTTPConfig
	Mail       MailConfig
}

// New создает и возвращает новый объект Config, загружая все микроконфиги.
func New() (*Config, error) {
	initConfig()

	dbConfig, err := NewPGConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load database config: %w", err)
	}

	loggerConfig, err := NewLoggerConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load logger config: %w", err)
	}

	serverConfig, err := NewHTTPConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load server config: %w", err)
	}

	mailCfg, err := NewMailConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load mail config: %w", err)
	}

	return &Config{
		Database:   dbConfig,
		Logger:     loggerConfig,
		HttpServer: serverConfig,
		Mail:       mailCfg,
	}, nil
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
