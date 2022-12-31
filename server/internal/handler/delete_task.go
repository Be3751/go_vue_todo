package handler

import (
	"net/http"

	"github.com/be3/go_vue_todo/server/internal/domain/model"
	"github.com/be3/go_vue_todo/server/internal/usecase"
	"github.com/gin-gonic/gin"
)

type DeleteTaskHandler struct {
	Usecase usecase.DeleteTaskUsecase
}

func (h *DeleteTaskHandler) Func(ctx *gin.Context) {
	var task *model.Task
	err := ctx.BindJSON(task)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	err = h.Usecase.Execute(task)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	ctx.Status(http.StatusOK)
}
