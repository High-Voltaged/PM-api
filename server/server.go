package server

import (
	"api/app"
	"api/utils"

	"github.com/gin-gonic/gin"
)

func InitializeApp() (a *app.App) {
	cfg := utils.ReadEnv()
	db := InitializeDB(&cfg)
	router := gin.Default()

	a = &app.App{
		DB:     db,
		Router: router,
		Config: &cfg,
	}

	InitializeRoutes(a)

	return
}
