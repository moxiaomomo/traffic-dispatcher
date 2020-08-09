package model

// WSMsgType websocket message type
type WSMsgType int

const (
	// CmdUnknown 未知command
	CmdUnknown WSMsgType = iota
	// CmdQueryGeo 查询geo信息
	CmdQueryGeo
	// CmdReportGeo 上报geo信息
	CmdReportGeo
)

// ClientRole 客户角色
type ClientRole int

const (
	// ClientUnknown 未知role
	ClientUnknown ClientRole = iota
	// ClientDriver 司机角色
	ClientDriver
	// ClientPassanger 乘客角色
	ClientPassanger
	// ClientAdmin 乘管理员角色
	ClientAdmin
)

// GeoLocation 经纬度位置
type GeoLocation struct {
	// Lat 纬度
	Lat float64 `json:"lat"`
	// Lng 经度
	Lng float64 `json:"lng"`
}

// User 用户client
type User struct {
	Name  string `json:"name"`
	UID   string `json:"uid"`
	Token string `json:"token"`
}

// WSMessage websocket message
type WSMessage struct {
	Command WSMsgType   `json:"cmd"`
	User    User        `json:"user"`
	Role    ClientRole  `json:"role"`
	Geo     GeoLocation `json:"geo"`
	Data    []byte      `json:"data"`
}
