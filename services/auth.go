package services

import (
	"api/app"
	"api/config"
	"api/consts"
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
	cfg *config.Config
}

func NewAuthService(a *app.App) *AuthService {
	return &AuthService{
		db:  a.DB,
		cfg: a.Config,
		ctx: context.Background(),
	}
}

func (svc *AuthService) Login(body *req.LoginBody) (gin.H, *response.Error) {
	db := svc.db

	entity, err := db.User.Query().Where(user.Email(body.Email)).First(svc.ctx)
	if err != nil {
		return nil, response.ClientError(http.StatusNotFound, consts.INCORRECT_CREDENTIALS)
	}

	ok := utils.ComparePasswords(entity.Password, body.Password)
	if !ok {
		return nil, response.ClientError(http.StatusUnauthorized, consts.INCORRECT_CREDENTIALS)
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
		return nil, response.ClientError(http.StatusBadRequest, consts.USER_EXISTS)
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

func (svc *AuthService) ForgotPassword(body *req.ForgotPasswordBody) *response.Error {
	db := svc.db

	entity, err := db.User.Query().Where(user.Email(body.Email)).First(svc.ctx)
	if err != nil {
		return response.ClientError(http.StatusNotFound, consts.INCORRECT_CREDENTIALS)
	}

	token, err := tokens.GenerateJWT(tokens.UserClaims{
		ID:    entity.ID,
		Email: entity.Email,
	})

	if err != nil {
		return response.ServerError(err)
	}

	err = utils.SendEmail(&svc.cfg.Email, entity.Email, "Reset your password", token)
	if err != nil {
		return response.ServerError(err)
	}

	return nil
}

func (svc *AuthService) ResetPassword(param string, body *req.ResetPasswordBody) *response.Error {
	db := svc.db

	userClaims, err := tokens.ParseToken(param)
	if err != nil {
		return response.ClientError(http.StatusBadRequest, consts.INVALID_RESET_TOKEN)
	}

	id := userClaims.ID
	password := utils.HashPassword(body.Password)

	_, err = db.User.UpdateOneID(id).SetPassword(password).Save(svc.ctx)
	if err != nil {
		return response.ClientError(http.StatusNotFound, consts.USER_NOEXIST)
	}

	return nil
}
