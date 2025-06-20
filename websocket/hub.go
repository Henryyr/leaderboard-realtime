package websocket

import (
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Hub struct {
	Clients map[*websocket.Conn]bool
	Mutex   sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients: make(map[*websocket.Conn]bool),
	}
}

func (h *Hub) AddClient(conn *websocket.Conn) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	h.Clients[conn] = true
}

func (h *Hub) RemoveClient(conn *websocket.Conn) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	delete(h.Clients, conn)
}

func (h *Hub) Broadcast(message []byte) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	for conn := range h.Clients {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			conn.Close()
			delete(h.Clients, conn)
		}
	}
}
