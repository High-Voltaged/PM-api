package routes

import (
	"api/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(cfg *config.Config) {
	r := gin.Default()

	InitializeAuthRoutes(r)

	serverConfig := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	err := r.Run(serverConfig)
	if err != nil {
		log.Fatalf("Error starting the server: %s\n", err)
	}
}
