package service

import (
	"database/sql"

	// "os"

	_ "github.com/go-sql-driver/mysql" // パッケージ自体は使用せず，パッケージに含まれるsqlを用いるためにアンダースコアを付ける
)

var Db *sql.DB
var Stmt *sql.Stmt

func init() {
	// user := os.Getenv("MYSQL_USER")
	// pass := os.Getenv("MYSQL_PASSWORD")
	// host := os.Getenv("MYSQL_HOST")
	// port := os.Getenv("MYSQL_PORT")
	// dbname := os.Getenv("MYSQL_DATABASE")
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, host, port, dbname)

	dsn := "root:root@tcp(db:3306)/todo?charset=utf8"
	var err error
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}
