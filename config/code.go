package config

// ErrorCode : 错误码
type ErrorCode int32

const (
	_ ErrorCode = iota + 9999
	// StatusOK : 正常
	StatusOK
	// StatusParamInvalid : 请求参数无效
	StatusParamInvalid
	// StatusServerError : 服务出错
	StatusServerError
	// StatusRegisterFailed : 注册失败
	StatusRegisterFailed
	// StatusLoginFailed : 登录失败
	StatusLoginFailed
	// StatusTokenInvalid : 10005 token无效
	StatusTokenInvalid
	// StatusNotPermitted : 10006 操作无效
	StatusNotPermitted
)
