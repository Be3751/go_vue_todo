package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/be3/go_vue_todo/server/internal/domain/model"
	"github.com/be3/go_vue_todo/server/internal/usecase"
	"github.com/gin-gonic/gin"
)

type PutTaskHandler struct {
	Usecase usecase.UpdateTaskUsecase
}

func (h *PutTaskHandler) Func(ctx *gin.Context) {
	var task *model.Task
	err := ctx.BindJSON(&task) // この時点でHTTPリクエストからのJSONを取得できていない可能性が高い
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	task.Id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	task.UpdatedAt = time.Now()

	err = h.Usecase.Execute(task)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.Status(http.StatusOK)
}
