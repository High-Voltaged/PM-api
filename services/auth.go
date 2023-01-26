package services

import (
	"api/ent"
	"api/ent/user"
	req "api/requests"
	"api/response"
	"api/tokens"
	"api/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	db  *ent.Client
	ctx context.Context
}

func NewAuthService(db *ent.Client) *AuthService {
	return &AuthService{db: db, ctx: context.Background()}
}

func (svc *AuthService) Login(body *req.LoginBody) (gin.H, *response.Error) {
	db := svc.db

	entity, err := db.User.Query().Where(user.Email(body.Email)).First(svc.ctx)
	if err != nil {
		return nil, response.ClientError(http.StatusNotFound, response.INCORRECT_CREDENTIALS)
	}

	ok := utils.ComparePasswords(entity.Password, body.Password)
	if !ok {
		return nil, response.ClientError(http.StatusUnauthorized, response.INCORRECT_CREDENTIALS)
	}

	token, err := tokens.GenerateJWT(tokens.UserClaims{
		ID:    entity.ID,
		Email: entity.Email,
	})

	if err != nil {
		return nil, response.ServerError(err)
	}

	return gin.H{"accessToken": token}, nil
}

func (svc *AuthService) Register(body *req.RegisterBody) (*ent.User, *response.Error) {
	db := svc.db

	users, _ := db.User.Query().Where(user.Email(body.Email)).All(svc.ctx)

	if len(users) > 0 {
		return nil, response.ClientError(http.StatusBadRequest, response.USER_EXISTS)
	}

	hashed := utils.HashPassword(body.Password)

	result, err := db.User.Create().
		SetName(body.Name).SetEmail(body.Email).
		SetPassword(hashed).Save(svc.ctx)

	if err != nil {
		return nil, response.ServerError(err)
	}

	return result, nil
}
