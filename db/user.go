package db

import (
	"errors"
	"fmt"
	"time"

	"traffic-dispatcher/config"
	dbmysql "traffic-dispatcher/db/mysql"
	dbredis "traffic-dispatcher/db/redis"
	"traffic-dispatcher/model/orm"
	"traffic-dispatcher/util"
)

func genUserID(user *orm.User) string {
	if len(user.UserId) == 32 {
		return user.UserId
	}
	tmp := fmt.Sprintf("%s%d", user.UserName, user.Role)
	return util.MD5([]byte(tmp))
}

func genSessionID(user *orm.User) string {
	return config.RedisSessionPrefix + genUserID(user)
}

func genToken(user *orm.User) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(genUserID(user) + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

// Signup 用户注册
func Signup(user *orm.User) error {
	user.UserId = genUserID(user)
	if err := dbmysql.Conn().Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Signin 用户登录
func Signin(user *orm.User) (dbUser orm.User, err error) {
	dbmysql.Conn().Where("user_name = ?", user.UserName).First(&dbUser)
	if dbUser.UserName != user.UserName {
		err = errors.New("No user matched")
		return
	} else if dbUser.UserPwd != user.UserPwd {
		err = errors.New("User password does not match")
		return
	}

	rConn := dbredis.Conn()
	defer rConn.Close()

	hKey := genSessionID(&dbUser)
	rConn.Do("HSET", hKey, "token", genToken(&dbUser))
	rConn.Do("EXPIRE", hKey, 43200) // 12h
	return
}
