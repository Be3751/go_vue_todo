package controller

import (
	"net/http"

	"github.com/be3/go_vue_todo/server/model"
	"github.com/be3/go_vue_todo/server/service"
	"github.com/gin-gonic/gin"
)

func TaskList(c *gin.Context) {
	taskService := service.TaskService{}
	tasks := taskService.GetTaskList(10)
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	task := model.Task{}
	err := c.Bind(&task)
	if err != nil {
		return
	}
	taskService := service.TaskService{}
	err = taskService.SetTask(&task)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
	c.JSON(http.StatusOK, task)
}
