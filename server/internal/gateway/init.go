package gateway

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // パッケージ自体は使用せず，パッケージに含まれるmysqlを用いるためにアンダースコアを付ける
)

var Db *sql.DB
var Stmt *sql.Stmt

func init() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, host, port, dbname)

	var err error
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}
