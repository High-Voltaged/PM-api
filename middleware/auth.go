package middleware

import (
	"api/consts"
	"api/ent"
	"api/ent/user"
	"api/response"
	"api/tokens"
	"context"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Authentication(db *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" {
			response.SendErrorResponse(ctx, response.ClientError(http.StatusBadRequest, consts.INVALID_TOKEN))
			return
		}

		token := strings.SplitAfter(header, "Bearer")[1]
		data, err := tokens.ParseToken(token)
		if err != nil {
			log.Error(err)
			response.SendErrorResponse(ctx, response.ClientError(http.StatusUnauthorized, consts.UNAUTHORIZED))
			return
		}

		exists, _ := db.User.Query().
			Where(user.ID(data.ID)).
			Exist(context.Background())

		if !exists {
			response.SendErrorResponse(ctx, response.ClientError(http.StatusNotFound, consts.USER_NOEXIST))
			return
		}

		ctx.Set("user", data)

		ctx.Next()
	}
}
