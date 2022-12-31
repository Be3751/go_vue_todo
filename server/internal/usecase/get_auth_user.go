package usecase

import (
	"github.com/be3/go_vue_todo/server/internal/domain/gateway"
	"github.com/be3/go_vue_todo/server/internal/domain/model"
	"github.com/gin-contrib/sessions"
)

type GetAuthUserUsecase interface {
	Execute(sessions.Session) (*model.User, error)
}

type GetAuthUserInteractor struct {
	Repository gateway.UserRepository
}

func NewGetAuthUserUsecase(r gateway.UserRepository) *GetAuthUserInteractor {
	return &GetAuthUserInteractor{
		Repository: r,
	}
}

func (i *GetAuthUserInteractor) Execute(session sessions.Session) (*model.User, error) {
	sessionId := session.Get("SessionId")
	user, err := i.Repository.FindBySessionId(sessionId.(string))
	if err != nil {
		return nil, err
	}
	return user, nil
}
