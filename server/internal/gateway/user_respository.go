package gateway

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/be3/go_vue_todo/server/internal/domain/model"
	"github.com/be3/go_vue_todo/server/internal/utils"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(d *sql.DB) *UserRepository {
	return &UserRepository{
		db: d,
	}
}

func (r *UserRepository) FindById(userID string) (*model.User, error) {
	var user *model.User
	err := r.db.QueryRow("select id, enc_pwd from users where id = ?", userID).Scan(&user.Id, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("query row error: %s", err.Error())
	}
	return user, nil
}

func (r *UserRepository) FindBySessionId(sessionId string) (*model.User, error) {
	sessionId = "\"" + sessionId + "\""
	stmt := "select id, enc_pwd from users where uuid = " + sessionId

	var user *model.User
	err := r.db.QueryRow(stmt).Scan(user.Id, user.Password)
	if err != nil {
		return nil, fmt.Errorf("query row error: %s", err.Error())
	}
	return user, nil
}

func (r *UserRepository) Create(id string, pwd string) error {
	var user model.User

	// パスワードからハッシュ値を生成
	encPwd, err := utils.EncPwd(pwd)
	if err != nil {
		return fmt.Errorf("couldn't generate hash from the password: %s", err.Error())
	}

	// ユーザの新規登録
	user = model.User{Id: id, Password: encPwd}
	stmt, err := r.db.Prepare("insert into users (id, enc_pwd) values (?, ?)")
	if err != nil {
		return fmt.Errorf("prepare error: %s", err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Password)
	if err != nil {
		return fmt.Errorf("exec error: %s", err.Error())
	}
	return nil
}

func (r *UserRepository) Exist(userId string) (bool, error) {
	// 既にユーザが登録済みかを確認
	var user model.User
	_ = r.db.QueryRow("select id from users where id = ?", userId).Scan(&user.Id)
	if user.Id != "" {
		return false, errors.New("This id is already in use.")
	}
	return true, nil
}

func (r *UserRepository) AttachSession(userId string) (uuid.UUID, error) {
	sessUuid := uuid.New()

	stmt, err := r.db.Prepare("update users set uuid=? where id = ?")
	if err != nil {
		return sessUuid, fmt.Errorf("prepare error: %s", err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessUuid.String(), userId)
	if err != nil {
		return sessUuid, fmt.Errorf("exec error: %s", err.Error())
	}
	return sessUuid, nil
}

// DBにユーザ情報を追加できているかを確認するための確認用メソッド
// func (UserService) GetUsers() ([]model.User, error) {
// 	fmt.Println("GetUsers")

// 	var users []model.User

// 	if rows, err := Db.Query("select id from users"); err == nil {
// 		for rows.Next() {
// 			user := model.User{}
// 			if err = rows.Scan(&user.Id); err != nil {
// 				return nil, err
// 			}
// 			fmt.Println(user)
// 			users = append(users, user)
// 		}
// 		rows.Close()
// 	}
// 	return users, nil
// }
