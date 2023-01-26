package controllers

import (
	"api/ent"
	req "api/requests"
	"api/services"
	"api/utils"
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
		utils.SendErrorResponse(ctx, utils.ClientError(
			http.StatusBadRequest,
			jsonErr.Error(),
		))
		return
	}

	result, err := c.service.Login(&body)
	if err != nil {
		utils.SendErrorResponse(ctx, err)
		return
	}

	utils.SendResponse(ctx, http.StatusOK, result)
}

func (c *AuthController) Register(ctx *gin.Context) {
	var body req.RegisterBody
	jsonErr := ctx.ShouldBindJSON(&body)

	if jsonErr != nil {
		utils.SendErrorResponse(ctx, utils.ClientError(
			http.StatusBadRequest,
			jsonErr.Error()),
		)
		return
	}

	result, err := c.service.Register(&body)
	if err != nil {
		utils.SendErrorResponse(ctx, err)
		return
	}

	utils.SendResponse(ctx, http.StatusOK, result)
}
