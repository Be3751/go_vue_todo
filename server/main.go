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
	store.Options(sessions.Options{MaxAge: 60 * 30})
	router.Use(sessions.Sessions("mysession", store))

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
		sessionId := session.Get("sessionId") // session.Set(key, val)実行時のkeyを指定
		fmt.Println(sessionId)

		// セッション情報を保持しているかを確認
		if sessionId == nil {
			fmt.Println("Couldn't authenticate...")
			c.Status(http.StatusUnauthorized)
			c.Abort()
		} else {
			fmt.Println("Successful to authenticate!")
			c.Next()
		}
	}
}

// CORS設定用ミドルウェア
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		// アクセスを許可したいオリジン
		AllowOrigins: []string{
			"http://localhost:8080", // 今回はフロントエンドアプリケーションのみを指定
		},
		// 許可したいHTTPリクエストヘッダー
		AllowHeaders: []string{
			"Cookie",     // 過去にSet-Cookieヘッダーでブラウザに保存したクッキーをクライアント側からサーバ側へ送信することを許可するヘッダー
			"Set-Cookie", // サーバ側からクライアント側にクッキーを送信することを許可するヘッダー
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding", // クライアント側がサポートしている圧縮（エンコーディング）方式をサーバ側に伝えるためのヘッダー
			"Authorization",
			"Access-Control-Allow-Credentials", // Cookie、認証ヘッダー、または TLS クライアント証明書といった資格情報をクライアント側に公開することを許可するヘッダー
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	})
}
