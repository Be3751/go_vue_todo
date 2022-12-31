package handler

import (
	"net/http"

	"github.com/be3/go_vue_todo/server/internal/usecase"
	"github.com/gin-gonic/gin"
)

type SignoutHandler struct {
	Usecase usecase.SignoutUsecase
}

func (h *SignoutHandler) Func(ctx *gin.Context) {
	err := h.Usecase.Execute(ctx)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	ctx.Status(http.StatusOK)
}
