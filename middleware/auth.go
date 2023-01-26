package middleware

import (
	"api/tokens"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			utils.SendErrorResponse(ctx, utils.ClientError(http.StatusBadRequest, utils.INVALID_TOKEN))
			return
		}

		data, err := tokens.ParseToken(token)
		if err != nil {
			utils.SendErrorResponse(ctx, utils.ClientError(http.StatusUnauthorized, utils.UNAUTHORIZED))
			return
		}

		ctx.Set("user", data)

		ctx.Next()
	}
}
