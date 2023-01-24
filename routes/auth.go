package routes

import (
	"api/controllers"
	"api/ent"

	"github.com/gin-gonic/gin"
)

func InitializeAuthRoutes(db *ent.Client, router *gin.Engine) {
	group := router.Group("/auth")

	auth := controllers.NewAuthController(db)

	group.POST("/login", auth.Login)
	group.POST("/register", auth.Register)
}
