package services

import (
	"api/app"
	"api/config"
	"api/ent"
	"api/ent/project"
	"api/ent/user"
	req "api/requests"
	"api/response"
	"api/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectService struct {
	db  *ent.Client
	ctx context.Context
	cfg *config.Config
}

func NewProjectService(a *app.App) *ProjectService {
	return &ProjectService{
		db:  a.DB,
		cfg: a.Config,
		ctx: context.Background(),
	}
}

// Get all projects of the currently logged in user
func (svc *ProjectService) GetAll(userId int) (gin.H, *response.Error) {
	db := svc.db

	entities, err := db.Project.Query().
		Where(project.HasUsersWith(user.ID(userId))).
		All(svc.ctx)

	if err != nil {
		return nil, response.ClientError(http.StatusNotFound, response.PROJECTS_NOT_FOUND)
	}

	result := gin.H{
		"projects": entities,
	}
	return result, nil
}

// Create a new project for the currently logged in user & set them as the creator
func (svc *ProjectService) Create(body *req.CreateProjectBody, userId int) *response.Error {
	db := svc.db

	dateInterval, err := utils.BulkStrToDate(body.Start_at, body.End_at)
	if err != nil {
		return response.ClientError(http.StatusBadRequest, "invalid date format")
	}
	_, err = db.Project.Create().
		SetName(body.Name).SetDescription(body.Description).
		SetStartAt(dateInterval[0]).SetEndAt(dateInterval[1]).
		SetCreator(userId).
		AddUserIDs(userId).Save(svc.ctx)

	if err != nil {
		return response.ServerError(err)
	}
	return nil
}
