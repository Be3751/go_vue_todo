package usecase

import (
	"github.com/be3/go_vue_todo/server/internal/domain/gateway"
	myerrors "github.com/be3/go_vue_todo/server/internal/errors"
	"github.com/be3/go_vue_todo/server/internal/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SigninUsecase interface {
	Execute(*gin.Context, string, string) (bool, error)
}

var _ SigninUsecase = &SigninInteractor{}

type SigninInteractor struct {
	repository gateway.UserRepository
}

func NewSigninUsecase(r gateway.UserRepository) *SigninInteractor {
	return &SigninInteractor{
		repository: r,
	}
}

func (u *SigninInteractor) Execute(ctx *gin.Context, id, pwd string) (bool, error) {
	user, err := u.repository.FindById(id)
	if err != nil {
		return false, &myerrors.NoSuchAUserError{UserID: id}
	}

	// DBに格納されたハッシュ値とハッシュ化したパスワードの比較
	if err = utils.CompHashAndPwd(user.Password, pwd); err != nil {
		return false, &myerrors.InvalidPasswordError{UserPassword: pwd}
	}

	sessUuid, err := u.repository.AttachSession(user.Id)
	if err != nil {
		return false, err
	}

	// TODO: セッション作成処理の依存を解消する
	session := sessions.Default(ctx)            // セッション情報の作成
	session.Set("SessionId", sessUuid.String()) // セッションIDをクッキーに保存=セッションクッキーの生成
	if err := session.Save(); err != nil {
		return false, err
	}
	return true, nil
}
