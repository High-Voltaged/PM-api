package controllers

import (
	"api/ent"
	"api/requests"
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
	var body requests.RegisterRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.service.Register()

	utils.SendResponse(ctx, "Here's your data", 200, body)
}
