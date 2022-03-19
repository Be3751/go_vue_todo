package main

import (
	"github.com/be3/go_vue_todo/server/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/list", controller.TaskList)
	router.POST("/create", controller.CreateTask)
	router.GET("/get/:id", controller.ReadTask)

	router.Run(":3000")
}
