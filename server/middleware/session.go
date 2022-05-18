package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

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

// セッションクッキー生成用ミドルウェア
func SetSessionCookie(authKey string, sessionName string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(authKey)) // クッキーに認証キーを作成
	store.Options(sessions.Options{MaxAge: 60 * 30})
	return sessions.Sessions(sessionName, store)
}
