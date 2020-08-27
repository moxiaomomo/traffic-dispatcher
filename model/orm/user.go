package orm

import (
	"time"
)

// User 对应table: tbl_user
type User struct {
	// 用户id
	Id uint
	// 用户角色
	Role int
	// 用户id
	UserId string
	// 用户名
	UserName string
	// 用户加密密码
	UserPwd string
	// 邮箱地址
	Email string
	// 手机号
	Phone string
	// 邮箱是否验证
	EmailValidated bool
	// 手机是否验证
	PhoneValidated bool
	// 注册时间 (定义为指针类型，可以避免0000-00-00此类问题)
	SignupAt *time.Time
	// 最近活跃时间
	LastActive *time.Time
	// profile
	Profile string
	// 状态
	Status int

	// // extra
	// Token string
}
