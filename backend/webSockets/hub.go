package websockets

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

var WebSocketHub = Hub{
	Clients:    make(map[*Client]bool),
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

func (h *Hub) Run() {
	for client := range h.Register {
		h.Clients[client] = true
	}

	for client := range h.Unregister {
		if _, ok := h.Clients[client]; ok {
			delete(h.Clients, client)
			close(client.Send)
		}
	}

	for message := range h.Broadcast {
		for client := range h.Clients {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(h.Clients, client)
			}
		}
	}
}
