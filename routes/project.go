package routes

import (
	"api/app"
	"api/controllers"
	"api/middleware"
)

func InitializeProjectRoutes(a *app.App) {
	group := a.Router.Group("/projects", middleware.Authentication())

	project := controllers.NewProjectController(a)

	group.GET("/", project.GetAll)
	group.POST("/", project.Create)
}
