package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SendResponse(c *gin.Context, message string, statusCode int, data any) {
	response := Response{
		Message: message,
		Data:    data,
	}

	c.JSON(statusCode, response)
}
