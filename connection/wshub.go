// Package connection is to define a websocket server and client connect.
// Author: Arthur Zhang
// Create Date: 20190101
package connection

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

// ClientManager is a websocket manager
type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// Client is a websocket client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

// Message is an object for websocket message which is mapped to json type
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Manager define a ws server manager
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

// Start is to start a ws server
func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.Send(jsonMessage, conn)
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.Send(jsonMessage, conn)
			}
		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}

// Send is to send ws message to ws client
func (manager *ClientManager) Send(message []byte, ignore *Client) {
	for conn := range manager.Clients {
		if conn != ignore {
			conn.Send <- message
		}
	}
}

// Read read message from ws client
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			c.Socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
		Manager.Broadcast <- jsonMessage
	}
}

func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
