package gateway

import "github.com/be3/go_vue_todo/server/internal/domain/model"

type TaskRepository interface {
	FindById(string) (*model.Task, error)
	FindByUserId(string) ([]*model.Task, error)
	Create(*model.Task) error
	Update(*model.Task) error
	Delete(*model.Task) error
}
