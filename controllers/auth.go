package controllers

import (
	"api/app"
	req "api/requests"
	"api/response"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(a *app.App) *AuthController {
	return &AuthController{service: services.NewAuthService(a)}
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

func (c *AuthController) ForgotPassword(ctx *gin.Context) {
	var body req.ForgotPasswordBody
	jsonErr := ctx.ShouldBindJSON(&body)

	if jsonErr != nil {
		response.SendErrorResponse(ctx, response.ClientError(
			http.StatusBadRequest,
			jsonErr.Error()),
		)
		return
	}

	err := c.service.ForgotPassword(&body)
	if err != nil {
		response.SendErrorResponse(ctx, err)
		return
	}

	response.SendResponse(ctx, http.StatusOK, nil)
}

func (c *AuthController) ResetPassword(ctx *gin.Context) {
	var body req.ResetPasswordBody
	jsonErr := ctx.ShouldBindJSON(&body)

	if jsonErr != nil {
		response.SendErrorResponse(ctx, response.ClientError(
			http.StatusBadRequest,
			jsonErr.Error()),
		)
		return
	}

	param := ctx.Param("resetToken")
	err := c.service.ResetPassword(param, &body)
	if err != nil {
		response.SendErrorResponse(ctx, err)
		return
	}

	response.SendResponse(ctx, http.StatusOK, nil)
}
