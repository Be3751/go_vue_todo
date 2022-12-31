package usecase

import (
	"github.com/be3/go_vue_todo/server/internal/domain/gateway"
	"github.com/be3/go_vue_todo/server/internal/domain/model"
)

type CreateTaskUsecase interface {
	Execute(*model.Task) error
}

var _ CreateTaskUsecase = &CreateTaskInteractor{}

type CreateTaskInteractor struct {
	Repository gateway.TaskRepository
}

func NewCreateTaskUsecase(r gateway.TaskRepository) *CreateTaskInteractor {
	return &CreateTaskInteractor{
		Repository: r,
	}
}

func (i *CreateTaskInteractor) Execute(task *model.Task) error {
	err := i.Repository.Create(task)
	if err != nil {
		return err
	}
	return nil
}
