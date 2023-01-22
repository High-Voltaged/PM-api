package controllers

import "github.com/gin-gonic/gin"

func Login(ctx *gin.Context) {
	ctx.JSON(200, "It's all good!")
}
