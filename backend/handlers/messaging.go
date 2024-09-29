package handlers

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{}
	clients  = make(map[*websocket.Conn]bool) 
)

func WebSocketHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	clients[ws] = true
	defer delete(clients, ws)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("Received message: %s\n", msg)

		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				delete(clients, client)
			}
		}
	}

	return nil
}