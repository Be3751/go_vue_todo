package main

import (
	"github.com/be3/go_vue_todo/server/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/list", controller.TaskList)
	router.POST("/create", controller.CreateTask)
	router.GET("/read/:id", controller.ReadTask)
	router.PUT("/update/:id", controller.UpdateTask)

	router.Run(":3000")
}
