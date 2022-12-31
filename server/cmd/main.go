package main

import (
	"github.com/be3/go_vue_todo/server/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handler.SetRouter(router).Run(":3000")
}
