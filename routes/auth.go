package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeAuthRoutes(router *gin.Engine) {
	group := router.Group("/auth")

	group.GET("/test", controllers.Login)
}
