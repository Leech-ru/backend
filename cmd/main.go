package main

import (
	"LutiLeech/internal/adapters/app"
	"LutiLeech/internal/adapters/controller/api/server"
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
