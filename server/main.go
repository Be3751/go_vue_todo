package main

import (
	"fmt"
	"net/http"
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
	router.Use(sessions.Sessions("yourkey", store))

	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout)
	router.GET("/users", controller.UserList) // 登録済みユーザの確認用ハンドラ（開発時のみ使用）

	// ミドルウェアによる認証時のみ利用可能なハンドラ
	authUserGroup := router.Group("/auth", Authenticate())
	{
		authUserGroup.GET("/list", controller.TaskList)
		authUserGroup.POST("/create", controller.CreateTask)
		authUserGroup.GET("/read/:id", controller.ReadTask)
		authUserGroup.PUT("/update/:id", controller.UpdateTask)
		authUserGroup.DELETE("/delete/:id", controller.DeleteTask)
	}

	router.Run(":3000")
}

// 認証用ミドルウェア
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		// セッション情報を保持しているかを確認
		if userId := session.Get("uuid"); userId == nil {
			fmt.Println("You are not logged in.")
			c.Redirect(http.StatusMovedPermanently, "/login") // ログイン画面にリダイレクト
			c.Abort()                                         // ハンドラの次処理を中断
		} else {
			c.Next() // 次のハンドラに処理を移行
		}
	}
}

// CORS設定用ミドルウェア
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
