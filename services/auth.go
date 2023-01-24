package services

import "api/ent"

type AuthService struct {
	db *ent.Client
}

func NewAuthService(db *ent.Client) *AuthService {
	return &AuthService{db: db}
}

func (service *AuthService) Register() {

}
