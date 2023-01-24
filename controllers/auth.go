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

func (controller *AuthController) Login(c *gin.Context) {
	c.JSON(200, "It's all good!")
}

func (c *AuthController) Register(ctx *gin.Context) {
	var body req.RegisterBody
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		utils.SendErrorResponse(ctx, utils.ClientError(
			http.StatusBadRequest,
			err.Error()),
		)
		return
	}

	result, registerErr := c.service.Register(&body)
	if registerErr != nil {
		utils.SendErrorResponse(ctx, registerErr)
		return
	}

	utils.SendResponse(ctx, http.StatusOK, result)
}
