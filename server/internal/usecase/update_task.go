package usecase

import (
	"github.com/be3/go_vue_todo/server/internal/domain/gateway"
	"github.com/be3/go_vue_todo/server/internal/domain/model"
)

type UpdateTaskUsecase interface {
	Execute(*model.Task) error
}

var _ UpdateTaskUsecase = &UpdateTaskInteractor{}

type UpdateTaskInteractor struct {
	Repository gateway.TaskRepository
}

func NewUpdateTaskUsecase(r gateway.TaskRepository) *UpdateTaskInteractor {
	return &UpdateTaskInteractor{
		Repository: r,
	}
}

func (i *UpdateTaskInteractor) Execute(task *model.Task) error {
	err := i.Repository.Update(task)
	if err != nil {
		return err
	}
	return nil
}
