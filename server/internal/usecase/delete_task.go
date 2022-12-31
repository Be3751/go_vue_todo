package usecase

import (
	"github.com/be3/go_vue_todo/server/internal/domain/gateway"
	"github.com/be3/go_vue_todo/server/internal/domain/model"
)

type DeleteTaskUsecase interface {
	Execute(*model.Task) error
}

var _ DeleteTaskUsecase = &DeleteTaskInteractor{}

type DeleteTaskInteractor struct {
	Repository gateway.TaskRepository
}

func NewDeleteTaskUsecase(r gateway.TaskRepository) *DeleteTaskInteractor {
	return &DeleteTaskInteractor{
		Repository: r,
	}
}

func (i *DeleteTaskInteractor) Execute(task *model.Task) error {
	err := i.Repository.Delete(task)
	if err != nil {
		return err
	}
	return nil
}
