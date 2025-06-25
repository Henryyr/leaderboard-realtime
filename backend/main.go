package main

import (
	"leaderboard-realtime/config"
	"leaderboard-realtime/handler"
	ws "leaderboard-realtime/websocket"

	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	// Serve static files from ../public
	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.Dir("../public"),
		Browse: true,
		Index:  "index.html",
		MaxAge: 3600,
	}))

	// Connect to Redis using REDIS_URL env variable
	redisURL := os.Getenv("REDIS_URL")
	if redisURL != "" {
		// override config.ConnectRedis to use redisURL if present
		os.Setenv("REDIS_URL", redisURL)
	}
	config.ConnectRedis()

	app.Post("/submit", handler.SubmitScore)
	app.Get("/leaderboard", handler.GetLeaderBoard)
	app.Get("/ws", ws.WebSocketHandler, websocket.New(ws.WebSocketConn))

	log.Fatal(app.Listen(":3100"))
}
