package websocket

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var HubInstance = NewHub()

func WebSocketHandler(c *fiber.Ctx) error {
	//Http ke Websocket
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func WebSocketConn(c *websocket.Conn) {
	HubInstance.AddClient(c)
	log.Println("Client Connected")

	defer func() {
		HubInstance.RemoveClient(c)
		log.Println("Client Disconnected")
	}()

	for {
		//read pesan Client
		if _, _, err := c.ReadMessage(); err != nil {
			break
		}
	}
}
