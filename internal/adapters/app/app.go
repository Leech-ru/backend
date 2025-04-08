package app

import (
	"LutiLeech/internal/adapters/app/service_provider"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	Server          *echo.Echo
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

	// Configure TLS if enabled
	if a.ServiceProvider.HTTPConfig().EnabledTLS() {
		// To enable TLS in Echo, you would use StartTLS() instead of Start()
		// But we'll handle that in the Start() method
		log.Println("TLS is enabled, but certificate paths are not specified.")
	}

	return a, nil
}

// Start starts the server.
func (a *App) Start() {
	addr := a.ServiceProvider.HTTPConfig().Address()
	log.Printf("Starting server on %s...\n", addr)

	if a.ServiceProvider.HTTPConfig().EnabledTLS() {
		// Example for TLS:
		// log.Fatal(a.Server.StartTLS(addr, "cert.pem", "key.pem"))
		log.Fatal(a.Server.Start(addr))
	} else {
		log.Fatal(a.Server.Start(addr))
	}
}
