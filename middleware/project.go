package middleware

import (
	"api/consts"
	"api/ent"
	"api/ent/project"
	"api/ent/user"
	"api/response"
	"api/tokens"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProjectAuthorAuth(db *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userClaims, _ := ctx.Get("user")
		userId := userClaims.(*tokens.UserClaims).ID
		projectId, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			response.SendErrorResponse(ctx, response.ClientError(
				http.StatusBadRequest,
				consts.INVALID_PROJECT_ID,
			))
			return
		}

		_, err = db.Project.Query().
			Where(
				project.And(
					project.ID(projectId),
					project.HasUsersWith(user.ID(userId)),
				),
			).
			First(context.Background())

		if err != nil {
			response.SendErrorResponse(ctx, response.ClientError(
				http.StatusBadRequest,
				consts.INVALID_PROJECT_ACTION,
			))
			return
		}

		ctx.Set("project_id", projectId)
	}
}
