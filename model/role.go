package model

import (
	"strings"
)

// ClientRole 客户角色
type ClientRole int

const (
	// // ClientUnknown 未知role
	// ClientUnknown ClientRole = iota

	// ClientPassenger 乘客角色
	ClientPassenger ClientRole = iota
	// ClientDriver 司机角色
	ClientDriver
	// ClientAdmin 乘管理员角色
	ClientAdmin
)

func IsDriver(roleStr string) bool {
	if roleStr == "0" || strings.ToLower(roleStr) == "passenger" {
		return true
	}
	return false
}

func IsPassenger(roleStr string) bool {
	if roleStr == "1" || strings.ToLower(roleStr) == "driver" {
		return true
	}
	return false
}
