package main

import (
	"github.com/be3/go_vue_todo/server/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.GET("/list", controller.TaskList)
	router.POST("/create", controller.CreateTask)
	router.Run(":3000")
}
