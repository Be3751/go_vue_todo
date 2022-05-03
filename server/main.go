package main

import (
	"github.com/be3/go_vue_todo/server/controller"
	"github.com/be3/go_vue_todo/server/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.Cors())                                  // CORSの設定
	router.Use(middleware.SetSessionCookie("secret", "mysession")) // クッキーに認証キーを作成

	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout)
	router.GET("/users", controller.UserList) // 登録済みユーザの確認用ハンドラ（開発時のみ使用）

	// ミドルウェアによる認証時のみ利用可能なハンドラ
	authUserGroup := router.Group("/auth", middleware.Authenticate())
	{
		authUserGroup.GET("/tasks", controller.TaskList)
		authUserGroup.POST("/tasks", controller.CreateTask)
		authUserGroup.GET("/tasks/:id", controller.ReadTask)
		authUserGroup.PUT("/tasks/:id", controller.UpdateTask)
		authUserGroup.DELETE("/tasks/:id", controller.DeleteTask)
	}

	router.Run(":3000")
}
