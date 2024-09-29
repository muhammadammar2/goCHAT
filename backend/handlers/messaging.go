package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all origins for testing; adjust as necessary for production
    },
}

func WebSocketHandler(c echo.Context) error {
    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer ws.Close()

    for {
        _, msg, err := ws.ReadMessage()
        if err != nil {
            return err
        }
        // Echo the message back to the client (for testing purposes)
        if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
            return err
        }
    }
}
