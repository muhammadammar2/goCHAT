package websockets

import (
	"golang.org/x/net/websocket"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		var message []byte
		err := websocket.Message.Receive(c.Conn, &message)
		if err != nil {
			break
		}
		c.Hub.Broadcast <- message
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for message := range c.Send {
		websocket.Message.Send(c.Conn, message)
	}
}

