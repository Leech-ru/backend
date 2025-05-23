package config

import (
	"github.com/spf13/viper"
	"net"
	"strconv"
)

// ServerConfig defines an interface for HTTP server configuration.
type ServerConfig interface {
	Address() string
	Port() int
	Host() string
	EnabledTLS() bool
	DevMode() bool
}

type serverConfig struct {
	host       string
	port       int
	enabledTLS bool
	tlsPort    string
	devMode    bool
}

// NewHTTPConfig initializes a new HTTP configuration from environment variables.
func NewHTTPConfig() (ServerConfig, error) {
	return &serverConfig{
		host:       viper.GetString("backend.host"),
		port:       viper.GetInt("backend.port"),
		enabledTLS: viper.GetBool("backend.tls.enabled"),
		tlsPort:    viper.GetString("backend.tls.port"),
		devMode:    viper.GetBool("backend.dev-mode"),
	}, nil
}

// Port returns port.
func (cfg *serverConfig) Port() int {
	return cfg.port
}

// Host returns host
func (cfg *serverConfig) Host() string {
	return cfg.host
}

// EnabledTLS returns true if TLS is enabled.
func (cfg *serverConfig) EnabledTLS() bool {
	return cfg.enabledTLS
}

// Address constructs and returns the full server address (host:port).
func (cfg *serverConfig) Address() string {
	if cfg.enabledTLS {
		return net.JoinHostPort(cfg.host, cfg.tlsPort)
	}
	return net.JoinHostPort(cfg.host, strconv.Itoa(cfg.port))
}

// DevMode returns true if dev mode is enabled.
func (cfg *serverConfig) DevMode() bool {
	return cfg.devMode
}
