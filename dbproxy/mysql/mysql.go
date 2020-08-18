package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const mySQLSource = "admin:admin2020@tcp(127.0.0.1:3306)/traffic-dispatcher?charset=utf8"

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", mySQLSource)
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}

// DBConn : 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
