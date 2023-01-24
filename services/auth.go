package services

import (
	"api/ent"
	"api/ent/user"
	req "api/requests"
	"api/utils"
	"context"
	"net/http"
)

type AuthService struct {
	db  *ent.Client
	ctx context.Context
}

func NewAuthService(db *ent.Client) *AuthService {
	return &AuthService{db: db, ctx: context.Background()}
}

func (svc *AuthService) Register(body *req.RegisterBody) (*ent.User, *utils.Error) {
	db := svc.db

	users, _ := db.User.Query().Where(user.Email(body.Email)).All(svc.ctx)

	if len(users) > 0 {
		return nil, utils.ClientError(http.StatusBadRequest, utils.USER_EXISTS)
	}

	hashed := utils.HashPassword(body.Password)

	result, err := db.User.Create().
		SetName(body.Name).SetEmail(body.Email).
		SetPassword(hashed).Save(svc.ctx)

	if err != nil {
		return nil, utils.ServerError(err)
	}

	return result, nil
}
