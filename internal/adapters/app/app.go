package app

import (
	"LutiLeech/internal/adapters/app/service_provider"
	"LutiLeech/internal/adapters/config"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"log"
)

type App struct {
	Server          *ghttp.Server
	Config          *config.Config
	ServiceProvider *service_provider.ServiceProvider
}

// initDeps initializes application dependencies
func (a *App) initDeps() error {
	inits := []func() error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}
	for _, f := range inits {
		err := f()
		if err != nil {
			return fmt.Errorf("init deps: %w", err)
		}
	}
	return nil
}

// New creates a new copy of the application and sets the server.
func New() (*App, error) {
	a := &App{}

	err := a.initDeps()
	if err != nil {
		return nil, fmt.Errorf("new app: %w", err)
	}
	// Настраиваем TLS, если он включен
	if a.Config.HttpServer.EnabledTLS() {
		// server.EnableHTTPS("path/to/cert.pem", "path/to/key.pem")
		log.Println("TLS is enabled, but certificate paths are not specified.")
	}
	a.Server.SetAddr(a.Config.HttpServer.Address())

	return a, nil
}

// Start start server.
func (a *App) Start() {

	log.Printf("Starting server...\n")
	a.Server.Run()
}
