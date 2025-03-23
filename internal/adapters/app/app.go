package app

import (
	"LutiLeech/internal/adapters/config"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"log"
)

type App struct {
	Server *ghttp.Server
	Config *config.Config
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

// New создает новый экземпляр приложения и настраивает сервер.
func New() (*App, error) {
	a := &App{}

	err := a.initDeps()
	if err != nil {
		return nil, fmt.Errorf("new app: %w", err)
	}
	a.Server = ghttp.GetServer()

	// Настраиваем TLS, если он включен
	if a.Config.HttpServer.EnabledTLS() {
		// server.EnableHTTPS("path/to/cert.pem", "path/to/key.pem")
		log.Println("TLS is enabled, but certificate paths are not specified.")
	}
	a.Server.SetAddr(a.Config.HttpServer.Address())

	return a, nil
}

// Start запускает сервер.
func (a *App) Start() {

	log.Printf("Starting server...\n")
	a.Server.Run()
}
