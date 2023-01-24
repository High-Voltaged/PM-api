package server

import (
	"api/routes"
)

func InitializeRoutes(server *Server) {
	routes.InitializeAuthRoutes(server.DB, server.Router)
}
