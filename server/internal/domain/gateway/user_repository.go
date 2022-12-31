package gateway

import (
	"github.com/be3/go_vue_todo/server/internal/domain/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindById(string) (*model.User, error)
	FindBySessionId(string) (*model.User, error)
	Create(string, string) error
	Exist(string) (bool, error)
	AttachSession(string) (uuid.UUID, error)
}
