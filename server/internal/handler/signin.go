package handler

import (
	"errors"
	"net/http"

	myerrors "github.com/be3/go_vue_todo/server/internal/errors"
	"github.com/be3/go_vue_todo/server/internal/usecase"
	"github.com/gin-gonic/gin"
)

type SigninHandler struct {
	Usecase usecase.SigninUsecase
}

type SigninRequest struct {
	UserID  string `form:"id"`
	UserPWD string `form:"pwd"`
}

func (h *SigninHandler) Func(ctx *gin.Context) {
	var req SigninRequest
	ctx.ShouldBindQuery(req)
	_, err := h.Usecase.Execute(ctx, req.UserID, req.UserPWD)
	if err != nil {
		if errors.Is(err, &myerrors.NoSuchAUserError{}) {
			ctx.Status(http.StatusBadRequest)
		} else if errors.Is(err, &myerrors.InvalidPasswordError{}) {
			ctx.Status(http.StatusBadRequest)
		} else {
			ctx.Status(http.StatusInternalServerError)
		}
	}
	ctx.Status(http.StatusOK)
}
