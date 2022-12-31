package usecase

import (
	"github.com/be3/go_vue_todo/server/internal/domain/gateway"
	"github.com/be3/go_vue_todo/server/internal/domain/model"
)

type GetUserTasksUsecase interface {
	Execute(string) ([]*model.Task, error)
}

type GetUserTasksInteractor struct {
	Repository gateway.TaskRepository
}

var _ GetUserTasksUsecase = &GetUserTasksInteractor{}

func NewGetUserTasksUsecase(r gateway.TaskRepository) *GetUserTasksInteractor {
	return &GetUserTasksInteractor{
		Repository: r,
	}
}

func (i *GetUserTasksInteractor) Execute(userID string) ([]*model.Task, error) {
	userTasks, err := i.Repository.FindByUserId(userID)
	if err != nil {
		return nil, err
	}
	return userTasks, nil
}
