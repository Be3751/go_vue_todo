package usecase

import (
	"github.com/be3/go_vue_todo/server/internal/domain/gateway"
	"github.com/be3/go_vue_todo/server/internal/domain/model"
)

// 上位レイヤが利用する際に依存が疎になるようにinterfaceを定義
type GetTaskUsecase interface {
	Execute(string) (*model.Task, error)
}

type GetTaskInteractor struct {
	Repository gateway.TaskRepository
}

var _ GetTaskUsecase = &GetTaskInteractor{}

func NewGetTaskUsecase(r gateway.TaskRepository) *GetTaskInteractor {
	return &GetTaskInteractor{Repository: r}
}

func (i *GetTaskInteractor) Execute(taskID string) (*model.Task, error) {
	task, err := i.Repository.FindById(taskID)
	if err != nil {
		return nil, err
	}
	return task, nil
}
