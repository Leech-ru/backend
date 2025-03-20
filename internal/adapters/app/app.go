package app

import (
	"LutiLeech/internal/adapters/config"
	"github.com/gogf/gf/v2/net/ghttp"
	"log"
)

type App struct {
	Server *ghttp.Server
}

// New создает новый экземпляр приложения и настраивает сервер.
func New() *App {

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	server := ghttp.GetServer()

	server.SetAddr(cfg.HttpServer.Address())

	// Настраиваем TLS, если он включен
	if cfg.HttpServer.EnabledTLS() {
		// server.EnableHTTPS("path/to/cert.pem", "path/to/key.pem")
		log.Println("TLS is enabled, but certificate paths are not specified.")
	}

	return &App{
		Server: server,
	}
}

// Start запускает сервер.
func (a *App) Start() {
	log.Printf("Starting server...\n")
	a.Server.Run()
}
