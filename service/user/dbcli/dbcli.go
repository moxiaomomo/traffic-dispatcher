package dbcli

import (
	myProxy "traffic-dispatcher/db/mysql"

	"github.com/micro/go-micro/v2/logger"
)

func QueryUserCountTest() int {
	stmt, err := myProxy.DBConn().Prepare(
		"select count(1) from tbl_user where status=0")
	if err != nil {
		logger.Info("Failed to query, err:" + err.Error())
		return 0
	}
	defer stmt.Close()

	var rowCount int
	err = stmt.QueryRow().Scan(&rowCount)
	if err != nil {
		logger.Info("Failed to query, err:" + err.Error())
		return 0
	}
	return rowCount
}
