package main

import (
	"time"

	"github.com/be3/go_vue_todo/server/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(Cors()) // CORSの設定

	router.GET("/list", controller.TaskList)
	router.POST("/create", controller.CreateTask)
	router.GET("/read/:id", controller.ReadTask)
	router.PUT("/update/:id", controller.UpdateTask)
	router.DELETE("/delete/:id", controller.DeleteTask)

	router.Run(":3000")
}

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:8080",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"*",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	})
}
