package main

import (
	"Leech-ru/internal/adapters/app"
	"Leech-ru/internal/adapters/controller/api/server"
	"log"
)

func main() {
	mainApp, err := app.New()
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}
	server.Setup(mainApp)
	mainApp.Start()
}
