package handler

import (
	"net/http"

	"github.com/be3/go_vue_todo/server/internal/usecase"
	"github.com/gin-gonic/gin"
)

type SignupHandler struct {
	Usecase usecase.SignupUsecase
}

type SignupRequest struct {
	UserID  string `form:"id"`
	UserPWD string `form:"pwd"`
}

func (h *SignupHandler) Func(ctx *gin.Context) {
	var req SignupRequest
	ctx.ShouldBindQuery(req)

	err := h.Usecase.Execute(req.UserID, req.UserPWD)
	if err != nil {
		ctx.Status(http.StatusConflict)
	}
	ctx.Status(http.StatusAccepted)
}
