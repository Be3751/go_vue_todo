package handler

import (
	"net/http"

	"github.com/be3/go_vue_todo/server/internal/usecase"
	"github.com/gin-gonic/gin"
)

type GetTaskHandler struct {
	Usecase usecase.GetTaskUsecase
}

type GetTaskRequest struct {
	taskID string `uri:"id" binding:"required"`
}

func (h *GetTaskHandler) Func(ctx *gin.Context) {
	var req GetTaskRequest
	ctx.ShouldBindUri(req)
	task, err := h.Usecase.Execute(req.taskID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, task)
}
