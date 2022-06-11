package controller

import (
	"fmt"
	"net/http"

	"github.com/be3/go_vue_todo/server/model"
	"github.com/be3/go_vue_todo/server/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func TaskList(c *gin.Context) {
	fmt.Println("POST /tasks")

	session := sessions.Default(c)
	user, err := getAuthedUser(session)
	if err != nil {
		fmt.Println("error")
		c.String(http.StatusInternalServerError, "Server Error")
	}

	// ユーザに紐づいたタスクを取得
	taskService := service.TaskService{}
	tasks, err := taskService.GetTaskList(user.Id)
	if err != nil {
		fmt.Println("error")
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

func CreateTask(c *gin.Context) {
	fmt.Println("POST /create")

	var task model.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}

	session := sessions.Default(c)
	user, err := getAuthedUser(session)
	task.User = &user

	taskService := service.TaskService{}
	err = taskService.AddTask(&task)
	if err != nil {
		fmt.Println("error")
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "succeeded to create the task."})
	}
}

func ReadTask(c *gin.Context) {
	fmt.Println("GET /read/:id")

	var task model.Task
	taskService := service.TaskService{}
	id := c.Param("id")
	task, err := taskService.GetTaskById(id)
	if err != nil {
		fmt.Println("error")
		c.String(http.StatusInternalServerError, "Server Error")
	}
	if task.Content == "" {
		c.JSON(http.StatusOK, "No such a task")
	} else {
		c.JSON(http.StatusOK, task)
	}
}

func UpdateTask(c *gin.Context) {
	fmt.Println("PUT /update/:id")

	task := model.Task{}
	err := c.BindJSON(&task) // この時点でHTTPリクエストからのJSONを取得できていない可能性が高い
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	id := c.Param("id")

	taskService := service.TaskService{}
	err = taskService.ChangeTaskById(id, &task)
	if err != nil {
		c.String(http.StatusInternalServerError, "a database error occurred.")
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "succeeded to update the task."})
	}
}

func DeleteTask(c *gin.Context) {
	fmt.Println("DELETE /delete/:id")

	id := c.Param("id")
	taskService := service.TaskService{}
	err := taskService.DeleteTaskById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "succeeded to delete the post."})
	}
}

// セッションからユーザを取得
func getAuthedUser(session sessions.Session) (model.User, error) {
	sessId := session.Get("SessionId")
	userService := service.UserService{}
	user, err := userService.GetUserBySessionId(sessId.(string))
	return user, err
}
