package handler

import (
	"net/http"

	"github.com/be3/go_vue_todo/server/internal/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type GetUserTasksHandler struct {
	getUserTasks usecase.GetUserTasksUsecase
	getAuthUser  usecase.GetAuthUserUsecase
}

func (h *GetUserTasksHandler) Func(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user, err := h.getAuthUser.Execute(session)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	tasks, err := h.getUserTasks.Execute(user.Id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, tasks)
}
