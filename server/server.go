package server

import (
	"api/config"
	"api/ent"
	"api/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	DB     *ent.Client
	Config *config.Config
}

func InitializeServer() (s *Server) {
	cfg := utils.ReadEnv()
	db := InitializeDB(&cfg)
	router := gin.Default()

	s = &Server{
		DB:     db,
		Router: router,
		Config: &cfg,
	}

	InitializeRoutes(s)

	return
}

func (s *Server) Run() {
	cfg := s.Config.Server
	sConfig := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	err := http.ListenAndServe(sConfig, s.Router)
	if err != nil {
		log.Fatalf("Error starting the server: %s\n", err)
	}
}
