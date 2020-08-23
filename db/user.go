package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"traffic-dispatcher/config"
	"traffic-dispatcher/model/orm"
)

var dbCli *gorm.DB

func init() {
	var err error
	dbCli, err = gorm.Open("mysql", config.MySQLSource)
	// defer db.Close()
	if err != nil {
		log.Printf("connect mysql error: %s", err.Error())
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

// Signup 用户注册
func Signup(user *orm.User) error {
	if err := dbCli.Create(user).Error; err != nil {
		return err
	}
	return nil
}
