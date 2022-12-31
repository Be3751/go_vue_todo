package usecase

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SignoutUsecase interface {
	Execute(*gin.Context) error
}

var _ SignoutUsecase = &SignoutInteractor{}

type SignoutInteractor struct {
}

func NewSignoutUsecase() *SignoutInteractor {
	return &SignoutInteractor{}
}

func (i *SignoutInteractor) Execute(ctx *gin.Context) error {
	session := sessions.Default(ctx)              // クライアントに紐付いたセッションの取得
	session.Clear()                               // セッションの破棄
	session.Options(sessions.Options{MaxAge: -1}) // MaxAgeが負の値である場合、即座にセッションクッキーを削除
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}
