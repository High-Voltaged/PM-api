package routes

import (
	"api/app"
	"api/controllers"
	"api/middleware"
)

func InitializeProjectRoutes(a *app.App) {
	group := a.Router.Group("/projects", middleware.Authentication(a.DB))

	project := controllers.NewProjectController(a)

	group.GET("/", project.GetAll)
	group.POST("/", project.Create)
	group.PATCH("/:id",
		middleware.ProjectAuthorAuth(a.DB),
		project.Update,
	)
	group.DELETE("/:id",
		middleware.ProjectAuthorAuth(a.DB),
		project.Delete,
	)
}
