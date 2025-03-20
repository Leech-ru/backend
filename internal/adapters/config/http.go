package config

import (
	"github.com/spf13/viper"
	"net"
	"strconv"
)

// HTTPConfig defines an interface for HTTP server configuration.
type HTTPConfig interface {
	Address() string
	Port() int
	Host() string
	EnabledTLS() bool
}

type httpConfig struct {
	host       string
	port       int
	enabledTLS bool
	tlsPort    string
}

// NewHTTPConfig initializes a new HTTP configuration from environment variables.
func NewHTTPConfig() (HTTPConfig, error) {
	return &httpConfig{
		host:       viper.GetString("backend.host"),
		port:       viper.GetInt("backend.port"),
		enabledTLS: viper.GetBool("backend.tls.enabled"),
		tlsPort:    viper.GetString("backend.tls.port"),
	}, nil
}

// EnabledTLS returns true if TLS is enabled.
func (cfg *httpConfig) Port() int {
	return cfg.port
}

func (cfg *httpConfig) Host() string {
	return cfg.host
}

// EnabledTLS returns true if TLS is enabled.
func (cfg *httpConfig) EnabledTLS() bool {
	return cfg.enabledTLS
}

// Address constructs and returns the full server address (host:port).
func (cfg *httpConfig) Address() string {
	if cfg.enabledTLS {
		return net.JoinHostPort(cfg.host, cfg.tlsPort)
	}
	return net.JoinHostPort(cfg.host, strconv.Itoa(cfg.port))
}
