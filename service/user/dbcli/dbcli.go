package dbcli

import (
	"log"
	myProxy "traffic-dispatcher/dbproxy/mysql"
)

func QueryUserCountTest() int {
	stmt, err := myProxy.DBConn().Prepare(
		"select count(1) from tbl_user where status=0")
	if err != nil {
		log.Println("Failed to query, err:" + err.Error())
		return 0
	}
	defer stmt.Close()

	var rowCount int
	err = stmt.QueryRow().Scan(&rowCount)
	if err != nil {
		log.Println("Failed to query, err:" + err.Error())
		return 0
	}
	return rowCount
}
