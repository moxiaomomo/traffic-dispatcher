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

// WSMessage websocket message
type WSMessage struct {
	Command WSMsgType   `json:"cmd"`
	Role    ClientRole  `json:"role"`
	Geo     GeoLocation `json:"geo"`
	Data    []byte      `json:"data"`
}

// GeoLocation 经纬度位置
type GeoLocation struct {
	// Lat 纬度
	Lat float64 `json:"lat"`
	// Lng 经度
	Lng float64 `json:"lng"`
}
