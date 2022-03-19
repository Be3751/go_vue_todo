package controller

import (
	"fmt"
	"net/http"

	"github.com/be3/go_vue_todo/server/model"
	"github.com/be3/go_vue_todo/server/service"
	"github.com/gin-gonic/gin"
)

func TaskList(c *gin.Context) {
	fmt.Println("GET /list")

	taskService := service.TaskService{}
	tasks, err := taskService.GetTaskList()
	if err != nil {
		fmt.Println("error")
		c.String(http.StatusInternalServerError, "Server Error")
	}
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	fmt.Println("POST /create")

	var task model.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
	taskService := service.TaskService{}
	err = taskService.SetTask(&task)
	if err != nil {
		fmt.Println("error")
		c.String(http.StatusInternalServerError, "Server Error")
	}
	c.JSON(http.StatusOK, task)
}
