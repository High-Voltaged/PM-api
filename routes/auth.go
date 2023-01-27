package routes

import (
	"api/app"
	"api/controllers"
)

func InitializeAuthRoutes(a *app.App) {
	group := a.Router.Group("/auth")

	auth := controllers.NewAuthController(a)

	group.POST("/login", auth.Login)
	group.POST("/register", auth.Register)
	group.POST("/forgot-password", auth.ForgotPassword)
	group.POST("/reset-password/:resetToken", auth.ResetPassword)
}
