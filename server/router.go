package server

import (
	"api/app"
	"api/routes"
)

func InitializeRoutes(a *app.App) {
	routes.InitializeAuthRoutes(a)
	routes.InitializeProjectRoutes(a)
}
