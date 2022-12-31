package handler

import (
	"net/http"
	"time"

	"github.com/be3/go_vue_todo/server/internal/domain/model"
	"github.com/be3/go_vue_todo/server/internal/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type PostTaskHandler struct {
	getAuthUser usecase.GetAuthUserUsecase
	createTask  usecase.CreateTaskUsecase
}

func (h *PostTaskHandler) Func(ctx *gin.Context) {
	var task *model.Task
	err := ctx.BindJSON(task)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	task.CreatedAt = time.Now()

	user, err := h.getAuthUser.Execute(sessions.Default(ctx))
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	task.User = user

	err = h.createTask.Execute(task)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.Status(http.StatusOK)
}
