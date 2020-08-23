package config

const (
	// MySQLSource : 要连接的数据库源；
	// 其中test:test 是用户名密码；
	// 127.0.0.1:3306 是ip及端口；
	// charset=utf8 指定了数据以utf8字符编码进行传输
	MySQLSource = "admin:test123456@tcp(127.0.0.1:3306)/traffic-dispatcher?charset=utf8"

	// RedisHost redis地址
	RedisHost = "127.0.0.1:16379"
	// RedisPass 密码
	RedisPass = "test123456"

	// MongoSource mongodb地址
	MongoSource = "mongodb://127.0.0.1:27017"
)
