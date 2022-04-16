package controller

import (
	"fmt"
	"net/http"

	"github.com/be3/go_vue_todo/server/crypto"
	"github.com/be3/go_vue_todo/server/model"
	"github.com/be3/go_vue_todo/server/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var userService service.UserService

// DBにユーザ情報を追加できているかを確認するための確認用ハンドラー
func UserList(c *gin.Context) {
	fmt.Println("/users")

	var users []model.User
	users, err := userService.GetUsers()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, users)
}

func SignUp(c *gin.Context) {
	fmt.Println("/signup")

	// リクエストからフォームのid, passwordの値を取得
	id := c.PostForm("id")
	pwd := c.PostForm("pwd")

	// ユーザの新規登録
	err := userService.AddUser(id, pwd)
	if err != nil {
		c.Status(http.StatusConflict)
		return
	}
	c.String(http.StatusOK, "Successful to signup!")
}

func Login(c *gin.Context) {
	fmt.Println("/login")

	// リクエストからフォームのid, passwordの値を取得
	id := c.PostForm("id")
	pwd := c.PostForm("pwd")

	user, err := userService.GetUserById(id) // idでユーザ情報の取得
	if err != nil {
		c.String(http.StatusBadRequest, "No such a user.")
	}

	// DBに格納されたハッシュ値とハッシュ化したパスワードの比較
	if err = crypto.CompHashAndPwd(user.Password, pwd); err != nil {
		c.String(http.StatusBadRequest, "Invalid password.")
	} else {
		// セッション情報の作成
		sessUuid, err := userService.CreateSession(id)
		if err != nil {
			c.String(http.StatusBadRequest, "Counldn't create a session.\n")
			return
		}
		session := sessions.Default(c) // セッション情報の作成
		session.Set("uuid", sessUuid)  // クッキーにUUIDの付与=セッションクッキーの生成
		session.Save()
		c.String(http.StatusOK, "Successful to login!\n")
	}
}

func Logout(c *gin.Context) {
	fmt.Println("/logout")
	session := sessions.Default(c) // クライアントに紐付いたセッションの取得
	session.Clear()                // セッションの破棄
	session.Save()
	c.String(http.StatusOK, "Successful to logout!")
}
