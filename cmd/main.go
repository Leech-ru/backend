package main

import (
	"Leech-ru/internal/adapters/app"
	"Leech-ru/internal/adapters/controller/api/server"
	"log"
)

// @title           Leech API
// @version         1.0
// @description     Backend service for Leech-ru platform. Uses cookie-based authentication with HttpOnly tokens.

// @contact.name    API Support
// @contact.email   mmishin2107@gmail.com

// @host            пиявкипобеда.рф
// @schemes         https

// @securityDefinitions.apikey  CookieAuth
// @in                          header
// @name                        Cookie
// @description                 Authentication via HttpOnly cookies. System uses two cookies:\n- `user_auth_access_token` (short-lived)\n- `user_auth_refresh_token` (long-lived)\n\nAll protected endpoints require valid cookies to be automatically sent by browser.

// @tag.name        ping
// @tag.description The main check of server performance

// @tag.name        order
// @tag.description Leech order operations

// @tag.name        user
// @tag.description User authentication and management

// @tag.name        auth
// @tag.description Work with authorization

// @tag.name        cosmetics
// @tag.description Cosmetics view and management

// @tag.name        image
// @tag.description Image view and management

// @tag.name        category
// @tag.description Category view and management

// @tag.name        info
// @tag.description Information about the center

// @tag.name        partner
// @tag.description Information about center's partners

// @tag.name        news
// @tag.description News view and management

// @tag.name        main
// @tag.description information about main page

func main() {
	mainApp, err := app.New()
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}
	server.Setup(mainApp)
	mainApp.Start()
}
