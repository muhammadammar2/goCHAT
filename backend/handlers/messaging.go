package handlers

import (
	"github.com/labstack/echo"
	websockets "github.com/muhammadammar2/goCHAT/webSockets"
	"golang.org/x/net/websocket"
)

func WebSocketHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		client := &websockets.Client{
			Hub:  &websockets.WebSocketHub,
			Conn: ws,
			Send: make(chan []byte, 256),
		}

		client.Hub.Register <- client

		go client.WritePump()
		go client.ReadPump()
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}