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
	// CmdSubscribeGeo 订阅geo推送
	CmdSubscribeGeo
	// CmdUnsubscribeGeo 取消订阅geo推送
	CmdUnsubscribeGeo
)

// OrderState 订单状态
type OrderState int

const (
	// OrderCreated 订单创建
	OrderCreated OrderState = iota
	// OrderAccepted 订单接受
	OrderAccepted
	// OrderProcessing 订单进行中
	OrderProcessing
	// OrderCanceled 订单取消
	OrderCanceled
	// OrderFinishde 订单完成
	OrderFinishde
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
	// 查询目标role
	QueryRole ClientRole `json:"qrole"`
	// 查询目标geo
	QueryGeo GeoLocation `json:"qgeo"`
}
