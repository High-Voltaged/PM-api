package server

import (
	"api/app"
	"api/utils"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func InitializeApp() (a *app.App) {
	InitializeLogger()
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

func InitializeLogger() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	})
	log.SetOutput(os.Stdout)
}
