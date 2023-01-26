package controllers

import (
	"api/ent"
	req "api/requests"
	"api/response"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(db *ent.Client) *AuthController {
	return &AuthController{service: services.NewAuthService(db)}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var body req.LoginBody
	jsonErr := ctx.ShouldBindJSON(&body)

	if jsonErr != nil {
		response.SendErrorResponse(ctx, response.ClientError(
			http.StatusBadRequest,
			jsonErr.Error(),
		))
		return
	}

	result, err := c.service.Login(&body)
	if err != nil {
		response.SendErrorResponse(ctx, err)
		return
	}

	response.SendResponse(ctx, http.StatusOK, result)
}

func (c *AuthController) Register(ctx *gin.Context) {
	var body req.RegisterBody
	jsonErr := ctx.ShouldBindJSON(&body)

	if jsonErr != nil {
		response.SendErrorResponse(ctx, response.ClientError(
			http.StatusBadRequest,
			jsonErr.Error()),
		)
		return
	}

	result, err := c.service.Register(&body)
	if err != nil {
		response.SendErrorResponse(ctx, err)
		return
	}

	response.SendResponse(ctx, http.StatusOK, result)
}
