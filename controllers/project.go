package controllers

import (
	"api/app"
	req "api/requests"
	"api/response"
	"api/services"
	"api/tokens"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	service *services.ProjectService
}

func NewProjectController(a *app.App) *ProjectController {
	return &ProjectController{service: services.NewProjectService(a)}
}

func (c *ProjectController) GetAll(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userId := user.(*tokens.UserClaims).ID

	query := ctx.Request.URL.Query()
	options := utils.GetPaginationFromQuery(query)

	result, err := c.service.GetAll(userId, options)
	if err != nil {
		response.SendErrorResponse(ctx, err)
		return
	}

	response.SendResponse(ctx, http.StatusOK, result)
}

func (c *ProjectController) Create(ctx *gin.Context) {
	var body req.ProjectBody
	jsonErr := ctx.ShouldBindJSON(&body)

	if jsonErr != nil {
		response.SendErrorResponse(ctx, response.ClientError(
			http.StatusBadRequest,
			jsonErr.Error(),
		))
		return
	}

	user, _ := ctx.Get("user")
	userId := user.(*tokens.UserClaims).ID

	err := c.service.Create(&body, userId)
	if err != nil {
		response.SendErrorResponse(ctx, err)
		return
	}

	response.SendResponse(ctx, http.StatusOK, nil)
}

func (c *ProjectController) Update(ctx *gin.Context) {
	var body req.ProjectBody
	jsonErr := ctx.ShouldBindJSON(&body)

	if jsonErr != nil {
		response.SendErrorResponse(ctx, response.ClientError(
			http.StatusBadRequest,
			jsonErr.Error(),
		))
		return
	}

	projectId, _ := ctx.Get("project_id")

	err := c.service.Update(&body, projectId.(int))
	if err != nil {
		response.SendErrorResponse(ctx, err)
		return
	}

	response.SendResponse(ctx, http.StatusOK, nil)
}

func (c *ProjectController) Delete(ctx *gin.Context) {
	projectId, _ := ctx.Get("project_id")

	err := c.service.Delete(projectId.(int))
	if err != nil {
		response.SendErrorResponse(ctx, err)
		return
	}

	response.SendResponse(ctx, http.StatusOK, nil)
}
