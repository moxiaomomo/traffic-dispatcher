package connection

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WsConnection : 连接结构体
type WsConnection struct {
	wsConnect *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte

	mutex  sync.Mutex
	closed bool // 保证closeChan仅被关闭一次
}

// WsUpgrader : http升级websocket协议
var WsUpgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 1024,
	// CORS设置，允许跨域请求
	CheckOrigin: func(r *http.Request) bool {
		if r.Method != "GET" {
			log.Println("method is not GET")
			return false
		}
		// if r.URL.Path != "/ws" {
		// 	log.Println("path error")
		// 	return false
		// }
		return true
	},
	Subprotocols: []string{"lbs"},
	//将获取的参数放进这个数组，问题解决
}

// InitConnection : 初始化一条连接
func InitConnection(wsConn *websocket.Conn) (conn *WsConnection, err error) {
	conn = &WsConnection{
		wsConnect: wsConn,
		inChan:    make(chan []byte, 1024),
		outChan:   make(chan []byte, 1024),
		closeChan: make(chan byte, 1),
	}
	// 创建协程处理read操作
	go conn.processRead()
	// 创建协程处理write操作
	go conn.processWrite()
	return
}

// Close : 关闭一条连接
func (conn *WsConnection) Close() {
	// 可多次调用Close (线程安全)
	conn.wsConnect.Close()

	// closeChan只关闭一次
	conn.mutex.Lock()
	defer conn.mutex.Unlock()

	if !conn.closed {
		close(conn.closeChan)
		conn.closed = true
	}
}

// ReadMessage : 读取client数据
func (conn *WsConnection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("Connection is closed")
	}
	return
}

// WriteMessage : 向client写数据
func (conn *WsConnection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("Connection is closed")
	}
	return
}

// 从ws连接读取数据
func (conn *WsConnection) processRead() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConnect.ReadMessage(); err != nil {
			goto OnErr
		}
		//阻塞在这里，等待inChan有空闲位置
		select {
		case conn.inChan <- data:
		case <-conn.closeChan: // closeChan 感知 conn断开
			goto OnErr
		}
	}

OnErr:
	conn.Close()
}

// 往ws连接写数据
func (conn *WsConnection) processWrite() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto OnErr
		}
		if err = conn.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
			goto OnErr
		}
	}

OnErr:
	conn.Close()
}
