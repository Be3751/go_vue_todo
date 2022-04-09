package main

import (
	"time"

	"github.com/be3/go_vue_todo/server/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORSの設定
	router.Use(Cors())

	// クッキーに認証キーを作成
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout)
	router.GET("/users", controller.UserList)
	// authUserGroup := router.Group("/auth", LoginCheck())
	// {
	// 	authUserGroup.GET("/list", controller.TaskList)
	// 	authUserGroup.POST("/create", controller.CreateTask)
	// 	authUserGroup.GET("/read/:id", controller.ReadTask)
	// 	authUserGroup.PUT("/update/:id", controller.UpdateTask)
	// 	authUserGroup.DELETE("/delete/:id", controller.DeleteTask)
	// }

	router.Run(":3000")
}

// func LoginCheck() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		session := sessions.Default(c)

// 	}
// }

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
