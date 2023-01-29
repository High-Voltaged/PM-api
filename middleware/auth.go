package middleware

import (
	"api/response"
	"api/tokens"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" {
			response.SendErrorResponse(ctx, response.ClientError(http.StatusBadRequest, response.INVALID_TOKEN))
			return
		}

		token := strings.SplitAfter(header, "Bearer")[1]
		data, err := tokens.ParseToken(token)
		if err != nil {
			log.Error(err)
			response.SendErrorResponse(ctx, response.ClientError(http.StatusUnauthorized, response.UNAUTHORIZED))
			return
		}

		ctx.Set("user", data)

		ctx.Next()
	}
}
