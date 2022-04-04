package service

import (
	"errors"
	"fmt"

	"github.com/be3/go_vue_todo/server/crypto"
	"github.com/be3/go_vue_todo/server/model"
)

type UserService struct{}

func (UserService) AddUser(id string, pwd string) error {
	fmt.Println("AddUser")

	var user model.User

	fmt.Println(id)
	fmt.Println(pwd)

	// 既にユーザが登録済みかを確認
	_ = Db.QueryRow("select id from users where id = ?", id).Scan(&user.Id)
	fmt.Println(user.Id)
	if user.Id != "" {
		err := errors.New("This id is already in use.")
		fmt.Println(err)
		return err
	}

	encPwd, err := crypto.EncPwd(pwd) // パスワードからハッシュ値を生成
	if err != nil {
		err := errors.New("Couldn't generate hash from the password.")
		fmt.Println(err)
		return err
	}

	// ユーザの新規登録
	user = model.User{Id: id, Password: encPwd}
	stmt, err := Db.Prepare("insert into users (id, enc_pwd) values (?, ?)")
	if err == nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(user.Id, user.Password)

	return nil
}

func (UserService) GetUserById(id string) (model.User, error) {
	fmt.Println("GetUserById")

	var user model.User
	err := Db.QueryRow("select * from users where id = ?", id).Scan(&user.Id, &user.Password)
	if err != nil {
		fmt.Println("Select error")
		return user, err
	}
	return user, nil
}
