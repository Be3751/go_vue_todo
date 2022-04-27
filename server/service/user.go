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
	Stmt, err = Db.Prepare("insert into users (id, enc_pwd) values (?, ?)")
	if err != nil {
		fmt.Println("Insert error")
		fmt.Println(err)
		return err
	}
	defer Stmt.Close()
	fmt.Println(user)
	_, err = Stmt.Exec(user.Id, user.Password)
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
	err := Db.QueryRow("select id, enc_pwd from users where id = ?", id).Scan(&user.Id, &user.Password)
	if err != nil {
		fmt.Println("Select error")
		return user, err
	}
	return user, nil
}

// 既存ユーザに対応したセッションを作成
func (UserService) CreateSession(id string) (sessUuid uuid.UUID, err error) {
	Stmt, err = Db.Prepare("insert into sess (id, uuid) values (?, ?)")
	if err != nil {
		fmt.Println("Prepare error")
		return
	}
	defer Stmt.Close()

	sessUuid = uuid.New()
	_, err = Stmt.Exec(id, sessUuid.String())
	if err != nil {
		fmt.Println("Exec error")
		return
	}
	return
}

// 既存ユーザに対応したセッションを取得
func (UserService) GetSession(model.Session, error) {

}
