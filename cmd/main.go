package main

import (
	"Leech-ru/internal/adapters/app"
	"Leech-ru/internal/adapters/controller/api/server"
	"log"
)

// @title           Leech API
// @version         1.0
// @description     Backend service for Leech-ru platform.

// @contact.name   API Support
// @contact.email  mmishin2107@gmail.com

// @host      localhost:8000

// @schemes http
func main() {
	mainApp, err := app.New()
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}
	server.Setup(mainApp)
	mainApp.Start()
}
