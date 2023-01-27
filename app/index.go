package app

import (
	"api/config"
	"api/ent"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	DB     *ent.Client
	Config *config.Config
}

func (a *App) Run() {
	cfg := a.Config.Server
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		log.Fatalf("Error starting the server: %s\n", err)
	}
}
