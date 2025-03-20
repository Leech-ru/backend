package main

import (
	"LutiLeech/internal/adapters/app"
	"LutiLeech/internal/adapters/controller/api/server"
)

func main() {
	mainApp := app.New()
	server.Setup(mainApp)

	mainApp.Start()
}
