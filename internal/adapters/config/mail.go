package config

import (
	"github.com/spf13/viper"
)

// HTTPConfig defines an interface for HTTP server configuration.
type MailConfig interface {
	Host() string
	Port() int
	Mail() string
	Password() string
}

type mailConfig struct {
	host     string
	port     int
	mail     string
	password string
}

// NewMailConfig initializes a new Mail configuration from environment variables.
func NewMailConfig() (MailConfig, error) {
	return &mailConfig{
		host:     viper.GetString("service.mail.host"),
		port:     viper.GetInt("service.mail.port"),
		mail:     viper.GetString("service.mail.mail"),
		password: viper.GetString("service.mail.password"),
	}, nil
}

func (cfg *mailConfig) Host() string {
	return cfg.host
}

func (cfg *mailConfig) Port() int {
	return cfg.port
}

func (cfg *mailConfig) Mail() string {
	return cfg.mail
}

func (cfg *mailConfig) Password() string {
	return cfg.password
}
