package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	Data any `json:"data"`
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func ClientError(status int, message string) *Error {
	return &Error{
		Status:  status,
		Message: message,
	}
}

func ServerError(err error) *Error {
	if err != nil {
		log.Errorf("A server error occurred: %s\n", err.Error())
	}

	return &Error{
		Message: "A server error occurred.",
		Status:  http.StatusInternalServerError,
	}
}

func SendResponse(ctx *gin.Context, statusCode int, data any) {
	response := Response{
		Data: data,
	}

	ctx.JSON(statusCode, response)
}

func SendErrorResponse(ctx *gin.Context, err any) {
	var response *Error

	switch res := err.(type) {
	case *Error:
		response = res
	case error:
		response = &Error{
			Status:  http.StatusInternalServerError,
			Message: res.Error(),
		}
	default:
		response = ServerError(nil)
	}

	log.Errorf("An error occurred during the request:\n%s\n", response.Message)
	ctx.AbortWithStatusJSON(response.Status, response)
}
