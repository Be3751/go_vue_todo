package service

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // パッケージ自体は使用せず，パッケージに含まれるsqlを用いるためにアンダースコアを付ける
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "be3:password@(172.17.0.2:3306)/todo_db")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected to DB!")
	}
}
