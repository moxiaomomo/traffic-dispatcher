package mysql

import (
	"os"

	// // mysql driver
	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/logger"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"traffic-dispatcher/config"
)

var dbCli *gorm.DB

func init() {
	var err error
	dbCli, err = gorm.Open("mysql", config.MySQLSource)
	// defer db.Close()
	if err != nil {
		logger.Errorf("Connect mysql error: %s", err.Error())
		os.Exit(1)
	}

	// 设置表名映射规则
	dbCli.SingularTable(true) // 禁用表名复数
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tbl_" + defaultTableName
	}
	// 设置连接
	dbCli.DB().SetMaxIdleConns(10)
	dbCli.DB().SetMaxOpenConns(100)
}

// Conn : 返回数据库连接对象
func Conn() *gorm.DB {
	return dbCli
}

// var db *sql.DB

// func init() {
// 	db, _ = sql.Open("mysql", config.MySQLSource)
// 	db.SetMaxOpenConns(1000)
// 	err := db.Ping()
// 	if err != nil {
// 		logger.Info("Failed to connect to mysql, err:" + err.Error())
// 		os.Exit(1)
// 	}
// }

// // DBConn : 返回数据库连接对象
// func DBConn() *sql.DB {
// 	return db
// }

// func CheckErr(err error) {
// 	if err != nil {
// 		logger.Fatal(err)
// 		panic(err)
// 	}
// }
