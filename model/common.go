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
	// ClientPassenger 乘客角色
	ClientPassenger
	// ClientAdmin 乘管理员角色
	ClientAdmin
)

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
