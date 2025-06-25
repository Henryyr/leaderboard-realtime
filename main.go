package main

import (
	"leaderboard-realtime/config"
	"leaderboard-realtime/handler"
	ws "leaderboard-realtime/websocket"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	config.ConnectRedis()

	app.Post("/submit", handler.SubmitScore)
	app.Get("/leaderboard", handler.GetLeaderBoard)
	app.Get("/ws", ws.WebSocketHandler, websocket.New(ws.WebSocketConn))

	log.Fatal(app.Listen(":3100"))
}
