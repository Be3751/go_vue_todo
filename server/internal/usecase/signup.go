package usecase

import (
	"fmt"

	"github.com/be3/go_vue_todo/server/internal/domain/gateway"
)

type SignupUsecase interface {
	Execute(string, string) error
}

var _ SignupUsecase = &SignupInteractor{}

type SignupInteractor struct {
	repository gateway.UserRepository
}

func NewSignupUsecase(r gateway.UserRepository) *SignupInteractor {
	return &SignupInteractor{
		repository: r,
	}
}

func (u *SignupInteractor) Execute(id, pwd string) error {
	if exist, err := u.repository.Exist(id); exist {
		return err
	}

	err := u.repository.Create(id, pwd)
	if err != nil {
		return fmt.Errorf("create error: %s", err.Error())
	}
	return nil
}
