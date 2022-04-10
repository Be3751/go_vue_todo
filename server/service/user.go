package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/be3/go_vue_todo/server/crypto"
	"github.com/be3/go_vue_todo/server/model"
	"github.com/google/uuid"
)

type UserService struct{}

func (UserService) AddUser(id string, pwd string) error {
	fmt.Println("AddUser")

	// 既にユーザが登録済みかを確認
	var user model.User
	_ = Db.QueryRow("select id from users where id = ?", id).Scan(&user.Id)
	if user.Id != 0 {
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
	int_id, _ := strconv.Atoi(id)
	user = model.User{Id: int_id, Password: encPwd}
	stmt, err := Db.Prepare("insert into users (id, enc_pwd) values (?, ?)")
	if err != nil {
		fmt.Println("Prepare error")
		return err
	}
	defer stmt.Close()
	fmt.Println(user)
	_, err = stmt.Exec(user.Id, user.Password)
	if err != nil {
		fmt.Println("Exec error")
		return err
	}
	return nil
}

// DBにユーザ情報を追加できているかを確認するための確認用メソッド
func (UserService) GetUsers() ([]model.User, error) {
	fmt.Println("GetUsers")

	var users []model.User

	if rows, err := Db.Query("select id from users"); err == nil {
		for rows.Next() {
			user := model.User{}
			if err = rows.Scan(&user.Id); err != nil {
				return nil, err
			}
			fmt.Println(user)
			users = append(users, user)
		}
		rows.Close()
	}
	return users, nil
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

// 既存ユーザに対応したセッションを作成
func (UserService) CreateSession(id string) (session model.Session, err error) {
	stmt, err := Db.Prepare("insert into sessions (id, uuid) values (?, ?)")
	if err != nil {
		return
	}
	defer Db.Close()

	uuid := uuid.New()
	err = stmt.QueryRow(id, uuid.String()).Scan(&session.Id, &session.Uuid)
	return
}

// 既存ユーザに対応したセッションを取得
func (UserService) GetSession(model.Session, error) {

}
