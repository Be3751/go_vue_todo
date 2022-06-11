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
		return
	}

	// DBに格納されたハッシュ値とハッシュ化したパスワードの比較
	if err = crypto.CompHashAndPwd(user.Password, pwd); err != nil {
		fmt.Println("Invalid password.")
		c.Status(http.StatusBadRequest)
		return
	} else {
		// セッション情報の作成
		sessUuid, err := userService.CreateSession(user.Id)
		if err != nil {
			fmt.Println("Counldn't create a session.")
			c.Status(http.StatusBadRequest)
			return
		}
		session := sessions.Default(c)              // セッション情報の作成
		session.Set("SessionId", sessUuid.String()) // セッションIDをクッキーに保存=セッションクッキーの生成
		if err := session.Save(); err != nil {
			fmt.Println("Counldn't save a session.")
			fmt.Println(err)
			c.Status(http.StatusBadRequest)
			return
		}
		fmt.Println("Successful to login!")
		c.Status(http.StatusOK)
	}
}

func Logout(c *gin.Context) {
	fmt.Println("/logout")
	session := sessions.Default(c)                // クライアントに紐付いたセッションの取得
	session.Clear()                               // セッションの破棄
	session.Options(sessions.Options{MaxAge: -1}) // MaxAgeが負の値である場合、即座にセッションクッキーを削除
	session.Save()
	c.String(http.StatusOK, "Successful to logout!")
}
