package middleware

import (
	"api/response"
	"api/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			response.SendErrorResponse(ctx, response.ClientError(http.StatusBadRequest, response.INVALID_TOKEN))
			return
		}

		data, err := tokens.ParseToken(token)
		if err != nil {
			response.SendErrorResponse(ctx, response.ClientError(http.StatusUnauthorized, response.UNAUTHORIZED))
			return
		}

		ctx.Set("user", data)

		ctx.Next()
	}
}
